# Set the execution policy to allow the script to run (might need to be run as an administrator)
Set-ExecutionPolicy Bypass -Scope Process -Force

# Define the tools and their download URLs
$tools = @(
    @{
        Name = "clumsy";
        Url = "https://github.com/jagt/clumsy/releases/download/0.3/clumsy-0.3-win64-a.zip";
        Destination = "$env:USERPROFILE\Downloads\clumsy.zip"
        BinPath = "$env:USERPROFILE\Downloads\clumsy"
    },
    @{
        Name = "diskpd";
        Url = "https://github.com/microsoft/diskspd/releases/download/v2.1/DiskSpd.ZIP";
        Destination = "$env:USERPROFILE\Downloads\diskspd.zip"
        BinPath = "$env:USERPROFILE\Downloads\diskspd"
    },
    @{
        Name = "Testlimit";
        Url = "https://download.sysinternals.com/files/Testlimit.zip";
        Destination = "$env:USERPROFILE\Downloads\testlimit.zip"
        BinPath = "$env:USERPROFILE\Downloads\Testlimit"
    }
)

# Download and extract each tool and add its binary path to the PATH environment variable
foreach ($tool in $tools) {
    Write-Host ("Downloading {0}..." -f $tool.Name)
    Invoke-WebRequest -Uri $tool.Url -OutFile $tool.Destination

    Write-Host ("Extracting {0}..." -f $tool.Name)
    Expand-Archive -Path $tool.Destination -DestinationPath $tool.BinPath

    # Check if the binary path is already in the PATH environment variable
    $currentPath = [Environment]::GetEnvironmentVariable("PATH", [EnvironmentVariableTarget]::User)
    if (-not ($currentPath -like "*$($tool.BinPath)*")) {
        Write-Host ("Adding {0} to PATH..." -f $tool.Name)
        $newPath = $currentPath + ";" + $tool.BinPath
        [Environment]::SetEnvironmentVariable("PATH", $newPath, [EnvironmentVariableTarget]::User)
    } else {
        Write-Host ("{0} is already in PATH." -f $tool.Name)
    }
}

Write-Host "All tools have been downloaded, extracted, and their paths have been added to the PATH environment variable."

# Reset the execution policy to its default (might need to be run as an administrator)
Set-ExecutionPolicy Default -Scope Process -Force
