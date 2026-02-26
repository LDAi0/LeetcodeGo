from itertools import *
cnt=0
bad=['50','05','52','25','45','54','65','56']
for i in product('0123456', repeat=5):
    a=''.join(i)
    if a[0]!='0' and a.count('5')==1 and all(x not in a for x in bad):
        cnt+=1
print('adasd')
print(cnt)