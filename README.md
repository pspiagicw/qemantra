# Qemantra
![GitHub issues](https://img.shields.io/github/issues-raw/pspiagicw/qemantra?logoColor=%23ffb86c&style=for-the-badge)
![GitHub](https://img.shields.io/github/license/pspiagicw/qemantra?style=for-the-badge)
![GitHub last commit](https://img.shields.io/github/last-commit/pspiagicw/qemantra?style=for-the-badge)
![GitHub pull requests](https://img.shields.io/github/issues-pr/pspiagicw/qemantra?style=for-the-badge)

![QEMantra Icon](./assets/qemantra_icon_128.png)


Qemantra is a tool for creating and running QEMU Virtual Machines.

QEMU is better and sometimes faster than VirtualBox , but does not have any
command-line central managing solution. This tool does not run your virtual machines.
It simply aims to configure and run Virtual Machines using `QEMU`.

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

### Documentation
Currently documentation only exists as a `--help` flag.

## Motivation ?

Virtualbox is good, but it has a QT interface and it's command line inteface is hectic at best.
QEMU has no official (or good) frontend , the command line interface is mature , but has no central managing solution.

Qemantra aims to become a simple and convinient way to manage Virtual Machines. It is designed for the casual Virtualizer.


## Installation ?

Currently you can only build from source.

- Qemantra is built in golang , you must have golang installed.
- Clone the repository.
- `make build` to build the binary.
- Move the binary to a folder in your $PATH variable. 
  Or use `make install` to install in default `$GOBIN`

## Contributing

Anybody is welcome to contribute!
 
Qemantra is written in Golang , so Golang developers can contribute in the technical aspect.
If you want to contribute non-technically , then too you are welcome!
There are lot's of work in documentation and other aspects!
For bugs and feature requests , open a issue.

