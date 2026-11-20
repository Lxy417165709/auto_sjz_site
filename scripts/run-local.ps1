param(
  [int]$Port = 8080,
  [string]$Image = "auto_sjz_site:local",
  [string]$Container = "auto_sjz_site"
)

$ErrorActionPreference = "Stop"
$root = Split-Path -Parent $PSScriptRoot
Push-Location $root

try {
  Write-Host "[1/3] building image $Image ..."
  if ($env:NPM_REGISTRY) {
    docker build --build-arg "NPM_REGISTRY=$env:NPM_REGISTRY" -t $Image .
  } else {
    docker build -t $Image .
  }

  $dataDir = Join-Path $root "data"
  New-Item -ItemType Directory -Force -Path $dataDir | Out-Null

  Write-Host "[2/3] starting container $Container on port $Port ..."
  docker rm -f $Container 2>$null | Out-Null
  docker run -d --name $Container --restart unless-stopped `
    -p "${Port}:8080" `
    -e "SITE_BASE_URL=http://localhost:${Port}" `
    -e "MANAGER_URL=$env:MANAGER_URL" `
    -v "${dataDir}:/data" `
    $Image | Out-Null

  Start-Sleep -Seconds 2
  Write-Host "[3/3] logs:"
  docker logs --tail 20 $Container
  Write-Host "site ready: http://localhost:$Port"
} finally {
  Pop-Location
}
