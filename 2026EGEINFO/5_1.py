def to_any_base(n, base):
    if n == 0:
        return "0"
    digits = []
    while n > 0:
        remainder = n % base
        if remainder < 10:
            digits.append(str(remainder))
        else:
            # Для систем счисления больше 10 (A=10, B=11, ...)
            digits.append(chr(ord('A') + remainder - 10))
        n //= base
    return "".join(digits[::-1])

for N in range(1,1000):
    r=to_any_base(N,2)
    r1=int(r)
    for i in len(r):
        x+=r1%10
        r1//=10
    r
