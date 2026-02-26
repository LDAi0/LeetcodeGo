def to_base(n,base):
    digits="0123456789ABCDEFGHI"
    if n==0:return '0'
    res=""
    while n>0:
        res=digits[n%base]+res
        n//=base
    return res
for n in range(4,1000):
    n1=bin(n)[2:]
    if n%3==0:
        n1=n1+n1[-3]+n1[-2]+n1[-1]
    else:
        n1=n1[-1]+to_base((n%3)*3,2)
    r=int(n1,2)
    if r>=200:
        print("asdasd")
        print(n)
        break