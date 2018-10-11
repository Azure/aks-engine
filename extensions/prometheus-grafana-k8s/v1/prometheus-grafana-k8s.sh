#!/bin/bash
set -x

echo $(date) " - Starting Script"

echo $(date) " - Setting kubeconfig"
export KUBECONFIG=/var/lib/kubelet/kubeconfig

echo $(date) " - Waiting for API Server to start"
kubernetesStarted=1
for i in {1..600}; do
    if [ -e /usr/local/bin/kubectl ]
    then
        /usr/local/bin/kubectl cluster-info
        if [ "$?" = "0" ]
        then
            echo "kubernetes started"
            kubernetesStarted=0
            break
        fi
    else
        /usr/bin/docker ps | grep apiserver
        if [ "$?" = "0" ]
        then
            echo "kubernetes started"
            kubernetesStarted=0
            break
        fi
    fi
    sleep 1
done
if [ $kubernetesStarted -ne 0 ]
then
    echo "kubernetes did not start"
    exit 1
fi

master_nodes() {
    kubectl get no -L kubernetes.io/role -l kubernetes.io/role=master --no-headers -o jsonpath="{.items[*].metadata.name}" | tr " " "\n" | sort | head -n 1
}

wait_for_master_nodes() {
    ATTEMPTS=90
    SLEEP_TIME=10

    ITERATION=0
    while [[ $ITERATION -lt $ATTEMPTS ]]; do
        echo $(date) " - Is kubectl returning master nodes? (attempt $(( $ITERATION + 1 )) of $ATTEMPTS)"

        FIRST_K8S_MASTER=$(master_nodes)

        if [[ -n $FIRST_K8S_MASTER ]]; then
            echo $(date) " - kubectl is returning master nodes"
            return
        fi

        ITERATION=$(( $ITERATION + 1 ))
        sleep $SLEEP_TIME
    done

    echo $(date) " - kubectl failed to return master nodes in the alotted time"
    return 1
}

agent_nodes() {
    kubectl get no -L kubernetes.io/role -l kubernetes.io/role=agent --no-headers -o jsonpath="{.items[*].metadata.name}" | tr " " "\n" | sort | head -n 1
}

wait_for_agent_nodes() {
    ATTEMPTS=90
    SLEEP_TIME=10

    ITERATION=0
    while [[ $ITERATION -lt $ATTEMPTS ]]; do
        echo $(date) " - Is kubectl returning agent nodes? (attempt $(( $ITERATION + 1 )) of $ATTEMPTS)"

        FIRST_K8S_AGENT=$(agent_nodes)

        if [[ -n $FIRST_K8S_AGENT ]]; then
            echo $(date) " - kubectl is returning agent nodes"
            return
        fi

        ITERATION=$(( $ITERATION + 1 ))
        sleep $SLEEP_TIME
    done

    echo $(date) " - kubectl failed to return agent nodes in the alotted time"
    return 1
}

should_this_node_run_extension() {
    FIRST_K8S_MASTER=$(master_nodes)
    if [[ $FIRST_K8S_MASTER = $(hostname) ]]; then
        echo $(date) " - Local node $(hostname) is found to be the first master node $FIRST_K8S_MASTER"
        return
    else
        FIRST_K8S_AGENT=$(agent_nodes)
        if [[ $FIRST_K8S_AGENT = $(hostname) ]]; then
            echo $(date) " - Local node $(hostname) is found to be the first agent node $FIRST_K8S_AGENT"
            return
        else
            echo $(date) " - Local node $(hostname) is not the first master node $FIRST_K8S_MASTER or the first agent node $FIRST_K8S_AGENT"
            return 1
        fi
    fi
}

storageclass_param() {
	kubectl get no -l kubernetes.io/role=agent -l storageprofile=managed --no-headers -o jsonpath="{.items[0].metadata.name}" > /dev/null 2> /dev/null
	if [[ $? -eq 0 ]]; then
		echo '--set server.persistentVolume.storageClass=managed-standard'
	fi
}

