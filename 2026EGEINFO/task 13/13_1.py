from ipaddress import ip_network

net = ip_network("191.128.66.93/255.192.0.0", 0)
print("govno")
print(net.broadcast_address)