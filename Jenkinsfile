import groovy.json.*

def defaultEnv = [
	FORK: "${params.FORK}",
	BRANCH: "${params.BRANCH}",
	REGION_OPTIONS: "${params.REGIONS}",
	CLEANUP_ON_EXIT: true,
	UPGRADE_CLUSTER: false,
	CREATE_VNET: false,
	SCALE_CLUSTER: false,
	]

def k8sVersions = ["1.12", "1.13", "1.14", "1.15", "1.16"]
def tasks = [:]
def testConfigs = []

stage ("build binary") {
	node {
		retry(5){
			sh("sudo rm -rf ./*")
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
			sh("sudo rm -rf ./*")
			checkout scm
		}

		testConfigs = findFiles(glob: '**/test/e2e/test_cluster_configs/**/*.json')
		testConfigs.each { cfgFile ->
			def jobCfg = readJSON(file: cfgFile.path)
			k8sVersions.each { version ->
				def jobName = cfgFile.path[cfgFile.path.indexOf("test_cluster_configs/") + 21..-6] // remove leader and trailing .json
				jobName = "v${version}/${jobName}"
				if(params.INCLUDE_JOB_REGEX.trim() && !(jobName ==~ ~/${params.INCLUDE_JOB_REGEX}/)){
					// the job is focused, so only run jobs matching the regex
					echo("This run is limited to jobs matching ${params.INCLUDE_JOB_REGEX}; not running ${jobName}")
					return // this is a continue and will not exit the entire iteration
				}

				if(params.EXCLUDE_JOB_REGEX.trim() && (jobName ==~ ~/${params.EXCLUDE_JOB_REGEX}/)){
					// the job is focused, so only run jobs matching the regex
					echo("This run excludes jobs matching ${params.EXCLUDE_JOB_REGEX}; not running ${jobName}")
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
										ORCHESTRATOR_RELEASE: "${version}",
										API_MODEL_INPUT: "${JsonOutput.toJson(jobCfg.apiModel)}",
									] + jobSpecificEnv
								withEnv(envVars.collect{ k, v -> "${k}=${v}" }) {
									// define any sensitive data needed for the test script
									def creds = [
											string(credentialsId: 'AKS_ENGINE_TENANT_ID', variable: 'TENANT_ID'),
											string(credentialsId: 'AKS_ENGINE_3014546b_CLIENT_ID', variable: 'CLIENT_ID'),
											string(credentialsId: 'AKS_ENGINE_3014546b_CLIENT_SECRET', variable: 'CLIENT_SECRET'),
											string(credentialsId: 'LOG_ANALYTICS_WORKSPACE_KEY', variable: 'LOG_ANALYTICS_WORKSPACE_KEY')
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
