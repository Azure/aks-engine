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
        def success = true
        Integer timeoutInMinutes = TEST_TIMEOUT.toInteger()

        dir(clone_dir) {
          def img = null
          try {
            stage('Init') {
              deleteDir()
              checkout scm
              img = docker.build('aks-engine-test', '--pull .')
            }
          }
          catch(exc) {
            echo "Exception ${exc}"
            success = false
          }
          img.inside("-u root:root") {
            String errorMsg = ""
            def log_dir = pwd()+"/_logs"
            try {
              stage('Test') {
                if(success) {
                  // Create log directory
                  sh("mkdir -p ${log_dir}")
                  // Create template, deploy and test
                  env.SERVICE_PRINCIPAL_CLIENT_ID="${SPN_USER}"
                  env.SERVICE_PRINCIPAL_CLIENT_SECRET="${SPN_PASSWORD}"
                  env.TENANT_ID="${TENANT_ID}"
                  env.SUBSCRIPTION_ID="${SUBSCRIPTION_ID}"
                  env.LOCATION = "${LOCATION}"
                  env.LOGFILE = "${log_dir}/${LOCATION}.log"
                  env.CLEANUP = "${CLEANUP}"

                  env.INSTANCE_NAME = "test-acs-ci-${ORCHESTRATOR}-${env.LOCATION}-${env.BUILD_NUM}"
                  env.INSTANCE_NAME_PREFIX = "test-acs-ci"
                  env.ORCHESTRATOR = "${ORCHESTRATOR}"
                  env.CLUSTER_DEFINITION="examples/${ORCHESTRATOR}.json"
                  env.CLUSTER_SERVICE_PRINCIPAL_CLIENT_ID="${CLUSTER_SERVICE_PRINCIPAL_CLIENT_ID}"
                  env.CLUSTER_SERVICE_PRINCIPAL_CLIENT_SECRET="${CLUSTER_SERVICE_PRINCIPAL_CLIENT_SECRET}"

                  script="test/cluster-tests/${ORCHESTRATOR}/test.sh"
                  def exists = fileExists script

                  if (exists) {
                    env.VALIDATE = script
                  } else {
                    echo 'Skip validation'
                  }
                  timeout(time: timeoutInMinutes, unit: 'MINUTES') {
                    sh('./test/deploy.sh')
                  }
                }
              }
            }
            catch(exc) {
              echo "Exception ${exc}"
              success = false
              errorMsg = "Please run \"make ci\" for verification"
            }

            archiveArtifacts(allowEmptyArchive: true, artifacts: "${log_dir}/**/*.log")

            // Allow for future removal from the host
            sh("chmod -R a+rwx ${WORKSPACE}")

            if(!success) {
              currentBuild.result = "FAILURE"
              String to = "${SEND_TO}".trim()
              if(errorMsg != "") {
                if(to != "") {
                  to += ";"
                }
                to += emailextrecipients([[$class: 'CulpritsRecipientProvider']])
              }
              if(to != "") {
                def url = "${env.BUILD_URL}\n\n"
                for(String addr : to.tokenize('[ \t\n;,]+')) {
                  if(!addr.endsWith("@microsoft.com")) {
                    url = ""
                  }
                }
                gitCommit = sh(returnStdout: true, script: 'git rev-parse HEAD').trim()
                emailext(
                  to: to,
                  subject: "[AKS Engine is BROKEN] ${env.JOB_NAME} #${env.BUILD_NUM}",
                  body: "Commit: ${gitCommit}\n\n${url}${errorMsg}"
                )
              }
            }
          }
        }
      }
    }
  }
}
