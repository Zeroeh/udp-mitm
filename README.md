# udp-mitm
Using NAT to MITM specific UDP traffic for traffic shaping purposes. The intent of this was to capture udp packets externally and edit them before they reach their destination.

# Instructions
Host = the machine running the target game/application. This machine is air gapped and connected via ethernet to the intercept.
Intercept = the machine offering the tethered internet connection to the host. This machine is running the mitm proxy application.
If you do not understand the above 2 concepts or want to see visually, please see "diagram1.png" in the repo

1. 

# Issues
- Still need to get the bash script working

# Todo
- Figure out how to NAT traffic without touching the hosts firewall/iptable rules (2nd interceptor?)
