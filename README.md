# udp-mitm
Using NAT to MITM specific UDP traffic for traffic shaping purposes. The intent of this was to capture udp packets externally and edit them before they reach their destination.

# Instructions
Host = the machine running the target game/application. This machine is air gapped and connected via ethernet to the intercept.
Intercept = the machine offering the tethered internet connection to the host. This machine is running the mitm proxy application.
If you do not understand the above 2 concepts or want to see visually, please see "diagram1.png" in the repo

(These notes are currently linux specific. Other OS details will come soon when I get more testing done)
1. Set up the intercept and host machines by going to ethernet settings and sharing the connection. See how to do this [here](https://askubuntu.com/questions/359856/share-wireless-internet-connection-through-ethernet) or [here](https://askubuntu.com/questions/22835/how-to-network-two-ubuntu-computers-using-ethernet-without-a-router), or you can just google "linux share internet connection through ethernet"
2. If the above instructions from the links dont work, try deleting *ALL* ethernet profiles on both intercept and host machines and trying step 1 again.
3. On the host machine, shut off or disconnect from wifi and see if you can get internet, if not, refer to step 2 again. If it succeeds, continue.
Note: set the dstport for both scripts to "123" (NTP)
4. On the intercept machine, run the "intercept-setup.sh" script. Arguments for this script will be dst port and external ntp server. (You can use the default in the script but you can change it to any ntp server ip)
5. On the host machine, run the "host-setup.sh" script. Arguments are dst port and dst ip (dst ip is the local ethernet router ip aka your laptop/pi)
6. Start the ``main.go`` NTP server script with ``go run main.go``
7. Run the python NTP client with ``python3 test_script.py``
8. If everything works correctly, the python script should print out "Time = ..." repeatedly every 20 seconds. If it doesn't work, you'll have to do some debugging on your own, sorry.


# Issues
- Still need to get the bash script working
- The example python and ntp mitm application stops transmitting packets after a bit. This is probably some protection mechanism on apples end however.

# Todo
- Figure out how to NAT traffic without touching the hosts firewall/iptable rules (2nd interceptor?)
