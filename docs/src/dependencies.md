# ⚡️ Dependencies

- `qemu-system-*` binaries.

Mostly packaged with `qemu-full` (Arch/Debian).

- `ovmf` (*Optional*): for UEFI. See [here](/uefi.html).

</br>

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

</br>

![dependencies](./gifs/dependencies.gif)