wait_for_tiller() {
    ATTEMPTS=90
    SLEEP_TIME=10

    ITERATION=0
    while [[ $ITERATION -lt $ATTEMPTS ]]; do
        echo $(date) " - Is Helm running? (attempt $(( $ITERATION + 1 )) of $ATTEMPTS)"

        helm version > /dev/null 2> /dev/null

        if [[ $? -eq 0 ]]; then
            echo $(date) " - Helm is running"
            return
        fi

        ITERATION=$(( $ITERATION + 1 ))
        sleep $SLEEP_TIME
    done

    echo $(date) " - Helm failed to start in the alotted time"
    return 1
}

install_helm() {
    echo $(date) " - Downloading helm"
    curl https://storage.googleapis.com/kubernetes-helm/helm-v2.6.2-linux-amd64.tar.gz > helm-v2.6.2-linux-amd64.tar.gz
    tar -zxvf helm-v2.6.2-linux-amd64.tar.gz
    mv linux-amd64/helm /usr/local/bin/helm
    echo $(date) " - Downloading prometheus values"

    curl "$1" > prometheus_values.yaml

    sleep 10

    echo $(date) " - helm version"
    helm version
    helm init

    echo $(date) " - helm installed"
}

daemonset_api() {
    SERVER_VERSION_REGEX="Server Version: v([0-9]+)\.([0-9]+)\.([0-9]+)$"
    K8S_SERVER_VERSION=$(kubectl version --short | grep Server)
    K8S_SERVER_VERSION_MAJOR=$(echo "$K8S_SERVER_VERSION" | sed -r "s/$SERVER_VERSION_REGEX/\1/g")
    K8S_SERVER_VERSION_MINOR=$(sed -r "s/$SERVER_VERSION_REGEX/\2/g" "$K8S_SERVER_VERSION")

    if [[ $K8S_SERVER_VERSION_MAJOR -gt 1 ]]; then
        echo "apps/v1"
    elif [[ $K8S_SERVER_VERSION_MAJOR -eq 1 ]]; then
        if [[ $K8S_SERVER_VERSION_MINOR -gt 8 ]]; then
            echo "apps/v1"
        else
            echo "apps/v1beta2"
        fi
    fi
}

install_cadvisor() {
    echo "$(date) - Installing cAdvisor"
    local NAMESPACE="$1"
    echo "$(date) - Using namespace $NAMESPACE"
    local CADVISOR_DS_CONFIG_URL="$2"
    echo "$(date) - Using cAdvisor config at $CADVISOR_DS_CONFIG_URL"

    curl -o cadvisor_ds.yml "$CADVISOR_DS_CONFIG_URL"
    DAEMONSET_API=$(daemonset_api)
    echo "$(date) - Using DaemonSet api group $DAEMONSET_API"
    sed -i 's|DAEMONSET_API|'"$DAEMONSET_API"'|g' cadvisor_ds.yml
    kubectl apply -f ./cadvisor_ds.yml
}

update_helm() {
    echo $(date) " - Updating Helm repositories"
    helm repo update
}

install_prometheus() {
    PROM_RELEASE_NAME=monitoring
    NAMESPACE=$1

    echo $(date) " - Installing the Prometheus Helm chart"

    ATTEMPTS=90
    SLEEP_TIME=10

    ITERATION=0
    while [[ $ITERATION -lt $ATTEMPTS ]]; do
        helm install -f prometheus_values.yaml \
            --name $PROM_RELEASE_NAME \
            --version 5.1.3 \
            --namespace $NAMESPACE stable/prometheus $(storageclass_param)

        if [[ $? -eq 0 ]]; then
            echo $(date) " - Helm install successfully completed"
            break
        else
            echo $(date) " - Helm install returned a non-zero exit code. Retrying."
        fi

        ITERATION=$(( $ITERATION + 1 ))
        sleep $SLEEP_TIME
    done

    PROM_POD_PREFIX="$PROM_RELEASE_NAME-prometheus-server"
    DESIRED_POD_STATE=Running

    ITERATION=0
    while [[ $ITERATION -lt $ATTEMPTS ]]; do
        echo $(date) " - Is the prometheus server pod ($PROM_POD_PREFIX-*) running? (attempt $(( $ITERATION + 1 )) of $ATTEMPTS)"

        kubectl get po -n $NAMESPACE --no-headers |
            awk '{print $1 " " $3}' |
            grep $PROM_POD_PREFIX |
            grep -q $DESIRED_POD_STATE

        if [[ $? -eq 0 ]]; then
            echo $(date) " - $PROM_POD_PREFIX-* is $DESIRED_POD_STATE"
            break
        fi

        ITERATION=$(( $ITERATION + 1 ))
        sleep $SLEEP_TIME
    done
}

