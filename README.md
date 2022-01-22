# image-protector
Image Protector - A CLI tool provide an adaptive self-repairing authentication method for allowing the user to encrypt and decrypt the image.

With the following image:

<img src="./docs/test1.png" width="400px">

Add the repairing information into the image:

<img src="./docs/test2.png" width="400px">

If the photo is damaged:

<img src="./docs/test3.png" width="400px">

Use self-repairing function to fix it:

<img src="./docs/test4.png" width="400px">

# CLI Overview

```
This tool provides an easy and extensible way to protect the image.

Usage:
  protector [command]

Available Commands:
  encrypt     Encrypt the repairing information on photo
  help        Help about any command
  repair      Repair the damaged image

Flags:
  -h, --help   help for protector

Use "protector [command] --help" for more information about a command.
```

## Getting started

- [Download](#download)
- [Commands](#commands)
    - [Encrypt](#encrypt)
    - [Repair](#repair)
- [Copyright Notice](#copyright-notice)

### Download

To download the latest release, go to the [release page](https://github.com/tsunejui/image-protector/releases)

### Commands

You can run the following commands for the different requirement:

<a id="encrypt"></a>
**Encrypt** - Encrypt the repairing information on photo

```
Hide the repairing information in each pixel on specified image

Usage:
  protector encrypt [flags]

Flags:
  -f, --file string     specify a file
  -h, --help            help for encrypt
  -o, --output string   export the image
```

<a id="repair"></a>
**Repair** - Repair the damaged image

```
Repair the damaged image

Usage:
  protector repair [flags]

Flags:
  -f, --file string     specify a file
  -h, --help            help for repair
  -o, --output string   export the image
```

### Copyright Notice

This tool is solely for personal and non-commercial use and for reference only. 
