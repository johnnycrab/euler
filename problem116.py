# Problem 116

def get_ways(tile_len):
	global n_ways

	for k in range(2, 5):
		n_k = 0
		t = tile_len
		while t - k >= 0:
			n_k += 1 + n_ways[k-2][t-k]
			t -= 1
		n_ways[k-2].append(n_k)
	return n_ways[0][-1] + n_ways[1][-1] + n_ways[2][-1]


n_ways = [
	[0, 0, 1, 2, 4],
	[0, 0, 0, 1, 2],
	[0, 0, 0, 0, 1]
]

for tile_len in range(5, 51):
	a = get_ways(tile_len)
	if tile_len == 50:
		print(a)

