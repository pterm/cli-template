# cli-template

## Usage
> This cli template shows the date and time in the terminal

cli-template

## Description

```
This is a template CLI application, which can be used as a boilerplate for awesome CLI tools written in Go.
This template prints the date or time to the terminal.
```
## Examples

```bash
cli-template date
cli-template date --format 20060102
cli-template time
cli-template time --live
```

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`cli-template completion`|Generate the autocompletion script for the specified shell|
|`cli-template date`|Prints the current date.|
|`cli-template help`|Help about any command|
|`cli-template time`|Prints the current time|
# ... completion
`cli-template completion`

## Usage
> Generate the autocompletion script for the specified shell

cli-template completion

## Description

```
Generate the autocompletion script for cli-template for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`cli-template completion bash`|Generate the autocompletion script for bash|
|`cli-template completion fish`|Generate the autocompletion script for fish|
|`cli-template completion powershell`|Generate the autocompletion script for powershell|
|`cli-template completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`cli-template completion bash`

## Usage
> Generate the autocompletion script for bash

cli-template completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(cli-template completion bash)

To load completions for every new session, execute once:

#### Linux:

	cli-template completion bash > /etc/bash_completion.d/cli-template

#### macOS:

	cli-template completion bash > /usr/local/etc/bash_completion.d/cli-template

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`cli-template completion fish`

## Usage
> Generate the autocompletion script for fish

cli-template completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	cli-template completion fish | source

To load completions for every new session, execute once:

	cli-template completion fish > ~/.config/fish/completions/cli-template.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`cli-template completion powershell`

## Usage
> Generate the autocompletion script for powershell

cli-template completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	cli-template completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`cli-template completion zsh`

## Usage
> Generate the autocompletion script for zsh

cli-template completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	cli-template completion zsh > "${fpath[1]}/_cli-template"

#### macOS:

	cli-template completion zsh > /usr/local/share/zsh/site-functions/_cli-template

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... date
`cli-template date`

## Usage
> Prints the current date.

cli-template date

## Flags
|Flag|Usage|
|----|-----|
|`-f, --format string`|specify a custom date format (default "02 Jan 06")|
# ... help
`cli-template help`

## Usage
> Help about any command

cli-template help [command]

## Description

```
Help provides help for any command in the application.
Simply type cli-template help [path to command] for full details.
```
# ... time
`cli-template time`

## Usage
> Prints the current time

cli-template time

## Description

```
You can print a live clock with the '--live' flag!
```

## Flags
|Flag|Usage|
|----|-----|
|`-l, --live`|live output|


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 24 November 2022**
