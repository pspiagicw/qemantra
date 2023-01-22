# uefi

UEFI support is not natively bundled with QEMU, you need to install a external package using your package manager.

You will need to install `ovmf` package.

Then when running a virtual machine, point to the appropriate (.fd) file.

In most systems, this would be installed in `/usr/share/ovmf/OVMF.fd`

</br>

![uefi](./gifs/run-uefi.gif)

</br>







