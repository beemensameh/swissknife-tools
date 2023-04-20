
# Swissknife Tools

A quick and amazing tools for speed up your work

## Badges

[![Release](https://github.com/beemensameh/swissknife-tools/actions/workflows/github-workflow.yml/badge.svg)](https://github.com/beemensameh/swissknife-tools/actions/workflows/github-workflow.yml)
[![Build](https://github.com/beemensameh/swissknife-tools/actions/workflows/go-ci.yml/badge.svg)](https://github.com/beemensameh/swissknife-tools/actions/workflows/go-ci.yml)

[![Go Version](https://img.shields.io/github/go-mod/go-version/beemensameh/swissknife-tools?logo=go)](https://go.dev/)
[![Go Report Card](https://goreportcard.com/badge/github.com/beemensameh/swissknife-tools)](https://goreportcard.com/report/github.com/beemensameh/swissknife-tools)
[![codecov](https://codecov.io/gh/beemensameh/swissknife-tools/branch/main/graph/badge.svg)](https://codecov.io/gh/vektra/mockery)

[![license](https://img.shields.io/github/license/beemensameh/swissknife-tools)](https://github.com/beemensameh/swissknife-tools/blob/main/LICENCE)
![Repo size](https://img.shields.io/github/repo-size/beemensameh/swissknife-tools)
![Number of downloads](https://img.shields.io/github/downloads/beemensameh/swissknife-tools/total)

[![Open in Visual Studio Code](https://img.shields.io/badge/-Open%20in%20Visual%20Studio%20Code-2365cb?logo=visualstudiocode&labelColor=2d2b32&logoColor=007acc)](https://github.dev/beemensameh/swissknife-tools)

## Features

- Get time with many format and live option
- Generate UUID

## Run Locally

### From git repo
- Download the appropriate version for you from release page
- Run this command and enjoy
    - For Linux
    ```sh
    $ make build
    $ ./bin/swisstool
    ```

### From release
- Choose the latest release
- Download the version which compatible with your OS
- Run this command and enjoy
    - For Linux and Mac
    ```sh
    ./swisstools
    ```
    - For Windows
    ```sh
    swisstools.exe
    ```
- **[Optional]** for using auto-complete commands for swissknife-tools:<br>
    If it is not installed already, you can install it via your OS's package manager.
    To load completions in your current shell session:
    ```sh
    source <(swisstools completion bash)
    ```
    To load completions for every new session, execute once:
    #### Linux:
    ```sh
    swisstools completion bash > /etc/bash_completion.d/swisstools
    ```
    #### macOS:
    ```sh
    swisstools completion bash > $(brew --prefix)/etc/bash_completion.d/swisstools
    ```
    You will need to start a new shell for this setup to take effect.

## Usage/Examples

* [Time](./docs/time.md)
* [JSON](./docs/json.md)
* [UUID](./docs/uuid.md)
* [Hash](./docs/hash.md)
