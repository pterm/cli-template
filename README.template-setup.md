<h1 align="center">{{ .Name }}</h1>
<p align="center">{{ .Short }}</p>

<p align="center">

<a style="text-decoration: none" href="https://github.com/{{ .ProjectPath }}/releases">
<img src="https://img.shields.io/github/v/release/{{ .ProjectPath }}?style=flat-square" alt="Latest Release">
</a>

<a style="text-decoration: none" href="https://github.com/{{ .ProjectPath }}/releases">
<img src="https://img.shields.io/github/downloads/{{ .ProjectPath }}/total.svg?style=flat-square" alt="Downloads">
</a>

<a style="text-decoration: none" href="https://github.com/{{ .ProjectPath }}/stargazers">
<img src="https://img.shields.io/github/stars/{{ .ProjectPath }}.svg?style=flat-square" alt="Stars">
</a>

<a style="text-decoration: none" href="https://github.com/{{ .ProjectPath }}/fork">
<img src="https://img.shields.io/github/forks/{{ .ProjectPath }}.svg?style=flat-square" alt="Forks">
</a>

<a style="text-decoration: none" href="https://github.com/{{ .ProjectPath }}/issues">
<img src="https://img.shields.io/github/issues/{{ .ProjectPath }}.svg?style=flat-square" alt="Issues">
</a>

<a style="text-decoration: none" href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

<br/>

<a style="text-decoration: none" href="https://github.com/{{ .ProjectPath }}/releases">
<img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads">
</a>

<br/>

</p>

----

<p align="center">
<strong><a href="{{ .GitHubPagesURL }}/#/installation">Installation</a></strong>
|
<strong><a href="{{ .GitHubPagesURL }}/#/docs">Documentation</a></strong>
|
<strong><a href="{{ .GitHubPagesURL }}/#/CONTRIBUTING">Contributing</a></strong>
</p>

----

{{ .Long }}

## Installation

Run the following command in a terminal and you're ready to go!

**Windows**
```powershell
{{ .InstallCommandWindows }}
```

**macOS**
```bash
{{ .InstallCommandMacOS }}
```

**Linux**
```bash
{{ .InstallCommandLinux }}
```
