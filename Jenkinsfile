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

stage ("discover tests") {
	node {
		checkout scm

		testConfigs = findFiles(glob: '**/test/e2e/test_cluster_configs/**/*.json')
		testConfigs.each { cfgFile ->
			def jobCfg = readJSON(file: cfgFile.path)
			k8sVersions.each { version ->
				def jobName = cfgFile.path[cfgFile.path.indexOf("test_cluster_configs/") + 21..-6] // remove leader and trailing .json
				tasks["${version}/${jobName}"] = {
					node {
						stage("v${version}/${jobName}") {
							checkout scm

							def jobSpecificEnv = (jobCfg.env == null) ? defaultEnv.clone() : defaultEnv + jobCfg.env
							// set environment variables needed for the test script
							def envVars = [
									"ORCHESTRATOR_RELEASE=${version}",
									"API_MODEL_INPUT=${JsonOutput.toJson(jobCfg.apiModel)}",
									"FORK=${jobSpecificEnv.fork}",
									"BRANCH=${jobSpecificEnv.branch}",
									"REGIONS=${jobSpecificEnv.regions}",
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
									sh "./test/e2e/cluster.sh"
								}
							}
						}
					}
				}
			}
		}
	}
}

stage ("AKS Engine E2E Tests") {
    parallel tasks
}
