# Problem 119: Digit power sum


def digit_sum(n):
	s = 0
	while n%10 != n:
		d = n%10
		s += d
		n = (n-d)//10
	s += n
	return s


highest_digit_sum = 15 * 9
highest_num = 10**16

seq = []

for dig_sum in range(2, highest_digit_sum + 1):
	power = 1
	exponent = 0
	while power*dig_sum <= highest_num:
		exponent += 1
		power *= dig_sum
		if power > 10 and digit_sum(power) == dig_sum:
			seq.append(power)

seq.sort()

print(len(seq))
print(seq[29])
