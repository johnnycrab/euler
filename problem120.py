
s = 0

for a in range(3, 1001):

	a_sq = a*a

	r_max = 0

	for b in range(1, a):
		r = (2*a*b)%a_sq
		if r > r_max:
			r_max = r
	
	s += r_max

print(s)