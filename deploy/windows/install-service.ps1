param(
  [string]$BinaryPath = "C:\Program Files\psstd\psstd.exe",
  [string]$Http = ":8080",
  [string]$AdvertiseHttp = "http://CHANGE_ME:8080",
  [string]$Gossip = ":7946",
  [string]$Seeds = "",
  [string]$Db = "C:\ProgramData\psstd",
  [string]$NssmPath = "nssm.exe"
)

$ErrorActionPreference = "Stop"

New-Item -ItemType Directory -Force -Path (Split-Path $BinaryPath) | Out-Null
New-Item -ItemType Directory -Force -Path $Db | Out-Null

if (-not (Get-Command $NssmPath -ErrorAction SilentlyContinue)) {
  throw "nssm.exe is required to run psstd as a Windows service. Install NSSM or pass -NssmPath."
}

& $NssmPath install psstd $BinaryPath
& $NssmPath set psstd AppDirectory $Db
& $NssmPath set psstd AppEnvironmentExtra `
  "PSSTD_HTTP=$Http" `
  "PSSTD_ADVERTISE_HTTP=$AdvertiseHttp" `
  "PSSTD_GOSSIP=$Gossip" `
  "PSSTD_SEEDS=$Seeds" `
  "PSSTD_DB=$Db" `
  "PSSTD_WEB=true"
& $NssmPath set psstd Start SERVICE_AUTO_START
& $NssmPath set psstd AppStdout "$Db\psstd.log"
& $NssmPath set psstd AppStderr "$Db\psstd.err.log"
& $NssmPath start psstd
