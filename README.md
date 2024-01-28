<img alt="qemantra icon" height="100" width="100" src="./assets/qemantra.svg" align="left" style="padding-right: 1em"> *qemantra is a cli application designed to create, run and manage virtual machines using <a href="qemu"> QEMU/KVM </a> and implemented in <a href="">Go</a>*

</br>

![carbon](./docs/src/gifs/intro.gif)

</br>


<!-- Visit the [website](https://qemantra.pspiagicw.xyz) for more information.
## Installation

See [here](https://qemantra.pspiagicw.xyz). -->


 `qemantra` is a command-line tool for creating and managing QEMU Virtual Machines. 

 QEMU is better and sometimes faster than VirtualBox , but does not have any 
 command-line central managing solution. This tool aims to manage and run your virtual machines using `QEMU`. 

 ![img](./docs/src/gifs/intro.gif) 

 ## Features 

 - Create , list and run virtual machines. 
 - Create configurations to run on demand. 
 - Use features of QEMU like KVM , multiple architecture support etc. 
 - Features like UEFI available in a simple flag! 

 ## Examples 
 ### Running a Virtual Machine 
 ![img](./docs/src/gifs/run-iso.gif) 

 ### Run a machine with boot menu 
 ![img](./docs/src/gifs/run.gif) 

 ### Create a Machine 
 ![img](./docs/src/gifs/create-no-disk.gif) 

 ### Create a Machine with Disk 
 ![img](./docs/src/gifs/create.gif) 

 <!-- ### Create a Image 
 ![img](./docs/src/gifs/create-img.gif)  -->

 ### Usage 
 Call `qemantra` in your terminal. 
 ```sh 
 $ qemantra 
 ``` 

 #### Check 
 Run `qemantra check` for checking for dependencies and configuration. 

 Highly recommended before using `qemantra`!. 

 #### Create Machine 

 The `qemantra create-machine` subcommand provides functionality to create machines. 
 | Option        | Description                                             | 
 |---------------|---------------------------------------------------------| 
 | `--name`      | Name of the machine                                     | 
 | `--no-disk`   | Don't create a disk                                     | 
 | `--disk-name` | Name of the disk(Not applicable when using `no-disk`)   | 
 | `--disk-size` | Size of the disk(Not applicable when using `disk-size`) | 
 | `--cpu-cores` | Cores to provide to the RAM.                            | 
 | `--mem-size`  | RAM to provide to the VM                                | 

 #### Running a machine 

 The `qemantra run` subcommand provides functionality to run a virtual machine. 
 | Option           | Description                                                                      | 
 |------------------|----------------------------------------------------------------------------------| 
 | `--name`         | Name of the machine                                                              | 
 | `--iso`          | Path to the ISO(Relative path works)                                             | 
 | `--disk`         | Disk name to add to boot order(Should be in default qemantra directory)          | 
 | `--externaldisk` | Path to a external disk to add to boot order(Any disk , not managed by qemantra) | 
 | `--boot`         | Boot options while starting the VM                                               | 
 | `--uefi`         | Enable UEFI support(Requires `OVMF` to be installed)                             | 
 | `--no-kvm`       | Disables KVM(Enabled by default)                                                                                 | 

 The boot options can be either `menu` which provides a menu to choose between boot devices. Or you can use `iso` option to directly boot the given iso. 

 #### Renaming a machine 
 The `qemantra rename` command is used to rename a existing virtual machine. 
 If `qemantra run` has no arguments , it will execute the last machine which was booted. 

      
 #### List machines 
 Use `qemantra list` to list currently configured machines. Use `--verbose` option to list more information about the VM(Mem , Cpu etc). 

 You can list the images managed by `qemantra` by using `--images` option to `qemantra list`. 

 ## Installation ? 

 ### Installation using Go 
 You can install using Go by using 
 ```sh 
 go install github.com/pspiagicw/qemantra 
 ``` 

 This requires `$GOBIN` to be in our `$PATH` variable. 

 ### Installation using script 
 You can also use the autoinstall script.  
 You should generally not execute scripts from the internet , so check the script 
 ```sh 
 curl https://raw.githubusercontent.com/pspiagicw/qemantra/main/scripts/install.sh | bash 
 ``` 

 ### Installing manually 
 - Download the static binary from the [release](https://github.com/pspiagicw/qemantra/releases) page. 
 - Move the binary to a location on your `$PATH` variable. 

 ## Roadmap ? 

 These are the major features planned to be added to `qemantra`. Other features are welcome to be discussed. 

 - [x] Running virtual machines 
 - [x] Creating virtual machines. 
 - [x] Can use memory and cpu cores. 
 - [x] Can use iso while running 
 - [x] Can use different disks while running 
 - [x] Can use external disk while running 
 - [x] Can use boot options 
 - [x] Make a auto install script. 
 - [x] Support UEFI using OVMF 
 - [ ] Control logging functionality(`-v` / `-vv` and `-vvv`) 
 - [ ] Add to major repositories(Debian , AUR , Gentoo) 
 - [ ] Configuration changes using ENVIRONMENT VARIABLES 
 - [ ] Multiple architecture support 
 - [ ] Suppport easy clipboard sharing 
 - [ ] Support easy shared folder support 
 - [ ] List currently running machines. 

 ## Motivation ? 

 Virtualbox is good, but it has a QT interface and it's command line inteface is hectic at best. 
 QEMU has no official (or good) frontend , the command line interface is mature , but has no central managing solution. 

 `qemantra` aims to become a simple and convinient way to manage Virtual Machines. It is designed for the casual Virtualizer. 

 ## Contributing 

 Anybody is welcome to contribute! 
   
 `qemantra` is written in Golang , so Golang developers can contribute in the technical aspect. 
 If you want to contribute non-technically , then too you are welcome! 
 There are lot's of work in documentation and other aspects! 
 For bugs and feature requests , open a issue. 

## Getting help

- You can look at the issue section, seach for a similar issue first. If you feel your issue is not found, file a issue.
