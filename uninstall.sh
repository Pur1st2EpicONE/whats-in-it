#!/bin/bash

if [[ $EUID -ne 0 ]]; then
  echo "Error — the script must be run with root permissions. Try sudo bash uninstall.sh"
  exit 1
fi

sudo rm /usr/local/bin/wii
sudo rm -rf /etc/whats-in-it

echo done
