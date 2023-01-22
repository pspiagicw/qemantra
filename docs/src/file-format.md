# file format

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


