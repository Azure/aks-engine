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
            stage("cluster create") {
                node {
                    withEnv(["ORCHESTRATOR_RELEASE=${version}"]) {
                        sh "ls -laR"
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
