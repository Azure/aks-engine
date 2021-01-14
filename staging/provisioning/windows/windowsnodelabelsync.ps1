<#
.DESCRIPTION
    This script is intended to update Kubelet node labels.
#>

$Global:ClusterConfiguration = ConvertFrom-Json ((Get-Content "c:\k\kubeclusterconfig.json" -ErrorAction Stop) | out-string)
$global:KubeletNodeLabels = $Global:ClusterConfiguration.Kubernetes.Kubelet.NodeLabels

$NodeLabels = $KubeletNodeLabels -split ","
ForEach ($NodeLabel in $NodeLabels) {
  c:\k\kubectl.exe --kubeconfig=c:\k\config --overwrite $env:computername.ToLower() $NodeLabel
}