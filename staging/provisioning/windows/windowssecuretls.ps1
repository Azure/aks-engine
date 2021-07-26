#***************************************************************************************************************
# It does the following:
#   *   Disable TLS 1.O, TLS 1.1, SSLv2, SSLv3 and enables TLS1.2
#   *   The CipherSuite order is set to the SDL approved version.
#   *   The FIPS MinEncryptionLevel is set to 3.
#   *   RC4 is disabled
#***************************************************************************************************************

#******************* FUNCTION THAT ACTUALLY UPDATES KEYS ***********************
function Set-CryptoSetting {
    param ( 
        $regKeyName, 
        $value, 
        $valuedata, 
        $valuetype      
    ) 

    # Check for existence of registry key, and create if it does not exist 
    If (!(Test-Path -Path $regKeyName)) { 
        New-Item $regKeyName | Out-Null 
    } 

    # Get data of registry value, or null if it does not exist 
    $val = (Get-ItemProperty -Path $regKeyName -Name $value -ErrorAction SilentlyContinue).$value 


    If ($val -eq $null) { 
        # Value does not exist - create and set to desired value 
        New-ItemProperty -Path $regKeyName -Name $value -Value $valuedata -PropertyType $valuetype | Out-Null 
    }
    Else { 
        # Value does exist - if not equal to desired value, change it 
        If ($val -ne $valuedata) { 
            Set-ItemProperty -Path $regKeyName -Name $value -Value $valuedata 
        } 
    } 
}
#***************************************************************************************************************

#******************* FUNCTION THAT DISABLES RC4 ***********************
function DisableRC4 {
    $subkeys = Get-Item -Path "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL" 
    $ciphers = $subkeys.OpenSubKey("Ciphers", $true) 

    Write-Log "----- Checking the status of RC4 -----"

    $RC4 = $false
    if ($ciphers.SubKeyCount -eq 0) { 
        $k1 = $ciphers.CreateSubKey("RC4 128/128") 
        $k1.SetValue("Enabled", 0, [Microsoft.Win32.RegistryValueKind]::DWord) 
        $k2 = $ciphers.CreateSubKey("RC4 64/128") 
        $k2.SetValue("Enabled", 0, [Microsoft.Win32.RegistryValueKind]::DWord) 
        $k3 = $ciphers.CreateSubKey("RC4 56/128") 
        $k3.SetValue("Enabled", 0, [Microsoft.Win32.RegistryValueKind]::DWord) 
        $k4 = $ciphers.CreateSubKey("RC4 40/128") 
        $k4.SetValue("Enabled", 0, [Microsoft.Win32.RegistryValueKind]::DWord) 
        
        Write-Log "RC4 was disabled "
        $RC4 = $true
    } 

    If ($RC4 -ne $true) {
        Write-Log "There was no change for RC4 "
    }
}
#***************************************************************************************************************

#******************* FUNCTION CHECKS FOR PROBLEMATIC FIPS SETTING AND FIXES IT  ***********************
function Test-RegistryValueForFipsSettings {            
    $fipsPath = @( 
        "HKLM:\System\CurrentControlSet\Control\Terminal Server\WinStations\RDP-Tcp",
        "HKLM:\SOFTWARE\Policies\Microsoft\Windows NT\Terminal Services",
        "HKLM:\System\CurrentControlSet\Control\Terminal Server\DefaultUserConfiguration"
    )
    
    $fipsValue = "MinEncryptionLevel"
    
    foreach ($path in $fipsPath) {
        Write-Log "Checking to see if $($path)\$fipsValue exists"

        $ErrorActionPreference = "stop"
        Try {
            $result = Get-ItemProperty -Path $path | Select-Object -ExpandProperty $fipsValue
            if ($result -eq 4) {
                set-itemproperty -Path $path -Name $fipsValue -value 3
                Write-Log "Regkey $($path)\$fipsValue was changed from value $result to a value of 3"
            }
            else {
                Write-Log "Regkey $($path)\$fipsValue left at value $result"
            }
    
        }
        Catch [System.Management.Automation.ItemNotFoundException] {
            Write-Log "Reg path $path was not found"
        }
        Catch [System.Management.Automation.PSArgumentException] {
            Write-Log "Regkey $($path)\$fipsValue was not found"
        }
        Catch {
            Write-Log "Error of type $($Error[0].Exception.GetType().FullName) trying to get $($path)\$fipsValue"
        }
        Finally {
            $ErrorActionPreference = "Continue"
        }
    }
}
#***************************************************************************************************************

#********************************TLS CipherSuite Settings *******************************************

