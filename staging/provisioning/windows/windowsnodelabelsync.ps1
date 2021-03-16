<#
.DESCRIPTION
    This script is intended to sync customized and system-assigned Kubelet node labels.
#>

$Global:ClusterConfiguration = ConvertFrom-Json ((Get-Content "c:\k\kubeclusterconfig.json" -ErrorAction Stop) | out-string)
$Global:KubeletNodeLabels = $Global:ClusterConfiguration.Kubernetes.Kubelet.NodeLabels

$NodeLabels = $KubeletNodeLabels -split ","
ForEach ($NodeLabel in $NodeLabels) {
  c:\k\kubectl.exe --kubeconfig=c:\k\config label node --overwrite $env:computername.ToLower() $NodeLabel
}
