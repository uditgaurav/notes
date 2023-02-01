param(
  [AllowEmptyString()]
  [Int32]$Duration=60,
  [string]$DestinationHosts,
  [string]$IPAddresses
  )
set-strictmode -version 2.0

Write-Host "Schedule job to delete the DNS rules that will be added later"
$refjob = Start-Job -ScriptBlock $chaosRevertjob
if($DestinationHosts) {
  $DestHosts = "$DestinationHosts".Split(",")
  $addresses = New-Object Collections.Generic.List[String]
  foreach ($DestinationHost in $DestHosts) {
    $addresses.Add("$($DestinationHost)")
  }
}else {
  $IPAddressList = "$IPAddresses".Split(",")
  $addresses = New-Object Collections.Generic.List[String]
  foreach ($IPAddress in $IPAddressList) {
    $addresses.Add("$($IPAddress)")
  }
}

foreach ($address in $addresses) {
  $addressips = Resolve-DnsName -Name $address -Type A -DnsOnly
  foreach($addressip in $addressips.IPAddress){          
    New-NetFirewallRule -DisplayName "CHAOS Block VM IP address" -Direction Outbound –LocalPort Any -Protocol UDP -Action Block -RemoteAddress $addressip | Out-Null
    Write-Host "Added $($addressip) for $($address) to the Firewall and Blocked for UDP"
    New-NetFirewallRule -DisplayName "CHAOS Block VM IP address" -Direction Outbound –LocalPort Any -Protocol TCP -Action Block -RemoteAddress $addressip | Out-Null
    Write-Host "Added $($addressip) for $($address) to the Firewall and Blocked for TCP"
  }
}
Write-Host "Wating for chaos duration of $Duration seconds"
Start-Sleep -s ($Duration)
Write-Host "About to stop jobs"
Remove-NetFirewallRule -DisplayName "CHAOS Block VM IP address"
Write-Host "Chaos Completed!!!"
