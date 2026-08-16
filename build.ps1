# Remove old binaries if they exist
if (Test-Path "bin") {
    Remove-Item -Recurse -Force "bin"
}

# Create the bin directory
New-Item -ItemType Directory -Path "bin" | Out-Null

# Define the targets
$targets = @("linux/amd64", "linux/arm64", "windows/amd64", "darwin/amd64", "darwin/arm64")

foreach ($target in $targets) {
    $os, $arch = $target.Split('/')
    
    $env:GOOS = $os
    $env:GOARCH = $arch
    $env:CGO_ENABLED = "0"
    
    # Determine the extension based on the OS
    $extension = ""
    if ($os -eq "windows") {
        $extension = ".exe"
    }
    
    $output = "bin/installment-schedule-$os-$arch$extension"
    
    Write-Host "Building for $os/$arch..."
    
    # Use '.' if your main.go is in the root, or './cmd' if it's in a cmd folder
    go build -o $output ./cmd
}

Write-Host "Build complete!" -ForegroundColor Green