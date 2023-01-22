# create

You can create virtual machines using the `qemantra create` command.

## Synopsis

```sh
qemantra [GLOBAL] create [OPTIONS]
```

Currently no options are available.

This will start a interactive prompt to ask details about the VM.

</br>

![create](./gifs/create.gif)

</br>

**Note: Currently no options are available for qemantra create**

## Details

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




