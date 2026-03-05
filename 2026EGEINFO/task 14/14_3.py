from string import printable

for x in printable[:16]:
    x1=int(f'1f3b{x}75',16)
    x2=int(f'5d{x}3b',16)
    if (x1+x2)%11==0:
        print('temp')
        print((x1+x2)/11)