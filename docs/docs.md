# cli-template

## Usage
> This cli template shows the date and time in the terminal

cli-template [global options] command [options] [arguments...]

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
|`cli-template date`|Prints the current date.|
|`cli-template help`|Help about any command|
|`cli-template time`|Prints the current time|
# ... date
`cli-template date`

## Usage
> Prints the current date.

cli-template [global options] command [options] [arguments...]

## Flags
|Flag|Usage|
|----|-----|
|`-f, --format string`|specify a custom date format (default "02 Jan 06")|
# ... help
`cli-template help`

## Usage
> Help about any command

cli-template [global options] command [options] [arguments...]

## Description

```
Help provides help for any command in the application.
Simply type cli-template help [path to command] for full details.
```
# ... time
`cli-template time`

## Usage
> Prints the current time

cli-template [global options] command [options] [arguments...]

## Description

```
You can print a live clock with the '--live' flag!
```

## Flags
|Flag|Usage|
|----|-----|
|`-l, --live`|live output|


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 01 June 2021**
