param(
  [string]$BinaryPath = "C:\Program Files\pulsed\pulsed.exe",
  [string]$Http = ":8080",
  [string]$AdvertiseHttp = "http://CHANGE_ME:8080",
  [string]$Gossip = ":7946",
  [string]$Seeds = "",
  [string]$Db = "C:\ProgramData\pulsed",
  [string]$NssmPath = "nssm.exe"
)

$ErrorActionPreference = "Stop"

New-Item -ItemType Directory -Force -Path (Split-Path $BinaryPath) | Out-Null
New-Item -ItemType Directory -Force -Path $Db | Out-Null

if (-not (Get-Command $NssmPath -ErrorAction SilentlyContinue)) {
  throw "nssm.exe is required to run pulsed as a Windows service. Install NSSM or pass -NssmPath."
}

& $NssmPath install pulsed $BinaryPath
& $NssmPath set pulsed AppDirectory $Db
& $NssmPath set pulsed AppEnvironmentExtra `
  "PULSED_HTTP=$Http" `
  "PULSED_ADVERTISE_HTTP=$AdvertiseHttp" `
  "PULSED_GOSSIP=$Gossip" `
  "PULSED_SEEDS=$Seeds" `
  "PULSED_DB=$Db" `
  "PULSED_WEB=true"
& $NssmPath set pulsed Start SERVICE_AUTO_START
& $NssmPath set pulsed AppStdout "$Db\pulsed.log"
& $NssmPath set pulsed AppStderr "$Db\pulsed.err.log"
& $NssmPath start pulsed