install_grafana() {
    GF_RELEASE_NAME=dashboard
    NAMESPACE=$1

    echo $(date) " - Installing the Grafana Helm chart"
    helm install \
        --name $GF_RELEASE_NAME \
        --version 0.8.5 \
        --namespace $NAMESPACE stable/grafana $(storageclass_param)

    GF_POD_PREFIX="$GF_RELEASE_NAME-grafana"
    DESIRED_POD_STATE=Running

    ATTEMPTS=90
    SLEEP_TIME=10

    ITERATION=0
    while [[ $ITERATION -lt $ATTEMPTS ]]; do
        echo $(date) " - Is the grafana pod ($GF_POD_PREFIX-*) running? (attempt $(( $ITERATION + 1 )) of $ATTEMPTS)"

        kubectl get po -n $NAMESPACE --no-headers |
            awk '{print $1 " " $3}' |
            grep $GF_POD_PREFIX |
            grep -q $DESIRED_POD_STATE

        if [[ $? -eq 0 ]]; then
            echo $(date) " - $GF_POD_PREFIX-* is $DESIRED_POD_STATE"
            break
        fi

        ITERATION=$(( $ITERATION + 1 ))
        sleep $SLEEP_TIME
    done
}

ensure_k8s_namespace_exists() {
    NAMESPACE_TO_EXIST="$1"

    kubectl get ns $NAMESPACE_TO_EXIST > /dev/null 2> /dev/null
    if [[ $? -ne 0 ]]; then
        echo $(date) " - Creating namespace $NAMESPACE_TO_EXIST"
        kubectl create ns $NAMESPACE_TO_EXIST
    else
        echo $(date) " - Namespace $NAMESPACE_TO_EXIST already exists"
    fi
}

NAMESPACE=default
RAW_PROMETHEUS_CHART_VALS="https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/prometheus-grafana-k8s/v1/prometheus_values.yaml"
CADVISOR_CONFIG_URL="https://raw.githubusercontent.com/Azure/aks-engine/master/extensions/prometheus-grafana-k8s/v1/cadvisor_daemonset.yml"

# retrieve and parse extension parameters
if [[ -n "$1" ]]; then
    IFS=';' read -ra INPUT <<< "$1"
    if [[ -n "${INPUT[0]}" ]]; then
        NAMESPACE="${INPUT[0]}"
        echo "$(date) - Custom namespace specified: $NAMESPACE"
    fi
    if [[ -n "${INPUT[1]}" ]]; then
        RAW_PROMETHEUS_CHART_VALS="${INPUT[1]}"
        echo "$(date) - Custom prometheus chart values url specified: $RAW_PROMETHEUS_CHART_VALS"
    fi
    if [[ -n "${INPUT[2]}" ]]; then
        CADVISOR_CONFIG_URL="${INPUT[2]}"
        echo "$(date) - Custom cAdvisor config url specified: $CADVISOR_CONFIG_URL"
    fi
fi

# this extension should only run on a single node
# the logic to decide whether or not this current node
# should run the extension is to alphabetically determine
# if this local machine is the first in the list of master nodes
# if it is, then run the extension. if not, exit
wait_for_master_nodes
if [[ $? -ne 0 ]]; then
    echo $(date) " - Error while waiting for kubectl to output master nodes. Exiting"
    exit 1
fi

wait_for_agent_nodes
if [[ $? -ne 0 ]]; then
    echo $(date) " - Error while waiting for kubectl to output agent nodes. Exiting"
    exit 1
fi

echo "$(date) - Dumping out sorted agent nodes"
kubectl get no -L kubernetes.io/role -l kubernetes.io/role=agent --no-headers -o jsonpath="{.items[*].metadata.name}" | tr " " "\n" | sort

