
# Qemantra
![GitHub issues](https://img.shields.io/github/issues-raw/pspiagicw/qemantra?logoColor=%23ffb86c&style=for-the-badge)
![GitHub](https://img.shields.io/github/license/pspiagicw/qemantra?style=for-the-badge)
![GitHub last commit](https://img.shields.io/github/last-commit/pspiagicw/qemantra?style=for-the-badge)
![GitHub pull requests](https://img.shields.io/github/issues-pr/pspiagicw/qemantra?style=for-the-badge)

![QEMantra Icon](./assets/qemantra_icon_128.png)
<a id="orgd792ca5"></a>




Qemantra is a tool for creating and running QEMU Virtual Machines.

![img](./assets/gifs/intro.gif)


## Table of Contents

1.  [Qemantra](#orgd792ca5)
    1.  [Motivation ?](#org7e245cb)
    2.  [Installation ?](#org1a144a1)
    3.  [Features](#org18f5296)
    4.  [Contributing](#orgfacc51e)


<a id="org18f5296"></a>

## Features

-   Create , list and run virtual machines.
-   Create configurations to run on demand.
-   Use features of QEMU like KVM , multiple architecture support etc.

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


<a id="org7e245cb"></a>

## Motivation ?

Virtualbox is good, but it has a QT interface and it's command line inteface is hectic at best.
QEMU has no official (or good) frontend , the command line interface is mature , but has no central managing solution.

Qemantra aims to become a simple and convinient way to manage Virtual Machines. It is designed for the casual Virtualizer.


<a id="org1a144a1"></a>

## Installation ?

Currently you can only build from source.

-   Clone the repository
-   `make build` to build the binary.
-   Move the binary to a folder in your $PATH variable.


<a id="orgfacc51e"></a>

## Contributing

Anybody is welcome to contribute , check out the `contribution guide` for more information.
For bugs and feature requests , open a issue.

