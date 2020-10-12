# Docxy CLI

This is a CLI for [Docxy](https://docxy.traction.one).

Docxy is a React based open-source documentation site
generator. Build beautiful, blazing fast documentation sites
for your projects with just markdown.

The Docxy CLI is used to:
- Create and initialize a Docxy application
- Spin up a hot-reloading local development server
- Build a production-ready website ready for deployment


## Installation

You can install the Docxy CLI in one of the two ways
mentioned below.

### Option 1 - Install from Source

> Note that, you need to have [Go](https://golang.org)
> installed for this.

You can directly install it from source simply by running
the following command:

```bash
go get -u github.com/docxy/cli
```

### Option 2 - Get Binaries

You can directly get the updated version of the Docxy CLI
binaries from [GitHub Releases](https://github.com/docxy/cli/releases),
or you can build it from source


## Usage

After installation, it can be accessed globally using the
`docxy` command.

```
Usage:
  docxy [command]

Available Commands:
  help        Help about any command
  create      Creates a Docxy project
  init        Initializes a Docxy project
  serve       Starts the development server
  build       Builds the Docxy website

Flags:
  -h, --help   help for docxy

Use "docxy [command] --help" for more information about a command.
```

## Contributing

Any and all contributions are welcomed.

If you want to contribute code, you need to have a
development environment set up for working with Go.

## License

Copyright © 2020 Sankarsan Kampa

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at https://mozilla.org/MPL/2.0/.
