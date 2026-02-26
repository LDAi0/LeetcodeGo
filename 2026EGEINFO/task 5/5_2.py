def to_base(n,base):
	digits="0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if n==0:
		return "0"
	res = ""
	while n>0:
		res=digits[n%base]+res
		n//=base
	return res

for n in range(1,1000):
	s=to_base(n,2)
	if n%2==0:
		s='1'+s+'0'
	else:
		s='11'+s+'11'
	r=int(s,2)
	if r>52:
		print(n)
		print(n)
		break