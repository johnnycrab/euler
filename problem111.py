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



#primes = prime_sieve(10**9, 10**10)
#print(len(primes))
a = [100] * (5* (10**9))