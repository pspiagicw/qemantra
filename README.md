
<a id="orgd792ca5"></a>

# Qemantra

Qemantra is a tool for creating and running QEMU Virtual Machines.

![img](./assets/gifs/intro.gif)

![GitHub issues](https://img.shields.io/github/issues-raw/pspiagicw/qemantra?logoColor=%23ffb86c&style=for-the-badge)

<a id="org7e245cb"></a>

## Table of Contents

1.  [Qemantra](#orgd792ca5)
    1.  [Motivation ?](#org7e245cb)
    2.  [Installation ?](#org1a144a1)
    3.  [Features](#org18f5296)
    4.  [Contributing](#orgfacc51e)

## Motivation ?

Virtualbox is good, but it has a QT interface and it's command line inteface is hectic at best.
QEMU has no official (or good) frontend , the command line interface is mature , but has no central managing solution.

Qemantra aims to become a simple and convinient way to manage Virtual Machines. It is designed for the casual Virtualizer
who struggles to control their VM's using the command line.


<a id="org1a144a1"></a>

## Installation ?

Currently you can only build from source.

-   Clone the repository
-   `make build` to build the binary.
-   Move the binary to a folder in your $PATH variable.


<a id="org18f5296"></a>

## Features

-   Create , list and run virtual machines.
-   Create configurations to run on demand.
-   Use features of QEMU like KVM , multiple architecture support etc.


<a id="orgfacc51e"></a>

## Contributing

Anybody is welcome to contribute , check out the `contribution guide` for more information.
For bugs and feature requests , open a issue.

