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