# CipherSuites for windows OS < 10
function Get-BaseCipherSuitesOlderWindows()
{
    param
    (
        [Parameter(Mandatory=$true, Position=0)][bool] $isExcellenceOrder
    )
    $cipherorder = @()

    if ($isExcellenceOrder -eq $true)
    {
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384_P384"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256_P256"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384_P384"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256_P256"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384_P384"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256_P256"
    }
    else
    {
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256_P256"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384_P384"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256_P256"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384_P384"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256_P256"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384_P384"
    }

    # Add additional ciphers when EnableOlderTlsVersions flag is set to true
    if ($EnableOlderTlsVersions)
    {
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA_P256"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA_P256"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA_P256"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA_P256"
        $cipherorder += "TLS_RSA_WITH_AES_256_GCM_SHA384" 
        $cipherorder += "TLS_RSA_WITH_AES_128_GCM_SHA256" 
        $cipherorder += "TLS_RSA_WITH_AES_256_CBC_SHA256" 
        $cipherorder += "TLS_RSA_WITH_AES_128_CBC_SHA256" 
        $cipherorder += "TLS_RSA_WITH_AES_256_CBC_SHA"
        $cipherorder += "TLS_RSA_WITH_AES_128_CBC_SHA"
    }
    return $cipherorder
}

# Ciphersuites needed for backwards compatibility with Firefox, Chrome
# Server 2012 R2 doesn't support TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
# Both firefox and chrome negotiate ECDHE_RSA_AES_256_CBC_SHA1, Edge negotiates ECDHE_RSA_AES_256_CBC_SHA384
function Get-BrowserCompatCipherSuitesOlderWindows()
{
    param
    (
        [Parameter(Mandatory=$true, Position=0)][bool] $isExcellenceOrder
    )
    $cipherorder = @()

    if ($isExcellenceOrder -eq $true)
    {
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA_P384"  # (uses SHA-1)  
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA_P256"  # (uses SHA-1)
    }
    else
    {
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA_P256"  # (uses SHA-1)
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA_P384"  # (uses SHA-1)  
    }
    return $cipherorder
}

# Ciphersuites for OS versions windows 10 and above
function Get-BaseCipherSuitesWin10Above()
{
    param
    (
        [Parameter(Mandatory=$true, Position=0)][bool] $isExcellenceOrder
    )

    $cipherorder = @()
    if ($isExcellenceOrder -eq $true)
    {
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256"
    }
    else
    {
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384"
    }
    # Add additional ciphers when EnableOlderTlsVersions flag is set to true
    if ($EnableOlderTlsVersions)
    {
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA_P256"
        $cipherorder += "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA_P256"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA_P256"
        $cipherorder += "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA_P256"
        $cipherorder += "TLS_RSA_WITH_AES_256_GCM_SHA384" 
        $cipherorder += "TLS_RSA_WITH_AES_128_GCM_SHA256" 
        $cipherorder += "TLS_RSA_WITH_AES_256_CBC_SHA256" 
        $cipherorder += "TLS_RSA_WITH_AES_128_CBC_SHA256" 
        $cipherorder += "TLS_RSA_WITH_AES_256_CBC_SHA"
        $cipherorder += "TLS_RSA_WITH_AES_128_CBC_SHA"
    }

    return $cipherorder
}

#******************************* TLS Version Settings ****************************************************

function Get-RegKeyPathForTls12()
{
    $regKeyPath = @(
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\TLS 1.2",        
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\TLS 1.2\Client", 
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\TLS 1.2\Server" 
    )
    return $regKeyPath
}

function Get-RegKeyPathForTls11()
{
    $regKeyPath = @(
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\TLS 1.1", 
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\TLS 1.1\Client",
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\TLS 1.1\Server" 
    )
    return $regKeyPath
}

function Get-RegKeypathForTls10()
{
    $regKeyPath = @(
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\TLS 1.0", 
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\TLS 1.0\Client", 
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\TLS 1.0\Server"
    )
    return $regKeyPath
}

function Get-RegKeyPathForSsl30()
{
    $regKeyPath = @(
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\SSL 3.0",        
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\SSL 3.0\Client", 
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\SSL 3.0\Server"
    )
    return $regKeyPath
}

function Get-RegKeyPathForSsl20()
{
    $regKeyPath = @(
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\SSL 2.0",
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\SSL 2.0\Client",  
        "HKLM:\SYSTEM\CurrentControlSet\Control\SecurityProviders\SCHANNEL\Protocols\SSL 2.0\Server"
    )
    return $regKeyPath
}

#******************************* Wrap TLSSettings.ps1 to a function without restart **********************

