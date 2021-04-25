#!/usr/bin/env bash
if [ `whoami` != 'root' ]
  then
    echo -e "\e[91mTry With Sudo"
    exit
fi
GPIO=$1
if [[ ! -n $GPIO ]]
then
  echo "Usage :
  ./setup.sh GPIO_PIN_NUM
  "
  exit
fi

if [[ -d /sys/class/gpio/gpio$GPIO ]]
then
  echo "Config Exist"
  exit
fi

echo $GPIO > /sys/class/gpio/export
echo out > /sys/class/gpio/gpio$GPIO/direction

echo -e "\n\e[96m Done!"

echo "make sure to put $GPIO at the beginning of main.go file"
