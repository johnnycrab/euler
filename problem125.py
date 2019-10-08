# Problem 125

def is_palindromic(n):
	digits = list(str(n))
	k = 0
	l = len(digits)
	while k < l/2:
		if digits[k] != digits[l - 1 - k]:
			return False
		k += 1

	return True

below = 10**8

square_sums = [1]
for n in range(2, (10**4)+1):
	prev_sum = square_sums[n-2]
	square_sums.append(prev_sum + n*n)


palindromic_sum = 0 # for 1 = 1*1
palindromes = []

for i in range(1, len(square_sums)):
	if square_sums[i] < below and is_palindromic(square_sums[i]) and square_sums[i] not in palindromes:
			palindromes.append(square_sums[i])
			palindromic_sum += square_sums[i]

	for j in range(0, i-1):
		n = square_sums[i] - square_sums[j]
		if n < below and is_palindromic(n) and n not in palindromes:
			#print(n, i, j)
			palindromes.append(n)
			palindromic_sum += n

print(palindromic_sum)
