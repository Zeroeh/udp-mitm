# udp-mitm
Using NAT to MITM specific UDP traffic for traffic shaping purposes. The primary purpose of this project is to capture UDP traffic from games and process/edit them before they reach their destination. Since UDP is connectionless, unlike TCP which can be rerouted via hosts file, it needs special handling such as a specific NAT reroute or an application with a pcap hook to capture and edit the packets en route.

# Features
- Stealth. Be able to hide from any anticheat that would otherwise be looking for pcap hooks, firewall rules, and DNS reroutes (hosts file).
- You make the rest. The included test application is enough to get started making your hooks / features.

# Requirements
- Computer running the target application
- Second computer used as the intercept, this can be a laptop or raspberry pi, as long as it runs linux and has an ethernet port and wifi capabilities.
- Golang
- Sudo priveleges

# Instructions
- Host = the machine running the target game/application. This machine is air gapped and connected via ethernet to the intercept.
- Intercept = the machine offering the tethered internet connection to the host. This machine is running the mitm proxy application.

If you do not understand the above 2 concepts or want to see visually, please see [diagram1.png](https://github.com/Zeroeh/udp-mitm/blob/master/diagram1.png).

0. Download the [python test script](https://github.com/Zeroeh/udp-mitm/blob/master/test/test_script.py) to the host machine and the [NTP golang script](https://github.com/Zeroeh/udp-mitm/blob/master/test/main.go) to the intercept. When the steps refer to the bash script to run, please look [here](https://github.com/Zeroeh/udp-mitm/blob/master/test/start_reroute.sh) for the script.
1. Set up the intercept and host machines by going to ethernet settings and sharing the connection. See how to do this [here](https://askubuntu.com/questions/359856/share-wireless-internet-connection-through-ethernet) or [here](https://askubuntu.com/questions/22835/how-to-network-two-ubuntu-computers-using-ethernet-without-a-router), or you can just google "linux share internet connection through ethernet"
2. If the above instructions from the links dont work, try deleting *ALL* ethernet profiles on both intercept and host machines and trying step 1 again. If it still doesn't work, try flushing the iptables with ``iptables --flush && iptables -t nat --flush`` on the intercept and reboot. If it *still* doesn't work, you *may* need to enable ipv4 forwarding. You can do this with ``sudo sysctl net.ipv4.ip_forward=1`` which *should* apply every reboot and to apply the change immediately do ``sudo echo "1" > /proc/sys/net/ipv4/ip_forward``. Do this on the intercept.
3. On the host machine, shut off or disconnect from wifi and see if you can get internet, if not, refer to step 2 again. If it succeeds, continue. (you can use ifconfig to see which IP addresses are assigned to eth0, which they should be if everything is working)
- Note: set the dstport for both scripts to "123" (NTP)
4. On the intercept machine, run the "start_reroute.sh" script. Arguments for this script will be: server game port, mitm port, intercept ip on ethernet, and intercept ip base cidr. (example: ``sudo ./start_reroute.sh 5056 5555 10.42.0.1 10.42.0.0/24``)
5. (THIS STEP NO LONGER REQUIRED SINCE THE UPDATE) On the host machine, run the "host-setup.sh" script. Arguments are dst port and dst ip (dst ip is the local ethernet router ip aka your laptop/pi)
- Note: use ``iptables -t nat -L -n -v`` to double check that the iptables rules were applied. There should be a rule in OUTPUT and a rule in PREROUTING.
6. Start the ``main.go`` NTP server script with ``go run main.go``
7. Run the python NTP client with ``python3 test_script.py``
8. If everything works correctly, the python script should print out "Time = ..." repeatedly every 60 seconds. You can see example outputs in the [success](https://github.com/Zeroeh/udp-mitm/tree/master/success) directory. If it doesn't work, then you are on your own, sorry.

# Issues
- The example python and ntp mitm application stops transmitting packets after a bit. This is probably some protection mechanism on apples end however.
- You may experience some cross talk from the intercept coming from your router on heavily used UDP applications. This can be negated by changing the iptables script to use the ethernet interface only.

# Todo
- Maybe make the bash script a little more user friendly by cutting down the arguments a little bit?
- Make a script that removes the iptable rules.
