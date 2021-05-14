# Quick Start - Install dops

> Dops is crossplatform compatible.\
> Use the following instructions to install or update dops.

<!-- tabs:start -->

#### ** Windows **

### Windows Command

?> **NOTICE**
This command has to be run in an elevated powershell prompt.

```powershell
iwr -useb dops-cli.com/get/windows | iex
```

!> **WARNING**
This command is executed with administrative rights!\
Of course dops, just like the installation script, is not harmful to your computer. But it's good practise to control every script that is run with administrative rights. You can copy the URL from the command and paste it into your browser to view the script. Also, all of our installation scripts are located in a GitHub repository at the URL: https://github.com/dops-cli/get-dops 

#### ** Linux **

### Linux Command

?> **NOTICE**
This command has to be run in an elevated shell.

```bash
curl -s https://dops-cli.com/get/linux | sudo bash
```

!> **WARNING**
This command is executed with administrative rights!\
Of course dops, just like the installation script, is not harmful to your computer. But it's good practise to control every script that is run with administrative rights. You can copy the URL from the command and paste it into your browser to view the script. Also, all of our installation scripts are located in a GitHub repository at the URL: https://github.com/dops-cli/get-dops 

#### ** macOS **

### macOS Command

?> **NOTICE**
This command has to be run in an elevated shell.

```bash
/bin/bash -c "$(curl -fsSL https://dops-cli.com/get/linux)"
```

!> **WARNING**
This command is executed with administrative rights!\
Of course dops, just like the installation script, is not harmful to your computer. But it's good practise to control every script that is run with administrative rights. You can copy the URL from the command and paste it into your browser to view the script. Also, all of our installation scripts are located in a GitHub repository at the URL: https://github.com/dops-cli/get-dops 

#### ** Compile from source **

### Compile from source with Golang

?> **NOTICE**
To compile dops from source, you have to have [go](https://golang.org/) installed.

Compiling dops from source has the benefit that the command is the same on every platform.\
You have to have [go](https://golang.org/) installed to use this command.\
It is not recommended to install go only for the installation of dops.


!> **WARNING**
Make sure you don't run this command in a folder which contains a `go.mod` file, otherwise you will add dops as a dependency to your go module.

```bash
go get -u github.com/dops-cli/dops
```

<!-- tabs:end -->