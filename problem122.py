# Problem 122

def merge_and_append_and_sort(a, b, to_append):
	c = list(set(a + b + [to_append]))
	c.sort()
	return c


m_k = [[[]], [[1]]]

the_sum = 0

for k in range(2, 201):
	print("Computing for k: " + str(k))
	best_n = -1
	best_cuts_candidates = []
	for l in range(1, k):
		m_l_candidates = m_k[l]

		for m_l in m_l_candidates:
			rest = k-l

			temp_candidates = m_k[rest]

			for temp in temp_candidates:
				current_n = len(m_l) - 1
				current_cuts = m_l
				
				while (len(temp) > 0) and (temp[0] in current_cuts):
					temp = temp[1:]

				current_n += len(temp) + 1

			
				current_cuts = merge_and_append_and_sort(current_cuts, temp, k)
				if (current_n < best_n) or (best_n == -1):
					best_n = current_n
					best_cuts_candidates = [current_cuts]
				elif current_n == best_n:
					if not current_cuts in best_cuts_candidates:
						best_cuts_candidates.append(current_cuts)

	m_k.append(best_cuts_candidates)
	the_sum += len(best_cuts_candidates[0]) - 1

print(the_sum)

