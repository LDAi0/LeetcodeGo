from ipaddress import ip_network
net=ip_network('192.168.64.176/255.255.255.240',0)
cnt=0
for ip in net:
    if bin(int(ip))[2:].count('1')%2==0:
        cnt+=1
print('dasd')
print(cnt)