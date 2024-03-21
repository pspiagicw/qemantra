# `qemantra`

`qemantra` is a tool to manage QEMU/KVM virtual machines.

## Features

- Use QEMU without a graphical interface.
- Single static binary, no dependencies.
- Blazingly fast.

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

## Dependencies

- `qemu-system-*` binaries.

Mostly packaged with `qemu-full` (Arch/Debian).

- `ovmf` (*Optional*): for UEFI. See [here].

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

![manual-install](./gifs/manual-install.gif)

## Getting Started

Let's boot a ISO using `qemantra`. You will need to have the `qemu-system-x86` binary installed on your system. See [here](#dependencies) for more information.

- You will need a config to get started. You can create a config at `/home/<username>/.config/qemantra/config.toml`
- It should have the following content, change it accordingly

```toml
imageDir = "~/.local/share/qemantra/images"
machineDir = "~/.local/share/qemantra/machines"
```
- For this example let's make a `Linux Mint` VM with `4G` of RAM and dedicating `3 cores` to the machine.
- We are not creating or attaching any disk to this VM, because we don't want to install it.

Run
```sh
qemantra create
```
- Answer questions accordingly

- `qemantra` should create a configuration file for your VM in  `machineDir`.
- For running the virtual machien and attach the ISO, use 

```sh
qemantra run -iso <path-to-iso>
```

- It should run QEMU in SDL (GUI) mode, booting the Linux Mint ISO.

![getting-started](./gifs/getting-started.gif)

## Usage

###  `create`
- You can create virtual machines using the `qemantra create` command.
  
```sh
qemantra [GLOBAL] create [OPTIONS]
```

- This will start a interactive prompt to ask details about the VM.

![create](./gifs/create.gif)

You will need to enter 
  - A valid name
  - A valid CPU core Count
  - A valid RAM size


- Attaching a disk is optional. 

> Creating a VM without a disk is very useful when only trying out a ISO without the intention of installing it.

![create-no-disk](./gifs/create-no-disk.gif)

If a disk is requested, you will need to provide
  - A disk name
  - A disk format
  - A disk size

### `run`

- You can run machines using `qemantra run`
- Running the virtual machine uses QEMU to run the machine using the given details.

```sh
qemantra [GLOBAL] run [OPTIONS]
```
![run](./gifs/run.gif)


#### `-boot`

This allows to select a different boot option.

The different boot options available are
- `iso` This boots the ISO.
- `menu` This enables the menu, interactively choose the disk to boot.

#### `-external`

Provide a external disk to attach. This would be attached in addition to the existing disk (if present).

#### `-kvm`

Enable or disable KVM. By default `kvm` is enabled. To disable use `-kvm false`.

> Features such as SMP depend on KVM.**

#### `-iso`

Provide ISO disk to attach.

#### `-uefi`

Provide a OVMF file to use as UEFI bios.

![run-iso](./gifs/run-iso.gif)

### `list`

You can list previously created machines using `qemantra list`.

```sh
qemantra [GLOBAL] list [OPTIONS]
```

![list](./gifs/list.gif)

### `rename`

You can use the `qemantra rename` command to rename any previously created machine.

```sh
qemantra [GLOBAL] rename [OPTIONS]
```

![rename](./gifs/rename.gif)

- You only need to choose a VM and provide a new name.
- If a machine already exists, it would inform you.

### `edit`

- You can use the `qemantra edit` command to edit any previously created machine.
- It would show prompts to change the details of any given VM.
- Shows current settings in brackets.

>  Edit is different from `rename`. For changing the name, see rename.

```sh
qemantra [GLOBAL] edit [OPTIONS]
```
![edit](./gifs/edit.gif)

## `UEFI`

> UEFI support is not natively bundled with QEMU, you need to install a external package using your package manager.

You will need to install `ovmf` package.

Then when running a virtual machine, point to the appropriate (.fd) file.

> In most systems, this would be installed in `/usr/share/ovmf/OVMF.fd`

![uefi](./gifs/run-uefi.gif)

## Similar Projects

- libvirt (GUI and CLI)
- VBoxManage (CLI for VirtualBox).
- Quickemu (CLI, has automatic ISO downloads)

## Contribution

⭐ Star the project on [GitHub](https://github.com/pspiagicw/qemantra) if you like it!

Anyone is free to Contribute to the project, either by raising a issue or opening a PR.
