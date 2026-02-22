def to_base(n,base):
	digits="0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if n==0:
		return "0"
	res = ""
	while n>0:
		res=digits[n%base]+res
		n//=base
	return res
res=0
for s in to_base(4**2020+2**2017-15,2):
	if s=='1':
		res=res+1
		
print("ahahahah")
print(res)