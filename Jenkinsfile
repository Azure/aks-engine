import groovy.json.*

defaultEnv = [
	CLEANUP_ON_EXIT: true,
	CREATE_VNET: false,
	] + params

def k8sVersions = ["1.14", "1.15", "1.16", "1.17", "1.18"]
def latestReleasedVersion = "1.17"
def tasks = [:]
def testConfigs = []

def tasksForUpgradeJob(jobCfg, aksEngineVersions, jobName, version) {
	def jobsByName = [:]
	def t = [:]
	def versions = aksEngineVersions.orchestrators.findAll {
		it.orchestratorVersion.startsWith(version)
	}

	if(!versions) {
		println("no versions starting with ${version}, so return an empty job map")
		return t
	}

	def latestVersion = versions.last()
	if(!latestVersion.upgrades) {
		println("no versions to upgrade for ${version}, so return an empty job map")
		return t
	}

	def upgradeVersion = latestVersion.upgrades.last().orchestratorVersion
	jobCfg.env["UPGRADE_VERSIONS"] = upgradeVersion

	jobName = "${jobName}/upgrade/${upgradeVersion}"
	t[jobName] = runJobWithEnvironment(jobCfg, jobName, version)
	return t
}

def taskForCreateJob(jobCfg, jobName, version) {
	def t = [:]
	t[jobName] = runJobWithEnvironment(jobCfg, jobName, version)
	return t
}

def runJobWithEnvironment(jobCfg, jobName, version) {
	def jobSpecificEnv = defaultEnv + jobCfg.env
	def opts = jobCfg.options
	return {
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

					// set environment variables needed for the test script
					def envVars = [
							ORCHESTRATOR_RELEASE: "${version}",
							API_MODEL_INPUT: "${JsonOutput.toJson(jobCfg.apiModel)}",
							ADD_NODE_POOL_INPUT: "${JsonOutput.toJson(jobCfg.addNodePool)}",
						] + jobSpecificEnv
					withEnv(envVars.collect{ k, v -> "${k}=${v}" }) {
						// define any sensitive data needed for the test script
						def creds = [
								azureServicePrincipal(params.SERVICE_PRINCIPAL),
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

		dir('./bin') {
			unstash(name: 'aks-engine-bin')
		}

		def aksEngineAllVersions = readJSON(text: sh(script: "bin/aks-engine get-versions -o json", returnStdout: true))

		testConfigs = findFiles(glob: '**/test/e2e/test_cluster_configs/**/*.json')
		testConfigs.each { cfgFile ->
			def jobCfg = readJSON(file: cfgFile.path)
			if(!jobCfg.env) {
				jobCfg.env = [:] // ensure env exists
			}

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

				// run the job if:
				// allowedOrchestratorVersions is not set OR
				// allowedOrchestratorVersions contains version being processed OR
				// (version being process equals latestReleasedVersion AND allowedOrchestratorVersions contains "latestReleasedVersion")
				def allowedVersions = jobCfg.options?.allowedOrchestratorVersions
				def isVersionAllowed = allowedVersions == null ? true  : version in allowedVersions
				isVersionAllowed |= version == latestReleasedVersion && allowedVersions && "latestReleasedVersion" in allowedVersions

				if(!isVersionAllowed) {
					// the job config has limited this job to not run for this verion of the orchestrator
					echo("${jobName} is limited to ${jobCfg.options?.allowedOrchestratorVersions}; not running for ${version}")
					return // this is a continue and will not exit the entire iteration
				} else {
					echo("${jobName} is limted to '${jobCfg.options?.allowedOrchestratorVersions}'; running for ${version}")
				}

				if(params.UPGRADE_CLUSTER || jobCfg.env["UPGRADE_CLUSTER"])  {
					// we are upgrading, so we need to determine the next logical version to upgrade
					tasks = tasks + tasksForUpgradeJob(jobCfg, aksEngineAllVersions, jobName, version)
				} else {
					// not a upgrade job, just run the config and version
					tasks = tasks + taskForCreateJob(jobCfg, jobName, version)
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
