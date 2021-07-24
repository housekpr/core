# How to install

## Minimum required hardware to run the software

- [Raspberry Pi Zero W](https://www.raspberrypi.org/products/raspberry-pi-zero-w/), [Raspberry Pi 3 Model B+](https://www.raspberrypi.org/products/raspberry-pi-3-model-b-plus/) or newer.
- 8GB SD Card or bigger.
- 5V 2A Power Adapter.

Any other ARMv7 or ARMv8 device that can run Linux Kernel 5.10 or higher and has GPIO ports may work.

## End user installation guide

The instruction below assumes you are using RaspberryPi Zero/3/4 running latest stable Raspbian Minimal/Light installation.

### Prepare an OS user with minumum privileges

    adduser housekpr
    sudo usermod -aG gpio housekpr
    sudo su - housekpr 
    mkdir -p bin

## Installing Housekpr

    // TODO... wget github.com/release/.. ?




## IGNORE ME - DEVELOPER NOTES AND DRAFTS
    // Is that really needed? minimum permissions better
    // pi user is member of pi adm dialout cdrom sudo audio video plugdev games users input netdev gpio i2c spi
    sudo usermod -aG adm,dialout,cdrom,floppy,sudo,audio,dip,video,plugdev,netdev,lxd housekpr
    sudo visudo
    # Add NOPASSWD to sudo without pass -> %sudo   ALL=(ALL:ALL) NOPASSWD:ALL

## Prepare gpio group, if not there already
## to avoid https://github.com/fivdi/onoff/issues/3
## https://www.raspberrypi.org/forums/viewtopic.php?t=5185#p161013
sudo groupadd gpio
sudo usermod -aG gpio pavel
sudo chgrp -R gpio /sys/class/gpio
sudo chmod -R g+w /sys/class/gpio


## libraspberrypi-bin provides vcgencmd
## python3-gpiozero   provides py3 import for gpiozero and pinout command
## python3-pigpio     provides py3 import 
## network-manager    provides nmtui 
## pigpio             provides pigpiod (http://abyz.me.uk/rpi/pigpio)
## unzip              provides unzip
## make               provides make
## gcc                provides gcc
## g++                provides g++ (needed for npm install onoff)
## golang             provides go
apt install libraspberrypi-bin python3-gpiozero network-manager pigpio unzip make gcc g++ golang


## if pigpio is not available do manual install ass explained in 
## http://abyz.me.uk/rpi/pigpio/download.html
sudo apt install python3-setuptools
## get latest version from github?
## then download & install 
mkdir pigpio
cd pigpio
wget https://github.com/joan2937/pigpio/archive/refs/tags/v79.zip
unzip v79.zip
cd pigpio-79
make
sudo make install
## run the tests and make sure they pass!

## install node lts

### install nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.38.0/install.sh | bash

### install node lts
nvm install --lts
### install latest npm :wq
npm install -g npm
