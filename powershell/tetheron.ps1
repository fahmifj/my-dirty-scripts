# https://stackoverflow.com/questions/57899096/tetheringwifiband-in-windows-networking-networkoperators

$connProfile = [Windows.Networking.Connectivity.NetworkInformation,Windows.Networking.Connectivity,ContentType=WindowsRuntime]::GetInternetConnectionProfile()
if ($connProfile -eq $null){
    # Write-Host "No internet connection available"
    return 
}
Write-Host $connProfile.ProfileName

$tetherMgr = [Windows.Networking.NetworkOperators.NetworkOperatorTetheringManager,Windows.Networking.NetworkOperators,ContentType=WindowsRuntime]::CreateFromConnectionProfile($connProfile)
$hotspotStatus = $tetherMgr.TetheringOperationalState
Write-Host "Hotspot Status: " $hotspotStatus

if ($hotspotStatus -eq "On") {
    Write-Host "Hotspot already active"
    return
}

$tetherMgr.StartTetheringAsync()

#([Windows.Networking.NetworkOperators.NetworkOperatorTetheringOperationResult])
