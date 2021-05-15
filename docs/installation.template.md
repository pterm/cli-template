# Quick Start - Install {{ .Name }}

> [!TIP]
> {{ .Name }} is installable via [instl.sh](https://instl.sh).\
> You just have to run the following command and you're ready to go!

<!-- tabs:start -->

#### ** Windows **

### Windows Command

```powershell
{{ .InstallCommandWindows }}
```

#### ** Linux **

### Linux Command

```bash
{{ .InstallCommandLinux }}
```

#### ** macOS **

### macOS Command

```bash
{{ .InstallCommandMacOS }}
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
