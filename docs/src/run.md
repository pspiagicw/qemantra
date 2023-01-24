# Run

You can run machines using `qemantra run`

Running the virtual machine uses QEMU to run the machine using the given details.

## Synopsis

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

## `-boot`

This allows to select a different boot option.

The different boot options available are
- `iso` This boots the ISO.
- `menu` This enables the menu, interactively choose the disk to boot.

## `-external`

Provide a external disk to attach. This would be attached in addition to the existing disk (if present).

## `-kvm`

Enable or disable KVM. By default `kvm` is enabled. To disable use `-kvm false`.

**Note: Features such as SMP depend on KVM.**

## `-iso`

Provide ISO disk to attach.

</br>

![run-iso](./gifs/run-iso.gif)

</br>

## `-uefi`

Provide a OVMF file to use as UEFI bios. See [here](/uefi.html) for more information.
