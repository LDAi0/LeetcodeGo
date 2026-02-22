def f(n):
    if n==1: return 1
    if n==2: return 3
    else:
        return f(n-1)*n+f(n-2)*(n-1)
print("das")
print(f(5))