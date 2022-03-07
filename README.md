<h1 align="center">cli-template ✨</h1>

<p align="center">⚗ A template for beautiful, modern, cross-platform compatible CLI tools written with Go!</p>

<p align="center">
<img src="https://user-images.githubusercontent.com/31022056/119876432-2e38bf00-bf28-11eb-859b-38f925b593e8.gif" alt="Screenshot">
</p>

----

<p align="center">
<strong><a href="#getting-started">Getting Started</a></strong>
|
<strong><a href="https://github.com/pterm/cli-template/wiki/">Wiki</a></strong>
</p>

----

|This template features|
|----------------------|
|[Modern Style 💎](#modern-style)|
|[Build on top of cobra 🐍](#build-on-top-of-cobra)|
|[Automatic Releases 🚀](#automatic-releases)|
|[Easy installation for your users (one command) 🐥](#install-ready)|
|[Automatic Website/Docs Generation 🌐](#automatic-website-generation)|
|[Automatic Deployment 🔝](#automatic-deployment)|
|[Update Checking ♻](#update-checking)|
|[Custom CI-System 🤖](#custom-ci-system)|
|[Custom Libraries 🔬](#custom-libraries)|

## Getting Started

**You can find an in-depth tutorial in the Wiki here: [Getting Started](https://github.com/pterm/cli-template/wiki/Getting-Started)**

1. Click on <kbd>Use this template</kbd> on the top of the page.
1. Enable GitHub Pages and set the path to `/docs`.
1. Create a personal access token (with repo scope) and add it as a repository secret (name: `REPO_ACCESS_TOKEN`).
1. Clone and open your repository to change the description of your CLI in `./cmd/root.go`.
1. After you have set up your programm you have to create the very first release manually (`v0.0.1`), to initialize the CI-System. (Don't worry if your CLI tool doesn't do anything yet. It's common that the `v0.0.1` release is just the plain project setup)
1. The setup is done and you can start to code!

This template uses [spf13/cobra](https://github.com/spf13/cobra) as CLI framework.  
You can find their documentation here: [cobra.dev](https://cobra.dev/)

## Features

### Modern Style

> [PTerm Documentation Link](https://pterm.sh/)

This template uses [PTerm](https://github.com/pterm/pterm) to provide colorful, cross-platform compatible output by default!  
By using PTerm, you can output progressbars, colored text, charts and many more.  

### Build on top of cobra
> [Cobra Documentation Link](https://cobra.dev/)

Cobra is a popular CLI framework for Go.

### Automatic Releases

> [Documentation link](https://github.com/pterm/cli-template/wiki/Automatic-Releases)

Our custom made CI system will detect when you change the version of your CLI and will **automatically create a new GitHub release for you**.
You'll never have to deploy your CLI tools manually again!

### Install Ready

> [Instl Documentation link](https://docs.instl.sh)

The [automatic releases](#automatic-releases) contain binaries, for the most common operating systems.  
This binaries can be installed using [instl.sh](https://docs.instl.sh).

This means that your users, can install your CLI Tool with a single command!  
The command will automatically be put into your README.md, after you click <kbd>Use this template</kbd>.

The commands will look like below, depending on the OS, your user has.  

**Windows**

```powershell
iwr instl.sh/username/reponame/windows | iex 
```

**macOS**

```shell
curl -sSL instl.sh/username/reponame/macos | bash   
```

**Linux**

```shell
curl -sSL instl.sh/username/reponame/linux | bash  
```

### Automatic Website Generation

> [Documentation link](https://github.com/pterm/cli-template/wiki/Automatic-Website-Generation)

Every time you push a new commit, a GitHub Pages website is created/updated, which documents your whole CLI tool automatically.
You don't need to document anything by yourself.

You only need to enable GitHub Pages by going to your repository settings (set the path to `/docs`).

### Automatic Deployment

> [Goreleaser Documentation link](https://goreleaser.com)

This template uses Goreleaser to build binaries of GitHub releases, for the most common operating systems.  
Since we feature [Automatic Releases](#automatic-releases), your whole deployment process is automated, when you increase
the version of your CLI tool.

### Update Checking

Your users will be notified if a new version of your tool is availble.

### Custom CI-System

> [Documentation link](https://github.com/pterm/cli-template/wiki/Custom-CI-System)

We wrote a custom CI-System, which will be included in your repository.  
It sets up the template, creates/updates the website, writes docs for you and run your tests. (And much more!)

### Custom Libraries

We use libraries, which were specially made for this template, to give you the best user experience without a ton of code.
