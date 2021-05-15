# Quick Start - Install {{ .Name }}

> [!TIP]
> {{ .Name }} is installable via [instl.sh](https://instl.sh).\
> You just have to run the following command and you're ready to go!

<!-- tabs:start -->

#### ** Windows **

### Windows Command

?> **NOTICE**
This command has to be run in an elevated powershell prompt.

```powershell
iwr -useb instl.sh/{{ .ProjectPath }}/windows | iex
```

#### ** Linux **

### Linux Command

?> **NOTICE**
This command has to be run in an elevated shell.

```bash
curl -s https://instl.sh/{{ .ProjectPath }}/linux | sudo bash
```

#### ** macOS **

### macOS Command

?> **NOTICE**
This command has to be run in an elevated shell.

```bash
/bin/bash -c "$(curl -fsSL instl.sh/{{ .ProjectPath }}/macos)"
```

#### ** Compile from source **

### Compile from source with Golang

?> **NOTICE**
To compile {{ .Name }} from source, you have to have [Go](https://golang.org/) installed.

Compiling {{ .Name }} from source has the benefit that the build command is the same on every platform.\
It is not recommended to install Go only for the installation of {{ .Name }}.

```command
go install github.com/{{ .ProjectPath }}@latest
```

<!-- tabs:end -->