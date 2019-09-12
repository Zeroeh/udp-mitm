#!/bin/bash

if [[ $EUID -ne 0 ]]; then
	echo "This script requires sudo/root priveleges"
	exit 1
fi

if [[ $# -ne 4 ]]; then
	echo "This script requires 4 arguments: Application endpoint port, mitm port, tether intercept ip, tether intercept ip base cidr"
	echo "Example: ./%s 123 5555 10.42.0.1 10.42.0.0/24"
	echo "Possible intercept IPs: (Usually starts with 10.42.0.x)"
	ourip=`hostname -I`
	echo $ourip
	exit 1
fi

sudo iptables -t nat -A OUTPUT -p udp --sport $1 -j DNAT --to-destination 127.0.0.1:$2
sudo iptables -t nat -A PREROUTING -p udp --src $4 --dport $1 -j DNAT --to-destination $3:$2
echo "sudo iptables -t nat -A OUTPUT -p udp --sport $1 -j DNAT --to-destination 127.0.0.1:$2"
echo "sudo iptables -t nat -A PREROUTING -p udp --src $4 --dport $1 -j DNAT --to-destination $3:$2"
echo "Everything looks good here. Time to play some games :)"
