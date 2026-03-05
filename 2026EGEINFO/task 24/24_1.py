s=open('C:\\Users\\abso\\Documents\\Visual Studio\\GoLeetcode\\2026EGEINFO\\task 24\\24_123123.txt').readline()
m=0
m2=0
for c in 'QRW': s=s.replace(c,'A')
for c in '124': s=s.replace(c,'1')

for i in range(1,len(s)):
    if s[i]!=s[i-1]:
        m2+=1
    else:
        m=max(m,m2)
        m2=0
print(m+1)
