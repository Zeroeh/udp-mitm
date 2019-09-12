#!/bin/bash

if [[ $EUID -ne 0 ]]; then
	echo "This script requires root privileges"
	exit 1
fi

# run through twice for the game im testing, but you can change this to be variable based on program args
sudo iptables -t nat -D PREROUTING 1 && sudo iptables -t nat -D OUTPUT 1
sudo iptables -t nat -D PREROUTING 1 && sudo iptables -t nat -D OUTPUT 1
