# Stop and remove Azure Agents to enable use in Azure Stack
# If deploying an Azure VM the agents will be re-added to the VMs at deployment time
Stop-Service WindowsAzureGuestAgent
Stop-Service WindowsAzureNetAgentSvc
Stop-Service RdAgent
& sc.exe delete WindowsAzureGuestAgent
& sc.exe delete WindowsAzureNetAgentSvc
& sc.exe delete RdAgent

# Remove the WindowsAzureGuestAgent registry key for sysprep 
# This removes AzureGuestAgent from participating in sysprep 
# There was an update that is missing VMAgentDisabler.dll
$path = "Registry::HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion\Setup\SysPrepExternal\Generalize"
$generalizeKey = Get-Item -Path $path
$generalizeProperties = $generalizeKey | Select-Object -ExpandProperty property
$values = $generalizeProperties | ForEach-Object {
    New-Object psobject -Property @{"Name"=$_;
    "Value" = (Get-ItemProperty -Path $path -Name $_).$_}
}

$values | ForEach-Object {
    $item = $_;
    if( $item.Value.Contains("VMAgentDisabler.dll")) {
            Write-HOST "Removing " $item.Name - $item.Value;
            Remove-ItemProperty -Path $path -Name $item.Name;
    }
}

# run Sysprep
if( Test-Path $Env:SystemRoot\\system32\\Sysprep\\unattend.xml ) {  Remove-Item $Env:SystemRoot\\system32\\Sysprep\\unattend.xml -Force }
& $env:SystemRoot\\System32\\Sysprep\\Sysprep.exe /oobe /generalize /mode:vm /quiet /quit

# when done clean up
while($true) { $imageState = Get-ItemProperty HKLM:\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Setup\\State | Select ImageState; if($imageState.ImageState -ne 'IMAGE_STATE_GENERALIZE_RESEAL_TO_OOBE') { Write-Output $imageState.ImageState; Start-Sleep -s 10  } else { break } }
Get-ChildItem c:\\WindowsAzure -Force | Sort-Object -Property FullName -Descending | ForEach-Object { try { Remove-Item -Path $_.FullName -Force -Recurse -ErrorAction SilentlyContinue; } catch { } }
Remove-Item -Path WSMan:\\Localhost\\listener\\listener* -Recurse