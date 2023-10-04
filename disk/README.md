```
$ sudo fdisk /dev/sda

Command (m for help): p

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1   *        2048     9437183     4717568   83  Linux

Command (m for help): d
Selected partition 1

Command (m for help): p

   Device Boot      Start         End      Blocks   Id  System

Command (m for help): n
Command action
   e   extended
   p   primary partition (1-4)
Select (default p): p
Partition number (1-4, default 1): 1
First sector (2048-10485759, default 2048):
Using default value 2048
Last sector, +sectors or +size{K,M,G} (2048-10485759, default 10485759):
Using default value 10485759

Created a new partition 1 of type 'Linux' and of size 10 GiB.
Partition #1 contains a ext4 signature.

Do you want to remove the signature? [Y]es/[N]o: N

Command (m for help): p

   Device Boot      Start         End      Blocks   Id  System
/dev/sda1            2048    10485759     5241856   83  Linux

Command (m for help): w
The partition table has been altered!

Calling ioctl() to re-read partition table.

WARNING: Re-reading the partition table failed with error 16: Device or resource busy.
The kernel still uses the old table. The new table will be used at
the next reboot or after you run partprobe(8) or kpartx(8)
Syncing disks.
```

Resize:

```
$ sudo resize2fs /dev/sda1
```
