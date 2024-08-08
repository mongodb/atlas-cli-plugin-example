# Atlas CLI Example Plugin
This repository was created for developers who want to extend the [Atlas CLI](https://github.com/mongodb/mongodb-atlas-cli) with their own functionality. 

## How does it work?

A plugin for the Atlas CLI is essentially a standalone CLI. When installed into the Atlas CLI, the plugin is intagrated as a normal subcommand to seamlessly extend the core functionality. When a plugin command is invocated, the Atlas CLI executes the plugin's binary and forwards all arguments, environment variables and standard input/output/error streams to it. This allows the plugin to function as if it was part of the Atlas CLI.

## Developing a Plugin

The only files a plugin strictly requires is an executable binary file and a `manifest.yaml` file. 

## Getting Started


