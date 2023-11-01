# Set the execution policy to allow the script to run (might need to be run as an administrator)
Set-ExecutionPolicy Bypass -Scope Process -Force

# Define the tools and their download URLs
$tools = @(
    @{ 
        Name = "clumsy";
        Url = "https://github.com/jagt/clumsy/releases/download/0.3/clumsy-0.3-win64-a.zip";
        Destination = "$env:USERPROFILE\Downloads\clumsy.zip"
    },
    @{
        Name = "diskpd";
        Url = "https://github.com/microsoft/diskspd/releases/download/v2.1/DiskSpd.ZIP";
        Destination = "$env:USERPROFILE\Downloads\diskspd.zip"
    },
    @{
        Name = "Testlimit";
        Url = "https://download.sysinternals.com/files/Testlimit.zip";
        Destination = "$env:USERPROFILE\Downloads\testlimit.zip"
    }
)

# Download and extract each tool
foreach ($tool in $tools) {
    Write-Host ("Downloading {0}..." -f $tool.Name)
    Invoke-WebRequest -Uri $tool.Url -OutFile $tool.Destination

    Write-Host ("Extracting {0}..." -f $tool.Name)
    Expand-Archive -Path $tool.Destination -DestinationPath "$env:USERPROFILE\Downloads\$($tool.Name)"
}

Write-Host "All tools have been downloaded and extracted."

# Reset the execution policy to its default (might need to be run as an administrator)
Set-ExecutionPolicy Default -Scope Process -Force
