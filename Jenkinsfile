import groovy.json.*

def defaultEnv = [
	fork: "${params.FORK}",
	branch: "${params.BRANCH}",
	regions: "${params.REGIONS}",
	cleanupOnExit: true,
	upgradeCluster: false,
	createVNet: false,
	scaleCluster: false,
	]

def k8sVersions = ["1.12", "1.13", "1.14", "1.15", "1.16"]
def tasks = [:]
def testConfigs = []

stage ("build binary") {
	node {
		retry(5){
			sh("sudo rm -rf ./bin ./_output ./_logs")
			checkout scm
		}

		echo "building binary for test runs"
		try {
			sh "./test/e2e/build.sh"
		} finally {
			sh "./test/e2e/jenkins_reown.sh"
		}

		dir('./bin') {
			stash(includes: 'aks-engine', name: 'aks-engine-bin')
		}

		archiveArtifacts(artifacts: 'bin/**/*')
	}
}

stage ("discover tests") {
	node {
		retry(5){
			sh("sudo rm -rf ./bin ./_output ./_logs")
			checkout scm
		}

		testConfigs = findFiles(glob: '**/test/e2e/test_cluster_configs/**/*.json')
		testConfigs.each { cfgFile ->
			def jobCfg = readJSON(file: cfgFile.path)
			k8sVersions.each { version ->
				def jobName = cfgFile.path[cfgFile.path.indexOf("test_cluster_configs/") + 21..-6] // remove leader and trailing .json
				jobName = "v${version}/${jobName}"
				if(params.JOB_FOCUS_REGEX.trim() && !(jobName ==~ ~/${params.JOB_FOCUS_REGEX}/)){
					// the job is focused, so only run jobs matching the regex
					echo("This run is limited to jobs matching ${params.JOB_FOCUS_REGEX}; not running ${jobName}")
					return // this is a continue and will not exit the entire iteration
				}

				def isAllowedVersion = jobCfg.options?.allowedOrchestratorVersions == null ? true : version in jobCfg.options.allowedOrchestratorVersions
				if(!isAllowedVersion) {
					// the job config has limited this job to not run for this verion of the orchestrator
					echo("${jobName} is limited to ${jobCfg.options?.allowedOrchestratorVersions}; not running ${version}")
					return // this is a continue and will not exit the entire iteration
				}

				tasks[jobName] = {
					node {
						ws("${env.JOB_NAME}-${jobName}") {
							stage(jobName) {
								retry(5){
									sh("sudo rm -rf ./bin ./_output ./_logs")
									cleanWs()
									checkout scm
								}

								dir('./bin') {
									unstash(name: 'aks-engine-bin')
								}

								def jobSpecificEnv = (jobCfg.env == null) ? defaultEnv.clone() : defaultEnv + jobCfg.env
								// set environment variables needed for the test script
								def envVars = [
										"ORCHESTRATOR_RELEASE=${version}",
										"API_MODEL_INPUT=${JsonOutput.toJson(jobCfg.apiModel)}",
										"FORK=${jobSpecificEnv.fork}",
										"BRANCH=${jobSpecificEnv.branch}",
										"REGION_OPTIONS=${jobSpecificEnv.regions}",
										"CLEANUP_ON_EXIT=${jobSpecificEnv.cleanupOnExit}",
										"UPGRADE_CLUSTER=${jobSpecificEnv.upgradeCluster}",
										"CREATE_VNET=${jobSpecificEnv.createVNet}",
										"SCALE_CLUSTER=${jobSpecificEnv.scaleCluster}",
									]
								withEnv(envVars) {
									// define any sensitive data needed for the test script
									def creds = [
											string(credentialsId: 'AKS_ENGINE_TENANT_ID', variable: 'TENANT_ID'),
											string(credentialsId: 'AKS_ENGINE_3014546b_CLIENT_ID', variable: 'CLIENT_ID'),
											string(credentialsId: 'AKS_ENGINE_3014546b_CLIENT_SECRET', variable: 'CLIENT_SECRET')
										]

									withCredentials(creds) {
										echo "Running tests for: ${jobName}"
										try {
											echo "EXECUTOR_NUMBER :: $EXECUTOR_NUMBER"
											echo "NODE_NAME :: $NODE_NAME"
											sh "./test/e2e/cluster.sh"
										} finally {
											sh "./test/e2e/jenkins_reown.sh"
										}
									}
								}

								archiveArtifacts(artifacts: '_output/**/*', allowEmptyArchive: true)
								archiveArtifacts(artifacts: '_logs/**/*', allowEmptyArchive: true)
							}
						}
					}
				}
			}
		}
	}
}

stage ("AKS Engine E2E Tests") {
	throttle(['k8s-matrix']) {
		parallel tasks
	}
}
