def k8sVersions = ["1.15"]
def clusterConfigurations = ["default-config"]
def tasks = [:]

stage("Before") {
    node {
        echo "run this static stuff before the matrix"
    }
}

for(int i=0; i< k8sVersions.size(); i++) {
    def version = k8sVersions[i]
    for(int j=0; j< clusterConfigurations.size(); j++) {
        def clusterConfig = clusterConfigurations[j]
        tasks["${version}/${clusterConfig}"] = {
            stages {
                stage("cluster create") {
                    steps {
                        println "export ORCHESTRATOR_RELEASE=${version}"
                        println "./test/e2e/cluster.sh"
                    }
                }
            }
        }
    }
}

stage ("Matrix") {
    parallel tasks
}

stage("After") {
    node {
        echo "run this static stuff after the entire"
    }
}
