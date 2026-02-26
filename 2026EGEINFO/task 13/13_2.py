from ipaddress import ip_network
net = ip_network('112.160.0.0/255.240.0.0', 0)
count=0
for ip in net:
    n = bin(int(ip))[2:]
    if n.count('1')%5==0:
        count+=1
print("dawdasda")
print(count)