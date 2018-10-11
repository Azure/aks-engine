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
        String locations_str = "${LOCATIONS}"
        String sendTo = "${SEND_TO}".trim()
        Integer timeoutInMinutes = STAGE_TIMEOUT.toInteger()
        def autoclean="${AUTOCLEAN}"

        if(locations_str.equals("all")) {
          locations_str = "\
australiaeast \
brazilsouth \
canadacentral canadaeast \
centralindia southindia \
centralus eastus2 eastus northcentralus southcentralus westcentralus westus2 westus \
eastasia southeastasia \
koreacentral koreasouth \
japaneast japanwest \
northeurope westeurope \
uksouth ukwest"
        }
        def locations = locations_str.tokenize('[ \t\n]+')

        dir(clone_dir) {
          def img = null
          stage('Init') {
            deleteDir()
            checkout scm
            img = docker.build('aks-engine-test', '--pull .')
          }

          img.inside("-u root:root") {
            def junit_dir = "_junit"
            try {
              String canonicalName = sh(returnStdout: true, script: 'echo "${CLUSTER_DEFINITION%.*}" | sed "s/\\//_/g"').trim()
              stage('Setup') {
                // Set up Azure
                sh("az login --service-principal -u ${SPN_USER} -p ${SPN_PASSWORD} --tenant ${TENANT_ID}")
                sh("az account set --subscription ${SUBSCRIPTION_ID}")
                // Create report directory
                sh("mkdir ${junit_dir}")
                // Build and test aks-engine
                sh('make ci')
                // Create template
                env.CLUSTER_DEFINITION = "examples/${CLUSTER_DEFINITION}"
                env.ORCHESTRATOR = sh(returnStdout: true, script: './test/step.sh get_orchestrator_type').trim()
                sh("printf 'acs-test%x' \$(date '+%s') > INSTANCE_NAME")
                env.INSTANCE_NAME = readFile('INSTANCE_NAME').trim()
                env.CLUSTER_SERVICE_PRINCIPAL_CLIENT_ID="${CLUSTER_SERVICE_PRINCIPAL_CLIENT_ID}"
                env.CLUSTER_SERVICE_PRINCIPAL_CLIENT_SECRET="${CLUSTER_SERVICE_PRINCIPAL_CLIENT_SECRET}"
                timeout(time: timeoutInMinutes, unit: 'MINUTES') {
                  sh('./test/step.sh generate_template')
                }
              }

              for (i = 0; i <locations.size(); i++) {
                env.LOCATION = locations[i]
                env.RESOURCE_GROUP = "test-acs-svc-${canonicalName}-${env.LOCATION}-${env.BUILD_NUM}"
                env.DEPLOYMENT_NAME = "${env.RESOURCE_GROUP}"
                env.LOGFILE = pwd()+"/${junit_dir}/${canonicalName}.${env.LOCATION}.log"
                env.CLEANUP = "y"
                def ok = true
                // Deploy
                try {
                  stage("${env.LOCATION} deploy") {
                    def test = "deploy-${env.LOCATION}"
                    sh("mkdir -p ${junit_dir}/${test}")
                    sh("cp ./test/shunit/deploy_template.sh ${junit_dir}/${test}/t.sh")
                    timeout(time: timeoutInMinutes, unit: 'MINUTES') {
                      sh("cd ${junit_dir}; shunit.sh -t ${test} > ${test}/junit.xml")
                    }
                    sh("grep 'failures=\"0\"' ${junit_dir}/${test}/junit.xml")
                  }
                }
                catch(exc) {
                  env.CLEANUP = autoclean
                  echo "Exception in [deploy ${canonicalName}/${env.LOCATION}] : ${exc}"
                  ok = false
                }
                // Verify deployment
                try {
                  stage("${env.LOCATION} validate") {
                    if(ok) {
                      def test = "validate-${env.LOCATION}"
                      sh("mkdir -p ${junit_dir}/${test}")
                      sh("cp ./test/shunit/validate_deployment.sh ${junit_dir}/${test}/t.sh")
                      timeout(time: timeoutInMinutes, unit: 'MINUTES') {
                        sh("cd ${junit_dir}; shunit.sh -t ${test} > ${test}/junit.xml")
                      }
                      sh("grep 'failures=\"0\"' ${junit_dir}/${test}/junit.xml")
                    }
                    else {
                      echo "Skipped verification for ${env.RESOURCE_GROUP}"
                    }
                  }
                }
                catch(exc) {
                  env.CLEANUP = autoclean
                  echo "Exception in [validate ${canonicalName}/${env.LOCATION}] : ${exc}"
                }
                // Clean up
                try {
                  sh('./test/step.sh cleanup')
                }
                catch(exc) {
                  echo "Exception ${exc}"
                }
              } // for (i = 0; i <locations...
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
