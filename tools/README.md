# Preparing EC2 Windows Instance for Harness Chaos Engineering

## Overview

This script ensures that the EC2 Windows instance is primed for Harness Chaos Engineering. The steps involve verifying that specific processes and services run with administrator privileges, and ensuring that the Windows Firewall is enabled across all profiles.

## Script

```bat
@echo off
setlocal enabledelayedexpansion

REM Check if the current script is running as Administrator
net session >nul 2>&1
if %errorlevel% NEQ 0 (
    echo Script is not running as Administrator. Please run the script as Administrator.
    pause
    exit
)

REM Check if AmazonSSMAgent is running with Administrator user
tasklist /fi "imagename eq amazon-ssm-agent.exe" /v | findstr /i "Administrator" >nul

if errorlevel 1 (
    echo AmazonSSMAgent is not running as Administrator.

    REM Stop the Amazon SSM Agent service
    echo Stopping the AmazonSSMAgent service.
    net stop AmazonSSMAgent

    REM Launch Amazon SSM Agent as Administrator in background
    echo Launching AmazonSSMAgent as Administrator.
    start "AmazonSSMAgent" /B "C:\Program Files\Amazon\SSM\amazon-ssm-agent.exe"

) else (
    echo AmazonSSMAgent is already running as Administrator.
)

REM Checking Firewall Status for all profiles
for %%p in (DomainProfile PrivateProfile PublicProfile) do (
    netsh advfirewall show %%p state | findstr /C:"State                                 OFF" >nul
    if !errorlevel! == 0 (
        echo Firewall for %%p is OFF. Enabling...
        netsh advfirewall set %%p state on
    ) else (
        echo Firewall for %%p is already ON.
    )
)

REM Keep the script open for user's observation
pause
```

## What the Script Does:

1. **Administrator Privileges Check**:
   - The script first checks if it is being run with administrator rights.
   - If not, it prompts the user to run it as an administrator.

2. **AmazonSSMAgent Check**:
   - It checks if the `AmazonSSMAgent` process is running under the Administrator user.
   - If not, it stops the service and then runs the `AmazonSSMAgent` as Administrator in the background.

3. **Windows Firewall Status Check**:
   - For each of the firewall profiles (Domain, Private, and Public):
     - It checks if the firewall is enabled.
     - If it is found to be off for any profile, it enables the firewall for that profile.

## How to Run:

1. **Prepare the Script**:
   - Copy the provided script into a new file on your Windows EC2 instance and save it with a `.bat` extension, for instance `setup_ec2.bat`.

2. **Run as Administrator**:
   - Right-click on the `.bat` file and select `Run as administrator`.

3. **Review the Output**:
   - The script will provide console outputs for each of its steps, informing you of its progress.
   - Once completed, press any key to exit the script.

---

**Note**: Always ensure that the script is run with administrator rights for it to function correctly.
