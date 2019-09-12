#!/bin/bash

# setup ip forwarding on raspberry pi. (tested & working on stretch. Kernel: 4.19.66-v7+)
# this script will enable packet forwarding between eth0 and wlan0 (devices connected via ethernet will get internet through the pi's wifi)

if [[ $EUID -ne 0 ]]; then
	echo "This script requires sudo/root privileges"
	exit 1
fi

sudo iptables -F #flush all existing rules in the normal table
sudo iptables -t nat -F #flush all existing rules in the nat table
sudo iptables -t nat -A POSTROUTING -o wlan0 -j MASQUERADE
sudo iptables -A FORWARD -i wlan0 -o eth0 -m state --state RELATED,ESTABLISHED -j ACCEPT
sudo iptables -A FORWARD -i eth0 -o wlan0 -j ACCEPT

echo "If no messages were displayed before this, you're good to go :)"

