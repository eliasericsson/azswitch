![CI](https://github.com/eliasericsson/azswitch/workflows/CI/badge.svg?branch=main)
![Release](https://github.com/eliasericsson/azswitch/workflows/Release/badge.svg)
# azswitch
Switch quicker between Azure subscriptions

## Requirements
* Azure CLI - [installation instructions](https://docs.microsoft.com/sv-se/cli/azure/install-azure-cli)
* fzf - [installation instructions](https://github.com/junegunn/fzf)

## Installation
* [Linux](#ubuntu/debian/arch)
* [Windows](#windows)

### Ubuntu/Debian/Arch
Copy the binary into your $PATH
```sh
# Download the latest release
wget https://github.com/eliasericsson/azswitch/releases/download/v0.1.1/azswitch_<version>_<os>_<arch>.tar.gz

# Extract the archive contents
tar -xvf azswitch_0.1.3_linux_amd64.tar.gz

# Move the binary to a directory in your $PATH
mv azswitch /usr/local/bin/
```

### Windows
Download and extract the executable into a folder of your choice. Then add this folder to your PATH to run the application from any terminal.

> Examples using `C:\Temp\azswitch_0.1.3_windows_64-bit` as the directory

CMD
```cmd
SET "PATH=C:\Temp\azswitch_0.1.3_windows_64-bit;%PATH%"
```
PowerShell
```powershell
$env:path = $env:path + ";C:\Temp\azswitch_0.1.3_windows_64-bit"
```

## Usage
Run the application in a terminal.
```
azswitch
```