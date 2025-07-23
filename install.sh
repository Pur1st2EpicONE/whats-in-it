#!/bin/bash

if [[ $EUID -ne 0 ]]; then
  echo "Error — the script must be run with root permissions. Try sudo bash install.sh"
  exit 1
fi

mkdir -p /etc/whats-in-it
mkdir -p /etc/whats-in-it/certs

echo | openssl s_client -showcerts -connect gigachat.devices.sberbank.ru:443 </dev/null 2>/dev/null | openssl x509 -outform PEM > /etc/whats-in-it/certs/first.pem
echo | openssl s_client -showcerts -connect ngw.devices.sberbank.ru:9443 </dev/null 2>/dev/null | openssl x509 -outform PEM > /etc/whats-in-it/certs/second.pem

go build -o wii cmd/*
mv wii /usr/local/bin/
cp ./configs/config.yaml /etc/whats-in-it/
chmod 777 /etc/whats-in-it/config.yaml

echo done
