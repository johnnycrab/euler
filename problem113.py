
num_digits = 100

num_of_incs_starting_with = [1] * 10 # we keep the zero index for simpler indexing
num_of_decs_starting_with = [1] * 10
num_of_incs_starting_with[0] = 0
num_of_decs_starting_with[0] = 0


total_sum = 18

for d in range(1,num_digits):
	num_of_incs_starting_with_new = [0] * 10
	num_of_decs_starting_with_new = [0] * 10

	for i in range(10):
		if i > 0:
			for j in range(i, 10):
				num_of_incs_starting_with_new[j] += num_of_incs_starting_with[i]
		for j in range(i, -1, -1):
			num_of_decs_starting_with_new[j] += num_of_decs_starting_with[i]

	num_of_incs_starting_with = num_of_incs_starting_with_new
	num_of_decs_starting_with = num_of_decs_starting_with_new

	for i in range(10):
		total_sum += num_of_decs_starting_with[i] + num_of_incs_starting_with[i]

print(total_sum - num_digits*9)
