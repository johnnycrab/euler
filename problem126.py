# Problem 126: Cuboid layers
# NOT YET SOLVED

import numpy as np


# Idea: We represent each cuboid as a 3-dimensional tensor, where each element "-1" represents one unit cube.
# To cover it with one layer, we expand the dimension by 1 in each direction and lay elements "1" around each "-1"


def make_cuboid(dim_x, dim_y, dim_z):
	cuboid = np.zeros((dim_x, dim_y, dim_z))
	cuboid.fill(-1)
	
	return cuboid

def expand_cuboid(cuboid):
	expanded = np.zeros((cuboid.shape[0] + 2, cuboid.shape[1] + 2, cuboid.shape[2] + 2))
	expanded[1:cuboid.shape[0] + 1, 1:cuboid.shape[1] + 1, 1:cuboid.shape[2] + 1] = cuboid
	return expanded

def cover_cube(cuboid, at_pos_x, at_pos_y, at_pos_z):
	if cuboid[at_pos_x + 1, at_pos_y, at_pos_z] == 0:
		cuboid[at_pos_x + 1, at_pos_y, at_pos_z] = 1	
	if cuboid[at_pos_x - 1, at_pos_y, at_pos_z] == 0:
		cuboid[at_pos_x - 1, at_pos_y, at_pos_z] = 1
	if cuboid[at_pos_x, at_pos_y + 1, at_pos_z] == 0:
		cuboid[at_pos_x, at_pos_y + 1, at_pos_z] = 1
	if cuboid[at_pos_x, at_pos_y - 1, at_pos_z] == 0:
		cuboid[at_pos_x, at_pos_y - 1, at_pos_z] = 1
	if cuboid[at_pos_x, at_pos_y, at_pos_z + 1] == 0:
		cuboid[at_pos_x, at_pos_y, at_pos_z + 1] = 1
	if cuboid[at_pos_x, at_pos_y, at_pos_z - 1] == 0:
		cuboid[at_pos_x, at_pos_y, at_pos_z - 1] = 1


# returns covered cuboid, all cubes set to -1, together with the number of cubes needed to layer the cuboid
def cover_cuboid(cuboid):
	current_num_of_cubes = -1 * np.sum(cuboid)
	expanded_cuboid = expand_cuboid(cuboid)

	for i in range(expanded_cuboid.shape[0]):
		for j in range(expanded_cuboid.shape[1]):
			for k in range(expanded_cuboid.shape[2]):
				if expanded_cuboid[i, j, k] == -1:
					cover_cube(expanded_cuboid, i, j, k)
	expanded_cuboid = np.absolute(expanded_cuboid)

	return -1 * expanded_cuboid, np.sum(expanded_cuboid) - current_num_of_cubes

cuboid = make_cuboid(3, 2, 1)



def C(n):
	cuboid_count = 0

	for a in range(1, n+1):
		for b in range(a, 0, -1):
			for c in range(b, 0, -1):
				if a*b*c > n:
					continue
				current_cubes = 0
				cuboid = make_cuboid(a, b, c)
				while current_cubes < n:
					cuboid, current_cubes = cover_cuboid(cuboid)

				if current_cubes == n:
					cuboid_count += 1

	return cuboid_count

n = 119

while True:
	c_n = C(n)
	print("n: " + str(n) + ", C(n): " + str(c_n))
	if c_n == 1000:
		break
	n += 1