should_this_node_run_extension
if [[ $? -ne 0 ]]; then
    echo $(date) " - Not the first master node or the first agent node, no longer continuing extension. Exiting"
    exit 1
fi

ensure_k8s_namespace_exists $NAMESPACE

K8S_SECRET_NAME=dashboard-grafana
DS_TYPE=prometheus
DS_NAME=prometheus1

PROM_URL=http://monitoring-prometheus-server

# When RBAC is enabled, need to grant permission to tiller
kubectl create clusterrolebinding tiller-cluster-admin --clusterrole=cluster-admin --serviceaccount=kube-system:default

install_helm "$RAW_PROMETHEUS_CHART_VALS"
wait_for_tiller
if [[ $? -ne 0 ]]; then
    echo $(date) " - Tiller did not respond in a timely manner. Exiting"
    exit 1
fi
update_helm
install_cadvisor $NAMESPACE $CADVISOR_CONFIG_URL
install_prometheus $NAMESPACE
install_grafana $NAMESPACE

sleep 5

echo $(date) " - Creating the Prometheus datasource in Grafana"
GF_USER_NAME=$(kubectl get secret -n $NAMESPACE $K8S_SECRET_NAME -o jsonpath="{.data.grafana-admin-user}" | base64 --decode)
echo $GF_USER_NAME
GF_PASSWORD=$(kubectl get secret -n $NAMESPACE $K8S_SECRET_NAME -o jsonpath="{.data.grafana-admin-password}" | base64 --decode)
echo $GF_PASSWORD
GF_URL=$(kubectl get svc -n $NAMESPACE -l "app=dashboard-grafana,component=grafana" -o jsonpath="{.items[0].spec.clusterIP}")
echo $GF_URL

echo retrieving current data sources...
CURRENT_DS_LIST=$(curl -s --user "$GF_USER_NAME:$GF_PASSWORD" "$GF_URL/api/datasources")
echo $CURRENT_DS_LIST | grep -q "\"name\":\"$DS_NAME\""
if [[ $? -eq 0 ]]; then
    echo data source $DS_NAME already exists
    echo $CURRENT_DS_LIST | python -m json.tool
    exit 0
fi

echo data source $DS_NAME does not exist, creating...
DS_RAW=$(cat << EOF
{
    "name": "$DS_NAME",
    "type": "$DS_TYPE",
    "url": "$PROM_URL",
    "access": "proxy"
}
EOF
)



ATTEMPTS=90
SLEEP_TIME=10

ITERATION=0
while [[ $ITERATION -lt $ATTEMPTS ]]; do
    echo $(date) " - Is the grafana api running? (attempt $(( $ITERATION + 1 )) of $ATTEMPTS)"

    response=$(curl \
        -X POST \
        --user "$GF_USER_NAME:$GF_PASSWORD" \
        -H "Content-Type: application/json" \
        -d "$DS_RAW" \
        "$GF_URL/api/datasources")

    if [[ $response == *"Datasource added"* ]]; then
        echo $(date) " - Data source added successfully"
        break
    fi

    ITERATION=$(( $ITERATION + 1 ))
    sleep $SLEEP_TIME
done

echo $(date) " - Creating the Kubernetes dashboard in Grafana"

cat << EOF > sanitize_dashboard.py
#!/usr/bin/python3

import fileinput
import json

dashboard = json.loads(''.join(fileinput.input()))
dashboard.pop('__inputs')
dashboard.pop('__requires')
print(json.dumps(dashboard).replace('\${DS_PROMETHEUS}', 'prometheus1'))

EOF

chmod u+x sanitize_dashboard.py

DB_RAW=$(cat << EOF
{
    "dashboard": $(curl -sL "https://grafana.com/api/dashboards/3119/revisions/2/download" | ./sanitize_dashboard.py),
    "overwrite": false
}
EOF
)

curl \
    -X POST \
    --user "$GF_USER_NAME:$GF_PASSWORD" \
    -H "Content-Type: application/json" \
    -d "$DB_RAW" \
    "$GF_URL/api/dashboards/db"

echo $(date) " - Script complete"
