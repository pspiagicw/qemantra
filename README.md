# Qemantra
![GitHub issues](https://img.shields.io/github/issues-raw/pspiagicw/qemantra?logoColor=%23ffb86c&style=for-the-badge)
![GitHub](https://img.shields.io/github/license/pspiagicw/qemantra?style=for-the-badge)
![GitHub last commit](https://img.shields.io/github/last-commit/pspiagicw/qemantra?style=for-the-badge)
![GitHub pull requests](https://img.shields.io/github/issues-pr/pspiagicw/qemantra?style=for-the-badge)
![GitHub Stars](https://img.shields.io/github/stars/pspiagicw/qemantra?style=for-the-badge)

Qemantra is a command-line tool for creating and managing QEMU Virtual Machines.

![QEMantra banner](./assets/banner.png)


QEMU is better and sometimes faster than VirtualBox , but does not have any
command-line central managing solution. This tool aims to manage and run your virtual machines using `QEMU`.

![img](./assets/gifs/intro.gif)

## Features

- Create , list and run virtual machines.
- Create configurations to run on demand.
- Use features of QEMU like KVM , multiple architecture support etc.

## Examples
### Running a Virtual Machine
![img](./assets/gifs/run-iso.gif)

### Run a machine with boot menu
![img](./assets/gifs/run-menu.gif)

### Create a Machine
![img](./assets/gifs/create-machine-disk.gif)

### Create a Machine with Disk
![img](./assets/gifs/create-machine-no-disk.gif)

### Create a Image
![img](./assets/gifs/create-img.gif)

### Usage
Call `qemantra` in your terminal.
```sh
$ qemantra
```

#### Check
Run `qemantra check` for checking for dependencies and configuration.

Highly recommended before using `qemantra`!.

#### Create Machine

The `qemantra create-machine` subcommand provides functionality to create machines.
| Option        | Description                                             |
|---------------|---------------------------------------------------------|
| `--name`      | Name of the machine                                     |
| `--no-disk`   | Don't create a disk                                     |
| `--disk-name` | Name of the disk(Not applicable when using `no-disk`)   |
| `--disk-size` | Size of the disk(Not applicable when using `disk-size`) |
| `--mem-size`  | RAM to provide to the VM                                |
| `--cpu-cores` | Cores to provide to the RAM.                            |

#### Running a machine

The `qemantra run` subcommand provides functionality to run a virtual machine.
| Option           | Description                                                                      |
|------------------|----------------------------------------------------------------------------------|
| `--name`         | Name of the machine                                                              |
| `--iso`          | Path to the ISO(Relative path works)                                             |
| `--disk`         | Disk name to add to boot order(Should be in default qemantra directory)          |
| `--externaldisk` | Path to a external disk to add to boot order(Any disk , not managed by qemantra) |
| `--boot`         | Boot options while starting the VM                                               |

The boot options can be either `menu` which provides a menu to choose between boot devices. Or you can use `iso` option to directly boot the given iso.
    
#### List machines
Use `qemantra list` to list currently configured machines. Use `--verbose` option to list more information about the VM(Mem , Cpu etc).

You can list the images managed by `qemantra` by using `--images` option to `qemantra list`.In most cases the virtual machine would automatically use the correct disk without any problem.

## Motivation ?

Virtualbox is good, but it has a QT interface and it's command line inteface is hectic at best.
QEMU has no official (or good) frontend , the command line interface is mature , but has no central managing solution.

Qemantra aims to become a simple and convinient way to manage Virtual Machines. It is designed for the casual Virtualizer.


## Installation ?

Currently you can only build from source.You will need `Golang` installed on your system.
As a prerequisite you also need QEMU installed.


Install using `go install` by running `go install github.com/pspiagicw/qemantra@latest`.

## Contributing

Anybody is welcome to contribute!
 
Qemantra is written in Golang , so Golang developers can contribute in the technical aspect.
If you want to contribute non-technically , then too you are welcome!
There are lot's of work in documentation and other aspects!
For bugs and feature requests , open a issue.

