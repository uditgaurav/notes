# Running the PowerShell Script on Windows

This guide provides instructions on how to run the provided PowerShell script to download and install the specified tools (`clumsy`, `diskpd`, and `Testlimit`) on a Windows machine.

Before you proceed, make sure:

- You have administrative rights on your computer.
- The script is saved on your computer, for example, as `install_tools.ps1`.

## Instructions

### 1. Open PowerShell as Administrator

- Press the `Windows` key, type `PowerShell`.
- Right-click on `Windows PowerShell`, then choose `Run as administrator`.

### 2. Navigate to the Script's Directory

If your script is saved in, for example, `C:\scripts`, you would navigate to it using:

```powershell
cd C:\scripts
```

### 3. Run the Script

Execute the script by typing its name:

```powershell
.\install_tools.ps1
```

Wait for the script to finish executing. It will download and extract the tools to the specified directories.


# Verification Guide for Tool Installation

This guide will help you verify the successful installation of the following tools:

1. `clumsy`
2. `diskpd`
3. `Testlimit`

## Prerequisites

Before we begin, ensure that:

- The tools have been downloaded and extracted using the provided PowerShell script.
- You have access to the `Downloads` folder of the user who ran the script.

## Verification Steps

### 1. Verify `clumsy`

1. Navigate to the `Downloads\clumsy` directory in File Explorer.

    ```plaintext
    %USERPROFILE%\Downloads\clumsy
    ```

2. Check for the presence of the `clumsy.exe` executable.

3. (Optional) Double-click on `clumsy.exe` to run the application. You should see the `clumsy` user interface.

### 2. Verify `diskpd`

1. Navigate to the `Downloads\diskpd` directory in File Explorer.

    ```plaintext
    %USERPROFILE%\Downloads\diskpd
    ```

2. Check for the presence of the `DiskSpd.exe` executable.

3. (Optional) Open a Command Prompt and navigate to the directory above. Run the following command to view the help text:

    ```bash
    DiskSpd.exe -?
    ```

### 3. Verify `Testlimit`

1. Navigate to the `Downloads\Testlimit` directory in File Explorer.

    ```plaintext
    %USERPROFILE%\Downloads\Testlimit
    ```

2. Check for the presence of the `Testlimit64.exe` (for 64-bit) or `Testlimit.exe` (for 32-bit) executable.

3. (Optional) Open a Command Prompt and navigate to the directory above. Run the following command to view the help text:

    ```bash
    Testlimit64.exe -?
    ```

## Conclusion

If you can find the respective executables for each tool and they run as expected, the installation was successful. If any issues arise, ensure that the PowerShell script completed without errors and that the respective ZIP files were correctly extracted.
