#!/usr/bin/env python3

# modified script from https://github.com/lettier/ntpclient/blob/master/source/python/ntpclient.py

# this script is for testing udp rerouting through iptables

import sys
import socket
import struct
import time

def main():
	while True:
		host = "time.apple.com" #ntp server time.apple.com 17.253.20.125
		port = 123
		read_buffer = 1024
		address = (host, port)
		data = b"\x1b" + 47 * b"\0"
		epoch = 2208988800
		client = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
		client.sendto(data, address)
		data, address = client.recvfrom(read_buffer)
		t = struct.unpack("!12I", data)[10]
		t -= epoch
		print("Time = %s" % time.ctime(t))
		time.sleep(60)

if __name__ == "__main__":
	main()
