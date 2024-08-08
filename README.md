# Atlas CLI Example Plugin
This repository was created for developers who want to extend the [Atlas CLI](https://github.com/mongodb/mongodb-atlas-cli) with their own functionality. 

## How does it work?

A plugin for the Atlas CLI is essentially a standalone CLI. When installed into the Atlas CLI, the plugin is intagrated as a normal subcommand to seamlessly extend the core functionality. When a plugin command is invocated, the Atlas CLI executes the plugin's binary and forwards all arguments, environment variables and standard input/output/error streams to it. This allows the plugin to function as if it was part of the Atlas CLI.

## Developing a Plugin

The only files a plugin strictly requires is an executable binary file and a `manifest.yaml` file.

### Manifest file
The manifest file (`manifest.yaml`) is an essential part of an Atlas CLI plugin and needs to live in the root folder of the release asset. Here is what the manifest file for this example plugin looks like:

```yaml
name: atlas-cli-plugin-example
description: this is an example plugin
version: 2.0.0
github:
    owner: mongodb
    name: atlas-cli-plugin-example
binary: example
commands: 
    example: 
        description: Root command of the atlas cli plugin example
```

| Field | Description | Optional |
| - | - | - |
| `name` | The name of the plugin | No
| `description` | A brief description of the plugin. | No
| `version` | The version of the plugin.  | No
| `github.owner` | The GitHub repository owner | Yes
| `github.name`| The GitHub repository name | Yes
| `binary`| The name of the plugin's binary file | Yes
| `commands.<command_name>` | An object where the key is the command name and the value is a description of the command | No (there needs to be at least one command)
| `commands.<command_name>.description` | A description of the command | No

> **Important:** 
> - The `version` field in the manifest file is crucial for the Atlas CLI to identify the installed plugin version and manage updates. Ensure that this field is always aligned with the plugin's release version to enable correct updates.
> - If your plugin is not part of Github repository you can omit the entire `github` section. However, if your plugin is hosted on GitHub and you want to enable update functionality, you must include this section.

## Getting Started

To get started developing your own plugin you can use this repository as a starting point. Just click on "Use this as a template" in the top right corner to create your plugin repository based on this example plugin.
