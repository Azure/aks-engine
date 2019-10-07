# This is a temporary file to test dot-sourcing functions stored in separate scripts in a zip file

filter Timestamp {"$(Get-Date -Format o): $_"}

function
Write-Log($message)
{
    $msg = $message | Timestamp
    Write-Output $msg
}

function DownloadFileOverHttp
{
    Param(
        [Parameter(Mandatory=$true)][string]
        $Url,
        [Parameter(Mandatory=$true)][string]
        $DestinationPath
    )

    # First check to see if a file with the same name is already cached on the VHD
    $fileName = [IO.Path]::GetFileName($Url)
    
    $search = @()
    if (Test-Path $global:CacheDir)
    {
        $search = [IO.Directory]::GetFiles($global:CacheDir, $fileName, [IO.SearchOption]::AllDirectories)
    }

    if ($search.Count -ne 0)
    {
        Write-Log "Using cached version of $fileName - Copying file from $($search[0]) to $DestinationPath"
        Move-Item -Path $search[0] -Destination $DestinationPath -Force
    }
    else 
    {
        $secureProtocols = @()
        $insecureProtocols = @([System.Net.SecurityProtocolType]::SystemDefault, [System.Net.SecurityProtocolType]::Ssl3)
    
        foreach ($protocol in [System.Enum]::GetValues([System.Net.SecurityProtocolType]))
        {
            if ($insecureProtocols -notcontains $protocol)
            {
                $secureProtocols += $protocol
            }
        }
        [System.Net.ServicePointManager]::SecurityProtocol = $secureProtocols
    
        $oldProgressPreference = $ProgressPreference
        $ProgressPreference = 'SilentlyContinue'
        Invoke-WebRequest $Url -UseBasicParsing -OutFile $DestinationPath -Verbose
        $ProgressPreference = $oldProgressPreference
        Write-Log "Downloaded file to $DestinationPath"
    }
}

# https://stackoverflow.com/a/34559554/697126
function New-TemporaryDirectory {
    $parent = [System.IO.Path]::GetTempPath()
    [string] $name = [System.Guid]::NewGuid()
    New-Item -ItemType Directory -Path (Join-Path $parent $name)
}

function Initialize-DataDirectories {
    # Some of the Kubernetes tests that were designed for Linux try to mount /tmp into a pod
    # On Windows, Go translates to c:\tmp. If that path doesn't exist, then some node tests fail

    $requiredPaths = 'c:\tmp'

    $requiredPaths | ForEach-Object {
        if (-Not (Test-Path $_)) {
            New-Item -ItemType Directory -Path $_
        }
    }
}

function Retry-Command
{
    Param(
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][string]
        $Command,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][hashtable]
        $Args,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][int]
        $Retries,
        [Parameter(Mandatory=$true)][ValidateNotNullOrEmpty()][int]
        $RetryDelaySeconds
    )

    for ($i = 0; $i -lt $Retries; $i++) {
        try {
            return & $Command @Args
        } catch {
            Start-Sleep $RetryDelaySeconds
        }
    }
}
