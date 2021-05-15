<h1 align="center">cli-template ✨</h1>

<p align="center">⚗ A template for beautiful, modern, cross-platform compatible CLI tools written with Go!</p>

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
|[Build on top of cobra 💎](#build-on-top-of-cobra)|
|[Automatic Releases 🚀](#automatic-releases)|
|[Automatic Website/Docs Generation 🌐](#automatic-website-generation)|
|[Automatic Deployment 🔝](#automatic-deployment)|
|[Custom CI-System 🤖](#custom-ci-system)|
|[Custom Libraries 🔬](#custom-libraries)|
|[Easy installation for your users (one command) 🐥](#install-ready)|

## Getting Started

You can find an in-depth tutorial in the Wiki here: [Getting Started](https://github.com/pterm/cli-template/wiki/Getting-Started)

1. Click on <kbd>Use this template</kbd> on the top of the page.
1. Wait for your repository to initialize.
1. Create a personal access token and add it as a repository secret (`REPO_ACCESS_TOKEN`). [(GitHub Docs)](https://docs.github.com/es/actions/reference/encrypted-secrets#creating-encrypted-secrets-for-a-repository)
1. Clone and open your repository to change the name and description of your CLI in `./cmd/root.go`.
1. After you have set up your programm you have to create the very first release manually (`v0.0.1`) to initialize the CI-System.
    1. Don't worry if your CLI tool doesn't do anything yet. It's common that the `v0.0.1` release is just the plain project setup.
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

### Automatic Website Generation

> [Documentation link](https://github.com/pterm/cli-template/wiki/Automatic-Website-Generation)

Everytime you push a new commit, a GitHub Pages website is created/updated, which documents your whole CLI tool automatically.
You don't need to document anything by yourself.

You only need to enable GitHub Pages by going to your repository settings. (set to `/docs`)

### Automatic Deployment

### Custom CI-System

### Custom Libraries

We use libraries, which were specially made for this template, to give you the best user experience without a ton of code.

### Install Ready
