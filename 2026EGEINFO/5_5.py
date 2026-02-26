for n1 in range(0,10000):
    n = bin(n1)[2:]
    if n.count('1')%2==0:
        n='10'+n[2:]
        n=n+'1'
    else:
        n=n+'11'
        n='1'+n[2:]
    if int(n,2)>=100:
        print("adsad")
        print(n1)
        break
