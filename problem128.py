# Problem 128: Hexagonal tile differences

from math import sqrt

def prime_sieve(upper_limit):
	sieve = [True] * (upper_limit + 1)
	sieve[0] = False
	sieve[1] = False
	sqrt_upper_limit = sqrt(upper_limit)

	p = 2
	while True:
		if sieve[p] == True:

			if p <= sqrt_upper_limit:
				n = p*p
				while n <= upper_limit:
					sieve[n] = False
					n += p

		if p > sqrt_upper_limit:
			break

		p += 1
			
			
	return sieve


def get_starting_number_of_layer(layer):
	if layer == 1:
		return 1

	return 3*layer*layer - 9*layer + 8

def get_num_of_positions_by_direction(direction, layer):
	if direction == 0 or direction == 4:
		return 1
	if direction == 2 or direction == 6:
		return layer
	else:
		return layer-2

def normalize_location(layer, direction, pos):
	if pos == 0:
		direction = (direction - 1)%8
		pos = get_num_of_positions_by_direction(direction, layer)
	if pos == get_num_of_positions_by_direction(direction, layer) + 1:
		direction = (direction + 1)%8
		pos = 1

	return layer, direction, pos

def get_number_by_location(layer, direction, pos):
	if layer == 1:
		return 1

	layer, direction, pos = normalize_location(layer, direction, pos)

	# position is from 0 to 7
	number = get_starting_number_of_layer(layer) - 1

	for i in range(0, direction):
		if i%2 == 1:
			number += layer-2
		else:
			if i == 4 or i == 0:
				number += 1
			else:
				number += layer
	number += pos

	return number

def simple_abs(x):
	return -x if x < 0 else x

def check_if_prime(p):
	return 1 if sieve[p] else 0

def check_if_diff_prime(current_number, neighbor):
	return check_if_prime(simple_abs(current_number - neighbor))

def get_neighbors_pd(current_number, layer, direction, pos):
	if layer == 1:
		return 3

	count = 0
	
	if direction == 0: # North
		count += check_if_diff_prime(current_number, get_starting_number_of_layer(layer +1) - 1)
		count += check_if_diff_prime(current_number, get_starting_number_of_layer(layer-1))
		count += check_if_diff_prime(current_number, get_starting_number_of_layer(layer+1))
		count += check_if_diff_prime(current_number, get_starting_number_of_layer(layer+1) + 1)
		count += check_if_diff_prime(current_number, get_starting_number_of_layer(layer+2) - 1)
	elif direction == 7 and pos == layer-2: # Northeast
		count += check_if_diff_prime(current_number, get_starting_number_of_layer(layer))
		above = get_number_by_location(layer+1, direction, pos)
		count += check_if_diff_prime(current_number, above)
		count += check_if_diff_prime(current_number, above + 1)
		below = get_number_by_location(layer-1, direction, pos)
		count += check_if_diff_prime(current_number, below)
		count += check_if_diff_prime(current_number, get_starting_number_of_layer(layer) - 1)
	else:
		return 0
			
	return count

sieve = prime_sieve(10000000)

sequence = [1]

layer = 2

while len(sequence) < 2000:

	# north
	current_number = get_starting_number_of_layer(layer)
	if get_neighbors_pd(current_number, layer, 0, 1) == 3:
		sequence.append(current_number)

	# northeast right next to north
	if layer > 2:
		current_number = get_starting_number_of_layer(layer + 1) - 1
		if get_neighbors_pd(current_number, layer, 7, layer-2) == 3:
			sequence.append(current_number)

	layer += 1

print(sequence[-1])