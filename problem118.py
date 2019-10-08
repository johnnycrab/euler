# Problem 118: Pandigital prime sets

# To make life easier we pre-calculated the possibilities of sets with 8-digit primes + 1-digit prime: 11483

from math import sqrt

count = 11483

def make_prime_digit_list(nums):
	new_list = [[],[],[],[],[],[],[],[]]

	for n in nums:
		digits = list(str(n))
		if "0" in digits or any(digits.count(x) > 1 for x in digits):
			continue

		new_list[len(digits)].append(digits)

	return new_list


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

def can_be_used(primes, digits):
	for d in primes:
		if d not in digits:
			return False
	return True

def remove_digits(digits, target):
	new = []
	for t in target:
		if t not in digits:
			new.append(t)
	return new

def check_set_possibility(free_digits, current_prime_length, current_prime_index):
	global prime_digit_list
	n_free_digits = len(free_digits)

	# if we come to zero, we are done
	if n_free_digits == 0:
		global count
		count += 1
		return

	# nothing to do here, there are no more possibilities
	if n_free_digits - current_prime_length < 0 or (free_digits == current_prime_length and current_prime_index >= len(prime_digit_list[current_prime_length])):
		return

	while n_free_digits - current_prime_length >= 0:
		current_prime_index += 1
		if current_prime_index == len(prime_digit_list[current_prime_length]):
			current_prime_length += 1
			current_prime_index = 0
			if current_prime_length == 8:
				return

		# get the prime
		prime_digits = prime_digit_list[current_prime_length][current_prime_index]
		if can_be_used(prime_digits, free_digits):
			check_set_possibility(remove_digits(prime_digits, free_digits), current_prime_length, current_prime_index)






primes = prime_sieve(2, 10000000)

prime_digit_list = make_prime_digit_list(primes)

check_set_possibility(['1', '2', '3', '4', '5', '6', '7', '8', '9'], 1, -1)


print(count)
#print(primes_without_double_digits[1:1000])