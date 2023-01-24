# Getting Started

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

