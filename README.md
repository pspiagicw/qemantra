# `qemantra`

`qemantra` is a tool to manage QEMU/KVM virtual machines.

<!-- TOC start (generated with https://github.com/derlin/bitdowntoc) -->

- [`qemantra`](#qemantra)
   * [Features](#features)
   * [Dependencies](#dependencies)
   * [Installation](#installation)
   * [Config](#config)
   * [Usage](#usage)
      + [`create`](#create)
      + [`run`](#run)
      + [`list`](#list)
      + [`rename`](#rename)
      + [`edit`](#edit)
   * [UEFI](#uefi)
   * [Similar Projects](#similar-projects)
   * [Contribution](#contribution)

<!-- TOC end -->

## Features

- Use QEMU without a graphical interface.
- Single static binary, no dependencies.
- No fuss, run VM's with a single command.

## Who is it for ?

`qemantra` is a opionated tool.

`qemantra` is 
- A simple tool for simple uses. It is for the casual virtualizer.
- Will probably never support highly complex features. 

`qemantra` is not
- A complex VM management tool.
- A performant or efficient tool.
- A production/enterprise tool. 

## Dependencies

- `qemu-system-*` binaries.

Mostly packaged with `qemu-full` (Arch/Debian).

- `ovmf` (*Optional*): for UEFI. See [here](#uefi).

### Debian

```sh
sudo apt install qemu-system-x86 qemu-system-sparc qemu-system-ppc qemu-system-arm
```

```sh
sudo apt install qemu-full
```

### Arch

```sh
sudo pacman -S qemu-system-x86 qemu-system-arm
```

```sh
sudo pacman -S qemu-full
```

## Installation

- You can download a binary from the releases section on [GitHub](https://github.com/pspiagicw/qemantra/releases).

If you have the `Go` compiler installed, you can install using this command.

```sh
go install github.com/pspiagicw/qemantra@latest
```

If you use [`gox`](https://github.com/pspiagicw/gox), you can also run.

If can also clone the repo and compile it manually.

```sh
git clone https://github.com/pspiagicw/qemantra
cd qemantra
go build .

# Or
groom build
```

## Config

- You will need a config to get started. 
- You can create a config at `/home/<username>/.config/qemantra/config.toml`
- It should have the following content, change it accordingly

```toml
imageDir = "~/.local/share/qemantra/images"
machineDir = "~/.local/share/qemantra/machines"
```

## Usage

###  `create`
- You can create virtual machines using the `qemantra create` command.
  
```sh
qemantra create
```

- This will start a interactive prompt to ask details about the VM.

![create](gifs/create.gif)

You will need to enter 
  - A valid name
  - A valid CPU core Count
  - A valid RAM size

> [!NOTE]
> Attaching a disk is optional. 

![disk](gifs/disk.gif)

If a disk is requested, you will need to provide
  - A disk name
  - A disk format
  - A disk size


### `run`

- You can run machines using `qemantra run`
- Running the virtual machine uses QEMU to run the machine using the given details.

```sh
qemantra run [FLAGS]
```
![run](gifs/run.gif)

### Flags

#### `--boot`

This allows to select a different boot option.

The different boot options available are
- `iso` This boots the ISO.
- `menu` This enables the menu, interactively choose the disk to boot.

#### `--external`

Provide a external disk to attach. This would be attached in addition to the existing disk (if present).

#### `--kvm`

Enable or disable KVM. By default `kvm` is enabled. To disable use `-kvm false`.

> [!WARNING]
> Features such as SMP depend on KVM.

#### `--iso`

Provide ISO disk to attach.

#### `--uefi`

Provide a OVMF file to use as UEFI bios.

### `list`

You can list machines using `qemantra list`.

```sh
qemantra list 
```

![list](gifs/list.gif)

### `rename`

You can use the `qemantra rename` command to rename any previously created machine.

```sh
qemantra rename 
```

![rename](gifs/rename.gif)


### `edit`

- You can use the `qemantra edit` command to edit any previously created machine.
- It would show prompts to change the details of any given VM.
- Shows current settings in brackets.

> [!IMPORTANT]
>  Edit is different from `rename`. For changing the name, see [rename](#rename).

```sh
qemantra [GLOBAL] edit [OPTIONS]
```

![edit](gifs/edit.gif)

## UEFI

- UEFI support is not natively bundled with QEMU, you need to install a external package using your package manager.
- You will need to install `ovmf` package.
- When running a virtual machine, point to the appropriate (.fd) file.
- In most systems, this would be installed in `/usr/share/ovmf/OVMF.fd`

## Similar Projects

- [libvirt](libvirt.org) (GUI and CLI)
- [VBoxManage](https://docs.oracle.com/en/virtualization/virtualbox/7.0/user/vboxmanage.html#vboxmanage) (CLI for VirtualBox).
- [Quickemu](https://github.com/quickemu-project/quickemu) (CLI, written in Bash)

## Contribution

Star the project on [GitHub](https://github.com/pspiagicw/qemantra) if you like it!

Anyone is free to contribute to the project, either by raising a issue or opening a PR.
