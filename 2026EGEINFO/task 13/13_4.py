from ipaddress import ip_network
ip1='95.24.2.9'
ip2='95.24.3.10'
res=[]

for mask in range(16,33):
    net1=ip_network(f'{ip1}/{mask}',strict=False)
    net2=ip_network(f'{ip2}/{mask}',strict=False)
    if net1==net2:
        count=0
        for ip in net1:
            bin_ip=bin(int(ip)).zfill(32)
            if bin_ip.count('0')%2==0:
                count+=1
        res.append(count)
print(min(res))
