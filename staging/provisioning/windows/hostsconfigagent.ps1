$Global:ClusterConfiguration = ConvertFrom-Json ((Get-Content "c:\k\kubeclusterconfig.json" -ErrorAction Stop) | out-string)
$clusterFQDN = $Global:ClusterConfiguration.Kubernetes.ControlPlane.IpAddress
$hostsFile="C:\Windows\System32\drivers\etc\hosts"
$retryDelaySeconds = 15

filter Timestamp { "$(Get-Date -Format o): $_" }

function Write-Log($message) {
    $msg = $message | Timestamp
    Write-Output $msg
}

function Retry-Command {
    Param(
        [Parameter(Mandatory = $true)][ValidateNotNullOrEmpty()][string]
        $Command,
        [Parameter(Mandatory = $true)][ValidateNotNullOrEmpty()][hashtable]
        $Args,
        [Parameter(Mandatory = $true)][ValidateNotNullOrEmpty()][int]
        $Retries,
        [Parameter(Mandatory = $true)][ValidateNotNullOrEmpty()][int]
        $RetryDelaySeconds
    )

    for ($i = 0; $i -lt $Retries; $i++) {
        try {
            return & $Command @Args
        }
        catch {
            Start-Sleep $RetryDelaySeconds
        }
    }
}

function Get-APIServer-IPAddress
{
    $uri = "http://169.254.169.254/metadata/instance/compute/tags?api-version=2019-03-11&format=text"
    $response = Retry-Command -Command "Invoke-RestMethod" -Args @{Uri=$uri; Method="Get"; ContentType="application/json"; Headers=@{"Metadata"="true"}} -Retries 3 -RetryDelaySeconds 5

    if(!$response) {
        return ""
    }

    foreach ($tag in $response.Split(";"))
    {
        $values = $tag.Split(":")
        if ($values.Length -ne 2)
        {
            return ""
        }

        if ($values[0] -eq "aksAPIServerIPAddress")
        {
            return $values[1]
        }
    }

    return ""
}

Write-Log "Get cluster APIServer FQDN: $clusterFQDN"
while ($true)
{
    $clusterIP = Get-APIServer-IPAddress
    if ($clusterIP -eq "") {
        Start-Sleep $retryDelaySeconds
        continue
    }

    $hostsContent=Get-Content -Path $hostsFile -Encoding UTF8
    if ($hostsContent -match "$clusterIP $clusterFQDN") {
        Write-Log "APIServer FQDN has already been set to $clusterIP in hosts file"
    } else {
        $hostsContent -notmatch "$clusterFQDN" | Out-File $hostsFile -Encoding UTF8
        Add-Content -Path $hostsFile -Value "$clusterIP $clusterFQDN" -Encoding UTF8
        Write-Log "Updated APIServer FQDN to $clusterIP in hosts file"
    }

    Start-Sleep $retryDelaySeconds
}
