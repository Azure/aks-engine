#!/usr/bin/env groovy

node("slave") {
  withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'AZURE_CLI_SPN_ACS_TEST',
                  passwordVariable: 'SPN_PASSWORD', usernameVariable: 'SPN_USER']]) {
    timestamps {
      wrap([$class: 'AnsiColorBuildWrapper', 'colorMapName': 'XTerm']) {
        env.GOPATH="${WORKSPACE}"
        env.PATH="${env.PATH}:${env.GOPATH}/bin"
        def clone_dir = "${env.GOPATH}/src/github.com/Azure/aks-engine"
        env.HOME=clone_dir
        String sendTo = "${SEND_TO}".trim()
        Integer timeoutInMinutes = STAGE_TIMEOUT.toInteger()
        def autoclean="${AUTOCLEAN}"

        dir(clone_dir) {
          def img = null
          stage('Init') {
            deleteDir()
            checkout scm
            img = docker.build('aks-engine-test', '--pull .')
          }

          img.inside("-u root:root") {
            def success = true
            def junit_dir = "_junit"
            def prefix = ""
            try {
              stage('Setup') {
                // Set up Azure
                sh("az login --service-principal -u ${SPN_USER} -p ${SPN_PASSWORD} --tenant ${TENANT_ID}")
                sh("az account set --subscription ${SUBSCRIPTION_ID}")

                env.SERVICE_PRINCIPAL_CLIENT_ID="${SPN_USER}"
                env.SERVICE_PRINCIPAL_CLIENT_SECRET="${SPN_PASSWORD}"
                env.TENANT_ID="${TENANT_ID}"
                env.SUBSCRIPTION_ID="${SUBSCRIPTION_ID}"
                env.CLUSTER_SERVICE_PRINCIPAL_CLIENT_ID="${CLUSTER_SERVICE_PRINCIPAL_CLIENT_ID}"
                env.CLUSTER_SERVICE_PRINCIPAL_CLIENT_SECRET="${CLUSTER_SERVICE_PRINCIPAL_CLIENT_SECRET}"

                // First check to see if var exists in context, then check for true-ness
                // In Groovy, null and empty strings are false...
                if(getBinding().hasVariable("CUSTOM_HYPERKUBE_SPEC") && CUSTOM_HYPERKUBE_SPEC) {
                    env.CUSTOM_HYPERKUBE_SPEC="${CUSTOM_HYPERKUBE_SPEC}"
                }

                sh("printf 'acs-features-test%x' \$(date '+%s') > INSTANCE_NAME_PREFIX")
                prefix = readFile('INSTANCE_NAME_PREFIX').trim()
                // Create report directory
                sh("mkdir -p ${junit_dir}")
                // Build and test aks-engine
                sh('make ci')
              }
              def pairs = "${SCENARIOS_LOCATIONS}".tokenize('|')
              for(i = 0; i < pairs.size(); i++) {
                def pair = pairs[i].tokenize('[ \t\n]+')
                if(pair.size() != 2) {
                  echo "Skipping '"+pairs[i]+"'"
                  continue
                }
                def subdir = pair[0]
                def names = sh(returnStdout: true, script: "cd examples; ls ${subdir}/*.json").split("\\r?\\n")
                env.LOCATION = pair[1]
                for(j = 0; j< names.size(); j++) {
                  def name = names[j].trim()
                  env.CLUSTER_DEFINITION = pwd()+"/examples/${name}"
                  env.INSTANCE_NAME = "${prefix}-${i}-${j}"
                  env.RESOURCE_GROUP = "test-acs-${subdir}-${env.LOCATION}-${env.BUILD_NUM}-${i}-${j}"
                  env.DEPLOYMENT_NAME = "${env.RESOURCE_GROUP}"
                  env.ORCHESTRATOR = sh(returnStdout: true, script: './test/step.sh get_orchestrator_type').trim()
                  env.LOGFILE = pwd()+"/${junit_dir}/${name}.log"
                  env.CLEANUP = "y"
                  // Generate and deploy template, validate deployments
                  try {
                    stage(name) {
                      def scripts = ["generate_template.sh", "deploy_template.sh"]
                      if(env.ORCHESTRATOR == "kubernetes") {
                        scripts += "validate_deployment.sh"
                      }
                      for(k = 0; k < scripts.size(); k++) {
                        def script = scripts[k]
                        def test = "${name}.${script}"
                        sh("mkdir -p ${junit_dir}/${test}")
                        sh("cp ./test/shunit/${script} ${junit_dir}/${test}/t.sh")
                        timeout(time: timeoutInMinutes, unit: 'MINUTES') {
                          sh("cd ${junit_dir}; shunit.sh -t ${test} > ${test}/junit.xml")
                        }
                        sh("grep 'failures=\"0\"' ${junit_dir}/${test}/junit.xml")
                      }
                    }
                  }
                  catch(exc) {
                    env.CLEANUP = autoclean
                    echo "Exception in [${name}] : ${exc}"
                  }
                  // Clean up
                  try {
                    sh('./test/step.sh cleanup')
                  }
                  catch(exc) {
                    echo "Exception ${exc}"
                  }
                } // for (j = 0; j <files...
              } // for (i = 0; i <subdirs...
              // Generate reports
              try {
                junit("${junit_dir}/**/junit.xml")
                archiveArtifacts(allowEmptyArchive: true, artifacts: "${junit_dir}/**/*.log")
                if(currentBuild.result == "UNSTABLE") {
                  currentBuild.result = "FAILURE"
                  if(sendTo != "") {
                    emailext(
                      to: "${sendTo}",
                      subject: "[AKS Engine Jenkins Failure] ${env.JOB_NAME} #${env.BUILD_NUM}",
                      body: "${env.BUILD_URL}testReport")
                  }
                }
              }
              catch(exc) {
                echo "Exception ${exc}"
              }
            }
            catch(exc) {
              currentBuild.result = "FAILURE"
              echo "Exception ${exc}"
            }
            // Allow for future removal from the host
            sh("chmod -R a+rwx ${WORKSPACE}")
          }
        }
      }
    }
  }
}
