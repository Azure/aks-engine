import groovy.json.*

defaultEnv = [
	CLEANUP_ON_EXIT: true,
	CREATE_VNET: false,
	UPGRADE_BRANCH: "master",
	BACK_COMPAT_VERSIONS: 0,
	] + params

def k8sVersions = ["1.12", "1.13", "1.14", "1.15", "1.16"]
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
	if(isBackCompat()) {
		def previousReleases = getPreviousVersions(params.UPGRADE_FORK)
		def backCompatVersions = []
		if(previousReleases.size() == 1) {
			backCompatVersions = ["master"]
		} else {
			backCompatVersions = ["master"] + previousReleases[0..-2]
		}
		def baseReleaseBranch = getBaseReleaseBranch(params.FORK)
		backCompatVersions.each { releaseBranch ->
			backCompatJobName = "${jobName}/back/${baseReleaseBranch}-${releaseBranch}"
			def backCompatEnv = jobCfg.env + [UPGRADE_BRANCH: releaseBranch]
			t[backCompatJobName] = runJobWithEnvironment(backCompatEnv, jobCfg.apiModel, backCompatJobName, version)
		}
	} else {
		t[jobName] = runJobWithEnvironment(jobCfg.env, jobCfg.apiModel, jobName, version)
	}
	return t
}

def isBackCompat() {
	return params.BACK_COMPAT_VERSIONS && params.BACK_COMPAT_VERSIONS.toInteger() > 0
}

def backCompatVersionCount() {
	return params.BACK_COMPAT_VERSIONS.toInteger()
}

def taskForCreateJob(jobCfg, jobName, version) {
	def t = [:]
	t[jobName] = runJobWithEnvironment(jobCfg.env, jobCfg.apiModel, jobName, version)
	return t
}

def runJobWithEnvironment(environmentVars, apiModel, jobName, version) {
	def jobSpecificEnv = defaultEnv + environmentVars
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
							API_MODEL_INPUT: "${JsonOutput.toJson(apiModel)}",
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
								// sh "./test/e2e/cluster.sh"
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

def getPreviousVersions(fork) {
	def envVars = [
			UPGRADE_FORK: fork,
		]
	withEnv(envVars.collect{ k, v -> "${k}=${v}" }) {
		output = sh(script: "./test/e2e/releases.sh", returnStdout: true)
		return output.split( '\n' )
	}
}

def getBaseReleaseBranch(fork) {
	def releases = getPreviousVersions(params.FORK)
	return releases[-1]
}

stage ("build binary") {
	node {
		retry(5){
			sh("sudo rm -rf ./*")
			checkout scm
		}

		echo "building binary for test runs"
		try {
			if(isBackCompat()) {
				def baseReleaseBranch = getBaseReleaseBranch(params.FORK)
				def envVars = [
					FORK: params.FORK,
					BRANCH: baseReleaseBranch,
				]
				withEnv(envVars.collect{ k, v -> "${k}=${v}" }) {
					sh "./test/e2e/build.sh"
				}
			} else {
				sh "./test/e2e/build.sh"
			}
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

				def isAllowedVersion = jobCfg.options?.allowedOrchestratorVersions == null ? true : version in jobCfg.options.allowedOrchestratorVersions
				if(!isAllowedVersion) {
					// the job config has limited this job to not run for this verion of the orchestrator
					echo("${jobName} is limited to ${jobCfg.options?.allowedOrchestratorVersions}; not running ${version}")
					return // this is a continue and will not exit the entire iteration
				}

				if(params.UPGRADE_CLUSTER) {
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
