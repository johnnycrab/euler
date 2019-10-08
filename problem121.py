# Problem 121

import math

def n_choose_k(n, k):
	return math.factorial(n) // (math.factorial(n-k) * math.factorial(k))

def n_choose_k_sets(n, k):
	n_sets = n_choose_k(n, k)

	resulting_sets = [[] for i in range(n_sets)]

	n_choose_k_sets_fill_slice(resulting_sets, 0, len(resulting_sets), n, k)

	return resulting_sets

def n_choose_k_sets_fill_slice(sets, from_idx, to_idx, n, k):
	last_num = 0
	if len(sets[from_idx]) > 0:
		last_num = sets[from_idx][-1]
	count_of_numbers_left = k - len(sets[from_idx])

	if count_of_numbers_left == 0 or last_num == n:
		return 

	fill_up_to_number = n - count_of_numbers_left + 1

	for i in range(last_num + 1, fill_up_to_number + 1):
		nums_left = n-i
		#print(i, nums_left, count_of_numbers_left)
		slots_needed = n_choose_k(nums_left, count_of_numbers_left - 1)
		for j in range(from_idx, from_idx + slots_needed):
			(sets[j]).append(i)
		n_choose_k_sets_fill_slice(sets, from_idx, to_idx, n, k)
		from_idx += slots_needed


def winning_probability(n_turns):
	num_of_blue_disks = int(math.floor(n_turns / 2) + 1)

	prob = 0.

	for k in range(num_of_blue_disks, n_turns + 1):
		draw_possibilities = n_choose_k_sets(n_turns, k)
		
		for draw in draw_possibilities:
			draw_winning_prob = 1.
			for t in range(1, n_turns + 1):
				if t in draw:
					draw_winning_prob *= 1./(t+1.)
				else:
					draw_winning_prob *= t/(t+1.)

			prob += draw_winning_prob

	return prob

winning_prob = winning_probability(n_turns=15)

prize = 1.
while 1. - winning_prob*prize > 0.:
	prize += 1

print(prize - 1)