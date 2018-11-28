
# Return codes:
#  0 - success
#  1 - install failure
#  2 - download failure
#  3 - unrecognized patch extension

param(
    [string[]] $URIs
)

function DownloadFile([string] $URI, [string] $fullName)
{
    try {
        Write-Host "Downloading $URI"
        $ProgressPreference = 'SilentlyContinue'
        Invoke-WebRequest -UseBasicParsing $URI -OutFile $fullName
    } catch {
        Write-Error $_
        exit 2
    }
}


$URIs | ForEach-Object {
    Write-Host "Processing $_"
    $uri = $_
    $pathOnly = $uri
    if ($pathOnly.Contains("?"))
    {
        $pathOnly = $pathOnly.Split("?")[0]
    }
    $fileName = Split-Path $pathOnly -Leaf
    $ext = [io.path]::GetExtension($fileName)
    $fullName = [io.path]::Combine($env:TEMP, $fileName)
    switch ($ext) {
        ".exe" {
            Start-Process -FilePath bcdedit.exe -ArgumentList "/set {current} testsigning on" -Wait
            DownloadFile -URI $uri -fullName $fullName
            Write-Host "Starting $fullName"
            $proc = Start-Process -Passthru -FilePath "$fullName" -ArgumentList "/q /norestart"
            Wait-Process -InputObject $proc
            switch ($proc.ExitCode)
            {
                0 {
                    Write-Host "Finished running $fullName"
                }
                3010 {
                    Write-Host "Finished running $fullName. Reboot required to finish patching."
                }
                Default {
                    Write-Error "Error running $fullName, exitcode $($proc.ExitCode)"
                    exit 1
                }
            }
        }
        ".msu" {
            DownloadFile -URI $uri -fullName $fullName
            Write-Host "Installing $localPath"
            $proc = Start-Process -Passthru -FilePath wusa.exe -ArgumentList "$fullName /quiet /norestart"
            Wait-Process -InputObject $proc
            switch ($proc.ExitCode)
            {
                0 {
                    Write-Host "Finished running $fullName"
                }
                3010 {
                    Write-Host "Finished running $fullName. Reboot required to finish patching."
                }
                Default {
                    Write-Error "Error running $fullName, exitcode $($proc.ExitCode)"
                    exit 1
                }
            }
        }
        Default {
            Write-Error "This script extension doesn't know how to install $ext files"
            exit 3
        }
    }
}

# No failures, schedule reboot now

schtasks /create /TN RebootAfterPatch /RU SYSTEM /TR "shutdown.exe /r /t 0 /d 2:17" /SC ONCE /ST $(([System.DateTime]::Now + [timespan]::FromMinutes(5)).ToString("HH:mm")) /V1 /Z
exit 0