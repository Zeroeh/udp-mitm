This guide will walk you through setting up a headless raspberry pi for intercepting packets.

This guide assumes you have good knowledge and competence of linux and the command line.

This guide also assumes you know how to get the raspberry pi all configured so that you can login via ssh.

The raspbian version used in this guide was ``2019-04-08-raspbian-stretch-lite`` and was updated to kernel ``4.19.66-v7+`` after updating packages.

Note that "raspberry pi" and "intercept" will be used interchangeably here on out.

You should be following this guide in root mode by running ``sudo -i``

All commands assume they are being run as root. If a command explicitly has sudo added to it, then add the sudo.

- Once you have configured the headless intercept and connected to your desired access point, you need to install a program called "dnsmasq". Run ``apt-get install dnsmasq`` 

- Next you need to disable eth0 in dhcpcd. This is often touted as the "newer" way to get networking in debian but the old method works just as well and is less confusing. Run ``nano /etc/dhcpcd.conf`` and at the very bottom of the file add ``denyinterfaces eth0``. Save and exit.

- For this next step, you will have to edit ``/etc/network/interfaces`` to create the static configuration for the ethernet interface. See the [example](https://github.com/Zeroeh/udp-mitm/blob/master/raspberrypi/interfaces.png) configuration that I use. I stripped some of the comments to make it easier to see. Yours should look similar, but feel free to customize the subnet to your liking. 

- If you haven't already, make sure your ``/etc/wpa_supplicant/wpa_supplicant.conf`` configuration is good and working. If you were able to install dnsmasq you should already be set, but it doesn't hurt to double check everything.

- Next, open up ``/etc/dnsmasq.conf`` for editing and add a configuration similar to [this](https://github.com/Zeroeh/udp-mitm/blob/master/raspberrypi/dnsmasq.png). You can change the "server" to any DNS server of your choosing.

- Almost done! Now we need to enable packet forwarding and set up our iptables rules. To enable packet forwarding, run the following commands:
  * ``sudo sysctl net.ipv4.ip_forward=1``
  * ``sudo echo "1" > /proc/sys/net/ipv4/ip_forward``
  * Open ``/etc/sysctl.conf`` and uncomment the line with ``net.ipv4.ip_forward=1`` if it is not already uncommented.
  
- Now we need to set up the iptables rules. To make this simple, I've created a setup script that does everything automatically. It can be downloaded [here](https://raw.githubusercontent.com/Zeroeh/udp-mitm/master/raspberrypi/setup.sh). Simply ``wget`` the script and execute it. You might have to ``chmod +x`` to make it executable. Once executed, you should be able to disable wifi in your host machine and get internet via the ethernet from the intercept. If you're able to get internet, reboot the intercept and apply the iptables rules again. If you can't get internet, reboot anyways.

To set up the actual proxy, follow the main guide at the root of the repository. If the host machine can get internet from the intercept, you can skip steps 1 through 3.

Any issues? Comments? Requests? Post an issue and I'll look into it. I wrote this guide based on memory so there very well could be a few issues or errors.

# Getting setup.sh to run at boot
It is extremely annoying to have to run the setup script to apply the iptables rules after each reboot. Especially if you dont have the luxury of a USB TTY cable to make things quick.

- To run the script at boot, open up ``/etc/rc.local`` and at the end of the file, but before the ``exit 0`` line, add the command to run your script. Mine is ``./root/route_scripts/setup.sh``
- Reboot and see if you are able to get internet on the host without manually running the setup script.


