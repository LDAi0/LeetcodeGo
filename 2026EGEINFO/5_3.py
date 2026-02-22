for n in range(0,256):
    a = bin(n)[2:]
    b="".join(['1'if bit =='0' else '0'for bit in a])
    y=int(b,2)
    if (n-y)==133:
        print("dasdawd")
        print(n)
        break