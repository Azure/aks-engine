$Global:ClusterConfiguration = ConvertFrom-Json ((Get-Content "c:\k\kubeclusterconfig.json" -ErrorAction Stop) | out-string)
$KubeproxyFeatureGates = $Global:ClusterConfiguration.Kubernetes.Kubeproxy.FeatureGates # This is the initial feature list passed in from aks-engine

$KubeNetwork = "azure"
if ($Global:ClusterConfiguration.Cni.Name -eq "kubenet") {
    $KubeNetwork = "l2bridge"
}

$env:KUBE_NETWORK = $KubeNetwork
$global:HNSModule = "c:\k\hns.psm1"
$global:KubeDir = $Global:ClusterConfiguration.Install.Destination
$global:KubeproxyArgList = @("--v=3", "--proxy-mode=kernelspace", "--hostname-override=$env:computername", "--kubeconfig=$KubeDir\config")

$hnsNetwork = Get-HnsNetwork | ? Name -EQ $KubeNetwork
while (!$hnsNetwork) {
    Write-Host "$(Get-Date -Format o) Waiting for Network [$KubeNetwork] to be created . . ."
    Start-Sleep 10
    $hnsNetwork = Get-HnsNetwork | ? Name -EQ $KubeNetwork
}

# enable WinDsr if WinDsr feature gate is enabled
if ($KubeproxyFeatureGates -contains "WinDSR=true") {
    $global:KubeproxyArgList += @("--enable-dsr=true")
}

if ($KubeproxyFeatureGates.Count -ne 0) {
    $global:KubeproxyArgList += @("--feature-gates='" + ($KubeproxyFeatureGates -join "','") + "'")
}

#
# cleanup the persisted policy lists
#
Import-Module $global:HNSModule
# Workaround for https://github.com/kubernetes/kubernetes/pull/68923 in < 1.14,
# and https://github.com/kubernetes/kubernetes/pull/78612 for <= 1.15
Get-HnsPolicyList | Remove-HnsPolicyList

# Use run-process.cs to set process priority class as 'AboveNormal'
# Load a signed version of runprocess.dll if it exists for Azure SysLock compliance
# otherwise load class from cs file (for CI/testing)
if (Test-Path "$global:KubeDir\runprocess.dll") {
    [System.Reflection.Assembly]::LoadFrom("$global:KubeDir\runprocess.dll")
} else {
    Add-Type -Path "$global:KubeDir\run-process.cs"
}
$exe = "$global:KubeDir\kube-proxy.exe"
$args = ($global:KubeproxyArgList -join " ")
[RunProcess.exec]::RunProcess($exe, $args, [System.Diagnostics.ProcessPriorityClass]::AboveNormal)
