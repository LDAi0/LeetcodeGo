print("x y z w")
for x in 0,1:
    for y in 0,1:
        for z in 0,1:
            for w in 0,1:
                if ((x==(w or y)) or ((w<=z) and (y<=w)))==0:
                    print(x, y, z, w)