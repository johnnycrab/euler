# Problem 124

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




primes = prime_sieve(2, 100000)

the_ns = range(1, 100001)
the_rads = [1]

def compare_rads(a, b):
	rad_a = the_rads[a-1]
	rad_b = the_rads[b-1]

	if rad_a > rad_b:
		return 1
	elif rad_a == rad_b:
		return a-b
	else:
		return -1

for i in range(1, len(the_ns)):
	n = the_ns[i]

	rad = 1
	temp = n
	prime_idx = 0
	prev_idx = -1
	
	while temp != 1:
		p = primes[prime_idx]

		if temp%p == 0:
			if prev_idx != prime_idx:
				rad *= p
				prev_idx = prime_idx
			temp /= p
		else:
			prime_idx += 1

	the_rads.append(rad)

the_ns_sorted = sorted(the_ns, cmp=compare_rads)

print(the_ns_sorted[9999])