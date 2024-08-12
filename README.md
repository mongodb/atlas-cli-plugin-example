# Atlas CLI Plugin Example
This repository was created for developers who want to extend the [Atlas CLI](https://github.com/mongodb/mongodb-atlas-cli) with their own functionality. 

## How does it work?

### Plugin Execution
A plugin for the Atlas CLI consists of a manifest file and a binary.
The manifest defines which subcommands the plugin exposes, and the binary is called when one of these subcommands are called.

When a subcommand is invoked, the Atlas CLI executes the plugin's binary and forwards:
- all arguments 
- environment variables
- standard input/output/error

### Plugin Installation
When you install a plugin, the Atlas CLI takes the following steps:
1. download the release asset for the user's OS from the plugin GitHub repository
1. unpack the assets (manifest + binary + optinal extras) into the Atlas CLI plugin directory

## Developing a Plugin

Plugins can be developed with any programming language. A plugin directory can contain any number of files, but it strictly requires two: an executable binary file and a `manifest.yml` file.

### Manifest File
The manifest file (`manifest.yml`) is an essential part of an Atlas CLI plugin and needs to live in the root folder of the release asset. Here is what the manifest file for this example plugin looks like:

```yaml
name: atlas-cli-plugin-example
description: this is an example plugin
version: 2.0.1
github:
    owner: mongodb
    name: atlas-cli-plugin-example
binary: binary
commands: 
    example: 
        description: Root command of the atlas cli plugin example
```

| Field | Description | Optional |
| - | - | - |
| `name` | The name of the plugin | No
| `description` | A brief description of the plugin. | No
| `version` | The [semantic versioning](https://semver.org/) of the plugin.  | No
| `github.owner` | The GitHub repository owner | Yes
| `github.name`| The GitHub repository name | Yes
| `binary`| The name of the plugin's binary file | No
| `commands.<command_name>` | An object where the key is the command name and the value is a description of the command | No (a plugin needs at least one command)
| `commands.<command_name>.description` | A description of the command | No

> **Important:** 
> - The `version` field in the manifest file is crucial for the Atlas CLI to identify the installed plugin version and manage updates. Ensure that this field is always aligned with the plugin's release version to enable correct updates.
> - If your plugin is not part of a Github repository you can omit the entire `github` section. However, if your plugin is hosted on GitHub and you want to enable update functionality, you must include this section.


You may notice that there is no `manifest.yml` file in this repository, only a `manifest.template.yml`. This is because we use a [GitHub Action](https://github.com/mongodb/atlas-cli-plugin-example/blob/master/.github/workflows/release.yml) to automatically generate the `manifest.yml` file. Every time a new release is published, the template is used to create the required `manifest.yml` file, and the release's version is automatically inserted into the `version` field. The [GoReleaser](https://goreleaser.com/) is used to create release assets that contain the plugin's binary and the `manifest.yml`.

### Executable Binary

The Atlas CLI runs the executable binary file when a plugin command is invoked. Arguments, environment variables, and standard input/output/error streams are passed through to this executable. The name of the binary file needs to be defined in the `binary` field of the `manifest.yml`.

In this example plugin, we use the [GoReleaser](https://goreleaser.com/) to automatically generate the executable binary as part of the release assets.

## Getting Started

To begin developing your own plugin, you can use this repository as a starting point. Click "Use this template" in the top right corner to create a new repository based on this example plugin. Afterward, you can start customizing your plugin by:
1. Removing any unnecessary commands from `./internal/cli/`
1. Updating the `rootCmd` in `./cmd/plugin/main.go` 
1. Updating the name from `atlas-cli-plugin-example` to your preferred name in `manifest.template.yml` and `.goreleaser.yaml`

To install your plugin in the Atlas CLI, simply create a new release using a tag that follows [semantic versioning](https://semver.org/). This will trigger the [release workflow](https://github.com/mongodb/atlas-cli-plugin-example/blob/master/.github/workflows/release.yml), which generates the release assets containing the executable binary and the manifest.yml. Once the release is complete, you can install your plugin in the Atlas CLI:

```bash
atlas plugin install <repository-owner>/<repository-name>
```
