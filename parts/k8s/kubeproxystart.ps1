$Global:ClusterConfiguration = ConvertFrom-Json ((Get-Content "c:\k\kubeclusterconfig.json" -ErrorAction Stop) | out-string)

$KubeNetwork = "azure"
if ($Global:ClusterConfiguration.Cni.Name -eq "kubenet") {
    $KubeNetwork = "l2bridge"
}

$env:KUBE_NETWORK = $KubeNetwork
$global:HNSModule = "c:\k\hns.psm1"
$KubeDir = $Global:ClusterConfiguration.Install.Destination

$hnsNetwork = Get-HnsNetwork | ? Name -EQ $KubeNetwork
while (!$hnsNetwork) {
    Write-Host "$(Get-Date -Format o) Waiting for Network [$KubeNetwork] to be created . . ."
    Start-Sleep 10
    $hnsNetwork = Get-HnsNetwork | ? Name -EQ $KubeNetwork
}

#
# cleanup the persisted policy lists
#
Import-Module $global:HNSModule
# Workaround for https://github.com/kubernetes/kubernetes/pull/68923 in < 1.14,
# and https://github.com/kubernetes/kubernetes/pull/78612 for <= 1.15
Get-HnsPolicyList | Remove-HnsPolicyList

if (("--feature-gates=IPv6DualStack=true" | ? { $Global:ClusterConfiguration.Kubernetes.Kubelet.ConfigArgs -match $_ }) -ne $null) {
    .$KubeDir\kube-proxy.exe --v=3 --proxy-mode=kernelspace --feature-gates=IPv6DualStack=true --hostname-override=$env:computername --kubeconfig=$KubeDir\config
}
else {
    .$KubeDir\kube-proxy.exe --v=3 --proxy-mode=kernelspace --hostname-override=$env:computername --kubeconfig=$KubeDir\config
}