# `qemantra`

`qemantra` is a tool to manage QEMU/KVM virtual machines.

## ✨ Features

- 🔥 Use QEMU without a graphical interface.
- 🧹 Single static binary, no dependencies.
- 🚀 Blazingly fast.

![carbon](./gifs/intro.gif)

## Who is it for ?

`qemantra` has a very simple usage.Provide a helpful CLI tool to manage VM's running on top of QEMU.
But because of it's simplicity it is not for everyone.

`qemantra` is 
- A simple tool for simple uses. It is for the casual virtualizer.
- Will probably never support highly complex features. 


`qemantra` is not
- A complex VM management tool.
- A production/enterprise tool. Project is still under heavy development.Use it for personal use only!

## ⚡️ Dependencies

- `qemu-system-*` binaries.

Mostly packaged with `qemu-full` (Arch/Debian).

- `ovmf` (*Optional*): for UEFI. See [here](/uefi.html).

### Debian

```sh
sudo apt install qemu-system-x86 qemu-system-sparc qemu-system-ppc qemu-system-arm
```

**OR** 

```sh
sudo apt install qemu-full
```

### Arch

```sh
sudo pacman -S qemu-system-x86 qemu-system-arm
```

**OR**

```sh
sudo pacman -S qemu-full
```

![dependencies](./gifs/dependencies.gif)


## Installation

### 🛠️ Binary Installation

You can download a binary from the releases section on [GitHub](https://github.com/pspiagicw/qemantra/releases).

Move it to a folder on your PATH.

Or use this [script](https://github.com/pspiagicw/qemantra/blob/main/scripts/install.sh).

Run this command to download the script and install it automatically.

```sh
curl https://raw.githubusercontent.com/pspiagicw/qemantra/main/scripts/install.sh | bash
```

This script would detect your system and download the appropriate binary and install it in  `/usr/local/bin/`.

If possible it would also install manpages and other meta information.

![binary-install](./gifs/binary-install.gif)

### 🏗️  Manual Installation

Manual installation is recommended when contributing, or making personal changes.

You will need some packages before you can run `qemantra`.See [here](/dependencies.html).

After installing the packages. 

- You can install the `go` compiler. 
- Clone the project and execute `groom build`

If you don't have [`groom`](https://github.com/pspiagicw/groom) installed, you can also run

```sh
go build main.go -o qemantra
```

This development binary would not have version info.

![manual-install](./gifs/manual-install.gif)


## Getting Started

Let's boot a ISO using `qemantra`. You will need to have the `qemu-system-x86` binary installed on your system. See [here](/dependencies.html) for more information.

After you have installed `qemantra` using one of the methods, we can create a virtual machine using the `qemantra create` command.

- For this example let's make a `Linux Mint` VM with `4G` of RAM and dedicating `3 cores` to the machine.

- We are not creating or attaching any disk to this VM, because we don't want to install it.

This can be achived by 2 simple steps.

Run
```sh
qemantra create
```

This would ask you a series of questions.
Answer them accordingly.

It should automatically create the `~/.qemantra` directory.

`qemantra` should create a configuration file for your VM in `/home/<username>/.qemantra/machines`.
The filename should be `<name>.json`.

For actually running the virtual machien and attach the ISO, use 

```sh
qemantra run -iso <path-to-iso>
```
It should run QEMU in SDL (GUI) mode, booting the Linux Mint ISO.

</br>

![getting-started](./gifs/getting-started.gif)

</br>

`qemantra` has a lot of configuration options, while creating and running virtual machines. Read the documentation to know more.


## Usage

### Create

You can create virtual machines using the `qemantra create` command.

#### Synopsis

```sh
qemantra [GLOBAL] create [OPTIONS]
```

This will start a interactive prompt to ask details about the VM.

</br>

![create](./gifs/create.gif)

</br>

#### Details

You will need to enter 
- A valid name
- A valid CPU core Count
- A valid RAM size

Attaching a disk is optional. 
Creating a VM without a disk is very useful when only trying out a ISO without the intention of installing it.

</br>

![create-no-disk](./gifs/create-no-disk.gif)

</br>


If a disk is requested, you will need to provide
- A disk name
- A disk format
- A disk size

### Run

You can run machines using `qemantra run`

Running the virtual machine uses QEMU to run the machine using the given details.

#### Synopsis

```sh
qemantra [GLOBAL] run [OPTIONS]
```
</br>

![run](./gifs/run.gif)

</br>

The options available are given below.

| Option | Description |
| -------| ------------|
| `-boot` | Define the boot order |
| `-external` | Add a external disk |
| `-iso` | Add a ISO to boot |
| `-kvm` | Enable kvm |
| `-uefi` | Path to UEFI bios file |

#### `-boot`

This allows to select a different boot option.

The different boot options available are
- `iso` This boots the ISO.
- `menu` This enables the menu, interactively choose the disk to boot.

#### `-external`

Provide a external disk to attach. This would be attached in addition to the existing disk (if present).

#### `-kvm`

Enable or disable KVM. By default `kvm` is enabled. To disable use `-kvm false`.

**Note: Features such as SMP depend on KVM.**

#### `-iso`

Provide ISO disk to attach.

</br>

![run-iso](./gifs/run-iso.gif)

</br>

#### `-uefi`

Provide a OVMF file to use as UEFI bios. See [here](/uefi.html) for more information.

### List

You can list previously created machines using `qemantra list`.

#### Synopsis

```sh
qemantra [GLOBAL] list [OPTIONS]
```

![list](./gifs/list.gif)

### Rename

You can use the `qemantra rename` command to rename any previously created machine.

#### Synopsis

```sh
qemantra [GLOBAL] rename [OPTIONS]
```

</br>

![rename](./gifs/rename.gif)

</br>

You only need to choose a VM and provide a new name. If a machine already exists, it would inform you.

### Edit

You can use the `qemantra edit` command to edit any previously created machine.
It would show prompts to change the details of any given VM. Shows current settings in brackets.

**Note: Edit is different from `rename`. For changing the name, see [rename](/rename.html)**

#### Synopsis

```sh
qemantra [GLOBAL] edit [OPTIONS]
```

</br>

![edit](./gifs/edit.gif)

</br>

#### Details

This follows the same information as [creating a disk](/create.html).

## Similar Projects

- libvirt (GUI and CLI)
- VBoxManage (CLI for VirtualBox).
- Quickemu (CLI, has automatic ISO downloads)

## `UEFI`

> UEFI support is not natively bundled with QEMU, you need to install a external package using your package manager.

You will need to install `ovmf` package.

Then when running a virtual machine, point to the appropriate (.fd) file.

In most systems, this would be installed in `/usr/share/ovmf/OVMF.fd`

</br>

![uefi](./gifs/run-uefi.gif)

</br>

## File Format

Each virtual machine is stored as JSON file.

All files are stored in `~/.qemantra/machines`.

**You should only edit the machine files when you know what are you doing.**.

*Example JSON file*

```json
{
  "name": "manjaro",
  "drivePath": "",
  "systemCommand": "qemu-system-x86_64",
  "memSize": "4G",
  "cpuCores": "5"
}
```

*VM with a disk attached*

```json
{
  "name": "Ubuntu",
  "drivePath": "/home/pratham/.qemantra/images/ubuntu.img",
  "systemCommand": "qemu-system-x86_64",
  "memSize": "4G",
  "cpuCores": "2"
}

```

## Contribution

⭐ Star the project on [GitHub](https://github.com/pspiagicw/qemantra) if you like it!

Anyone is free to Contribute to the project, either by raising a issue or opening a PR.