function Enable-SecureTls {
    $registryPathGoodGuys = @()
    $registryPathBadGuys = @()

    Write-Log -Message "========== Start of logging for Enable-SecureTls =========="
    # we enable TLS 1.2 and disable others
    $registryPathGoodGuys += Get-RegKeyPathForTls12

    $registryPathBadGuys += Get-RegKeyPathForSsl20
    $registryPathBadGuys += Get-RegKeyPathForSsl30
    $registryPathBadGuys += Get-RegKeypathForTls10
    $registryPathBadGuys += Get-RegKeyPathForTls11
    Write-Log "Enabling TLS1.2. Disabling TLS1.1, TLS1.0, SSL3.0, SSL2.0"

    Write-Log "Check which registry keys exist already and which registry keys need to be created." 

    #******************* CREATE THE REGISTRY KEYS IF THEY DON'T EXIST********************************
    # Check for existence of GoodGuy registry keys, and create if they do not exist 
    For ($i = 0; $i -lt $registryPathGoodGuys.Length; $i = $i + 1) { 
        Write-Log "Checking for existing of key: $($registryPathGoodGuys[$i]) "
        If (!(Test-Path -Path $registryPathGoodGuys[$i])) { 
            New-Item $registryPathGoodGuys[$i] | Out-Null
            Write-Log "Creating key: $($registryPathGoodGuys[$i]) "
        }
    } 
    
    # Check for existence of BadGuy registry keys, and create if they do not exist 
    For ($i = 0; $i -lt $registryPathBadGuys.Length; $i = $i + 1) { 
        Write-Log "Checking for existing of key: $($registryPathBadGuys[$i]) "
        If (!(Test-Path -Path $registryPathBadGuys[$i])) { 
            Write-Log "Creating key: $($registryPathBadGuys[$i]) "
            New-Item  $registryPathBadGuys[$i] | Out-Null
        }
    }

    #******************* EXPLICITLY DISABLE SSLV2, SSLV3, TLS10 AND TLS11 ********************************
    For ($i = 0; $i -lt $registryPathBadGuys.Length; $i = $i + 1) {
        if ($registryPathBadGuys[$i].Contains("Client") -Or $registryPathBadGuys[$i].Contains("Server")) {
            Write-Log "Disabling this key: $($registryPathBadGuys[$i]) "
            Set-CryptoSetting $registryPathBadGuys[$i].ToString() Enabled 0 DWord  
            Set-CryptoSetting $registryPathBadGuys[$i].ToString() DisabledByDefault 1 DWord  
        }
    }

    #********************************* EXPLICITLY Enable TLS12 ****************************************
    For ($i = 0; $i -lt $registryPathGoodGuys.Length; $i = $i + 1) {
        if ($registryPathGoodGuys[$i].Contains("Client") -Or $registryPathGoodGuys[$i].Contains("Server")) {
            Write-Log "Enabling this key: $($registryPathGoodGuys[$i]) " 
            Set-CryptoSetting $registryPathGoodGuys[$i].ToString() Enabled 1 DWord  
            Set-CryptoSetting $registryPathGoodGuys[$i].ToString() DisabledByDefault 0 DWord 
        }
    }

    #************************************** Disable RC4 ************************************************
    DisableRC4
        
    #************************************** Set Cipher Suite Order **************************************
    Write-Log "----- starting ciphersuite order calculation for excellence cipher suite order -----" 
    $cipherlist = @()

    if ([Environment]::OSVersion.Version.Major -lt 10) 
    {
        $cipherlist += Get-BaseCipherSuitesOlderWindows -isExcellenceOrder $true
        $cipherlist += Get-BrowserCompatCipherSuitesOlderWindows -isExcellenceOrder $true
    }
    else
    {
        $cipherlist += Get-BaseCipherSuitesWin10Above -isExcellenceOrder $true
    }
    $cipherorder = [System.String]::Join(",", $cipherlist)
    Write-Log "Appropriate ciphersuite order : $cipherorder"
    
    $CipherSuiteRegKey = "HKLM:\SOFTWARE\Policies\Microsoft\Cryptography\Configuration\SSL\00010002" 
    
    if (!(Test-Path -Path $CipherSuiteRegKey)) 
    { 
        New-Item $CipherSuiteRegKey | Out-Null 
        Write-Log "Creating key: $($CipherSuiteRegKey) "
    } 
    
    $val = (Get-Item -Path $CipherSuiteRegKey -ErrorAction SilentlyContinue).GetValue("Functions", $null)
    Write-Log "Previous cipher suite value: $val  "
    Write-Log "New cipher suite value     : $cipherorder  "		 
        
    if ($val -ne $cipherorder) 
    { 
        Write-Log "Cipher suite order needs to be updated. "
        Set-ItemProperty -Path $CipherSuiteRegKey -Name Functions -Value $cipherorder 
        Write-Log "Cipher suite value was updated. "
    }
    else
    {
        Write-Log "Cipher suite order does not need to be updated. "
        Write-Log "Cipher suite value was not updated as there was no change. "
    }
            
    #****************************** CHECK THE FIPS SETTING WHICH IMPACTS RDP'S ALLOWED CIPHERS **************************
    #Check for FipsSettings
    Write-Log "Checking to see if reg keys exist and if MinEncryptionLevel is set to 4"
    Test-RegistryValueForFipsSettings

    Write-Log -Message "========== End of logging for Enable-SecureTls =========="
}