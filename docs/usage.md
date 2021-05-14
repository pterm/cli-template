# Using Modules

<a id="asciicast-9w0cCZ34umw2xfYfxUy9rNoMm" data-autoplay="true" data-loop="true"></a>

> Using modules in dops is super easy!

## The two ways to use dops

You can use dops in two different ways.\
**[Interactive mode](#interactive-mode)** or just by **[using the command line](#command-line-usage)**.

Not sure which mode fits best to your or your task?\
This list might help you:

| [Interactive mode](#interactive-mode) | [Command line](#command-line-usage)             |
| --------------------------------------|-------------------------------------------------|
| Testing modules                       | The user has experience with command line tools |
| Experimenting with modules            | Something has to be done fast                   |
| Learning the usage of a module        | Dops is running in a CI environment             |

## Interactive Mode

> [!NOTE]
> The interactive mode is good for getting familiar with a module and experimenting with it. For frequent use, or use inside a CI/CD system, we recommend [running dops from the command line](#command-line-usage).

### Selecting a module

To use the interactive mode of dops, you can simply start dops without a module name. 

```command
dops
```

In interactive mode you can select a module with the arrow keys of your keyboard and confirm it with the return key.

### Running a module

Once you have selected a module, a form appears showing all the options you can use for the selected module. If the selected module does not have any options, you will still be asked for confirmation, to prevent unintentional execution of a module.

## Command Line Usage

> [!NOTE]
> This is the recommended way to use dops.

To use dops from the command line, you must specify the module name and all options.

```command
dops <module> [module options] [submodule] [submodule options] [arguemnts...]
```

### Display Help Messages

To show all modules with their short description you can call the help of dops:

```command
dops -h
```

To display the help of a module you can call the help of that module:

```command
dops <module> -h
```

To display the help of a submodule:

```command
dops <module> <submodule> -h
```

### What are options/flags?

A module might have different options. We will call the **flags** now.\
\
Flags are identified by a leading dash, as for example: `-f`, `--flag`.\
A single leading dash (`-f`) means that it's the short form of a flag. Most times the short form is the first letter of the long flag name. If you want to combine flags you can do it like this:\
 `--pretty --json --nocolor`. You can also combine their short versions like this: `-pjn`.\
 \
Flags can have different values. For example a flag, which accepts text can be used like this: `--name MarvinJWendt`, or like this: `--name "Marvin Wendt"`.\
A flag that accepts numbers is used like this: `--count 3`.\
A flag that accepts booleans (yes/no on/off) is handled the following way: If the flag is set, it returns `true` so the option is activated. If the flag is not set it returns `false` so the flag is not activated. It has no value, but you can still set the value to `true` or `false` if you want.\
\
All values can also be set like this: `--flag=value`, `--count=3` if you preffer this style.

### What are submodules

They are modules nested inside another module. They have the following syntax: `dops <module> [module options] <subcommand> [subcommand options]`.\

#### Example

```command
//Long version
dops random-generator --seed HelloWorld string --length 15

// Short version:
dops rg -s HelloWorld s -l 15
```

This runs the module [`random-generator`](commands/random-generator.md) with a random seed of `HelloWorld`.\
The chosen command is `string` so the random generator knows to generate letters.\
The flag `--lenght 15` tells the `string` command that the text, which should be generated, should have a length of 15 letters.
