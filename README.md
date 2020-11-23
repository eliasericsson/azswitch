![Go](https://github.com/eliasericsson/azswitch/workflows/Go/badge.svg)
# azswitch
Switch quicker between Azure subscriptions

## Requirements
* Azure CLI - [installation instructions](https://docs.microsoft.com/sv-se/cli/azure/install-azure-cli)
* fzf - [installation instructions](https://github.com/junegunn/fzf)

## Installation
* [Linux](#ubuntu/debian/arch)

### Ubuntu/Debian/Arch
Copy the binary into your $PATH
```sh
# Download the latest release
wget https://github.com/eliasericsson/azswitch/releases/download/v0.1.1/azswitch_<version>_<os>_<arch>.tar.gz

# Extract the archive contents
tar -xvf azswitch_0.1.1_linux_amd64.tar.gz

# Move the binary to a directory in your $PATH
mv azswitch /usr/local/bin/
```