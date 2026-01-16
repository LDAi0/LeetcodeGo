print("x y z")
for x in 0,1:
    for y in 0,1:
        for z in 0,1:
            if ((x or y)<=(z==x)) == 0:
                print(x,y,z)