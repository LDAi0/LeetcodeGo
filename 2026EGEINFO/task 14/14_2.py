from string import printable
for p in range(10,36):
    for x in printable[:p]:
        for y in printable[:p]:
            for w in printable[:p]:
                for z in printable[:p]:
                    if y!='0' and w!='0':
                        x1 = int(f'{y}18{x}',p)
                        x2 = int(f'{w}{y}98',p)
                        if x1+x2==int(f'{x}{x}{z}4{y}',p):
                            print("sh")
                            print(x,y,z,w,p)
                            break
print("eshkere")
print(int("19b5",13))