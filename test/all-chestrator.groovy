#!/usr/bin/env groovy

node("slave") {
  withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'AZURE_CLI_SPN_AKS_TEST',
                  passwordVariable: 'SPN_PASSWORD', usernameVariable: 'SPN_USER']]) {
    timestamps {
      wrap([$class: 'AnsiColorBuildWrapper', 'colorMapName': 'XTerm']) {
        def jobname = "${JOBNAME}"
        def tests = [:]

        def pairs = "${ORCHESTRATOR_LOCATION}".tokenize('|')
        for(i = 0; i < pairs.size(); i++) {
          def pair = pairs[i].tokenize('[ \t\n]+')
          if(pair.size() != 2) {
            echo "Skipping '"+pairs[i]+"'"
            continue
          }
          def orchestrator = pair[0]
          def location = pair[1]
          def name = "${orchestrator}-${location}"

          tests[name] = {
            stage(name) {
              build job: jobname,
              parameters:
               [[$class: 'StringParameterValue', name: 'ORCHESTRATOR', value: orchestrator],
                [$class: 'StringParameterValue', name: 'LOCATION',     value: location]]
            }
          }
        }
        parallel tests
      }
    }
  }
}
