<#
.DESCRIPTION
    This script is intended to update Kubelet node labels.
#>

$global:KubeletNodeLabels = $Global:ClusterConfiguration.Kubernetes.Kubelet.NodeLabels

$NodeLabels = $KubeletNodeLabels -split ","
ForEach ($NodeLabel in $NodeLabels) {
  c:\k\kubectl.exe --kubeconfig='c:\\k\\config' --overwrite $env:computername $HOSTNAME $NodeLabel
}