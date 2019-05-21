# udp-mitm
Using NAT to MITM specific UDP traffic for traffic shaping purposes. The primary purpose of this project is to capture UDP traffic from games and process/edit them before they reach their destination. Since UDP is connectionless, unlike TCP which can be rerouted via hosts file, it needs special handling such as a special NAT routing or an application with a pcap hook to capture and edit the packets en route.

# Requirements
- Computer running the target application
- Second computer used as the intercept, this can be a laptop or raspberry pi, as long as it runs linux and has an ethernet port and wifi capabilities.
- Golang
- Sudo priveleges
- A working brain, for debugging

# Instructions
- Host = the machine running the target game/application. This machine is air gapped and connected via ethernet to the intercept.
- Intercept = the machine offering the tethered internet connection to the host. This machine is running the mitm proxy application.

If you do not understand the above 2 concepts or want to see visually, please see [diagram1.png](https://github.com/Zeroeh/udp-mitm/blob/master/diagram1.png).

0. Download the [python test script](https://github.com/Zeroeh/udp-mitm/blob/master/test/test_script.py) to the host machine and the [NTP golang script](https://github.com/Zeroeh/udp-mitm/blob/master/test/main.go) to the intercept. When the steps refer to the bash scripts to run, please look [here](https://github.com/Zeroeh/udp-mitm/blob/master/disregard.txt) for the commands to run. I don't have the bash scripts set up yet.
1. Set up the intercept and host machines by going to ethernet settings and sharing the connection. See how to do this [here](https://askubuntu.com/questions/359856/share-wireless-internet-connection-through-ethernet) or [here](https://askubuntu.com/questions/22835/how-to-network-two-ubuntu-computers-using-ethernet-without-a-router), or you can just google "linux share internet connection through ethernet"
2. If the above instructions from the links dont work, try deleting *ALL* ethernet profiles on both intercept and host machines and trying step 1 again. If it still doesn't work, try flushing the iptables with ``iptables --flush && iptables -t nat --flush`` and rebooting. If it *still* doesn't work, you *may* need to enable ipv4 forwarding. You can do this with ``sudo sysctl net.ipv4.ip_forward=1`` which *should* apply every reboot and to apply the change immediately do ``sudo echo "1" > /proc/sys/net/ipv4/ip_forward``. Do this on both the host and intercept.
3. On the host machine, shut off or disconnect from wifi and see if you can get internet, if not, refer to step 2 again. If it succeeds, continue. (you can use ifconfig to see which IP addresses are assigned to eth0, which they should be if everything is working)
- Note: set the dstport for both scripts to "123" (NTP)
4. On the intercept machine, run the "intercept-setup.sh" script. Arguments for this script will be dst port and external ntp server. (You can use the default in the script but you can change it to any ntp server ip)
5. On the host machine, run the "host-setup.sh" script. Arguments are dst port and dst ip (dst ip is the local ethernet router ip aka your laptop/pi)
- Note: use ``iptables -t nat -L -n -v`` to double check that the iptables rules were applied.
6. Start the ``main.go`` NTP server script with ``go run main.go``
7. Run the python NTP client with ``python3 test_script.py``
8. If everything works correctly, the python script should print out "Time = ..." repeatedly every 60 seconds. You can see example outputs in the [success](https://github.com/Zeroeh/udp-mitm/tree/master/success) directory. If it doesn't work, then you are on your own, sorry.


# Issues
- Still need to get the bash script working
- The example python and ntp mitm application stops transmitting packets after a bit. This is probably some protection mechanism on apples end however.
- The current implementation works flawlessly, but see the todo section as to why it could be better.

# Todo
- Figure out how to NAT traffic without touching the hosts firewall/iptable rules (2nd interceptor?). Having this method would be ideal as no anticheat would be able to read firewall rules. If anyone knows a fix or workaround (preferrably *without* adding extra hardware), submit an issue or pull request. The contribution will be greatly appreciated. Edit: I have the fix, just need to do some furthur testing and update the bash script. See disregard2.txt for the fix. Apply that and the other rule for the intercept machine.
