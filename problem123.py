# Problem 123: Prime square remainders
from math import sqrt

def prime_sieve(lower_limit, upper_limit):
	sieve = [True] * (upper_limit + 1)
	sqrt_upper_limit = sqrt(upper_limit)
	primes = []

	p = 2
	while True:
		if sieve[p] == True:
			if p >= lower_limit:
				primes.append(p)

			if p <= sqrt_upper_limit:
				n = p*p
				while n <= upper_limit:
					sieve[n] = False
					n += p

		if p >= upper_limit:
			break

		p += 1
			
			
	return primes

primes = prime_sieve(2, 10**6)


for n in range(1, len(primes), 2):
	p_n = primes[n-1]


	remainder = (2*p_n*n) % (p_n*p_n)

	if remainder > 10**10:

		print(n)
		break