# Problem 144

from math import sqrt
import matplotlib.pyplot as plt
from matplotlib.patches import Ellipse

def get_new_direction(old_direction_x, old_direction_y, impact_pt_x, impact_pt_y):
	d_1 = -1. * old_direction_x
	d_2 = -1. * old_direction_y
	m = -4. * (impact_pt_x / impact_pt_y)

	new_d_1 = 1./(m*m+1) * ((m*m - 1.)*d_1 + (-2.*m)*d_2)
	new_d_2 = 1./(m*m+1) * ((-2.*m)*d_1 + (1. - m*m)*d_2)

	norm = sqrt(new_d_1*new_d_1 + new_d_2*new_d_2)

	return (new_d_1 / norm, new_d_2 / norm)

def get_new_impact_point(old_impact_x, old_impact_y, direction_x, direction_y):
	a = 4. * direction_x * direction_x + direction_y * direction_y
	b = 8. * old_impact_x * direction_x + 2. * old_impact_y * direction_y
	c = 4. * old_impact_x * old_impact_x + old_impact_y * old_impact_y - 100.

	under_root = b*b - 4*a*c
	print(under_root)
	l = (-1. * b + sqrt(b*b - 4*a*c))/(2. * a)

	return (old_impact_x + l * direction_x, old_impact_y + l * direction_y)


# draw the ellipse
ellipse = Ellipse((0,0), width=10,height= 20)

fig, ax = plt.subplots(subplot_kw={'aspect': 'equal'})
ax.add_artist(ellipse)
ellipse.set_alpha(0.3)

plt.axis([-10., 10., -15., 15.])

impact_x = 0.0
impact_y = 10.1
direction_x = 1.4
direction_y = -19.7

i = 0

while True:
	prev_x = impact_x
	prev_y = impact_y
	impact_x, impact_y = get_new_impact_point(impact_x, impact_y, direction_x, direction_y)
	direction_x, direction_y = get_new_direction(direction_x, direction_y, impact_x, impact_y)

	print("New impact point: (" + str(impact_x) + ", " + str(impact_y) + ")")
	print(4 * impact_x * impact_x + impact_y * impact_y)
	print("New direction: (" + str(direction_x) + ", " + str(direction_y) + ")")
	print("------------------")

	plt.plot([prev_x, impact_x], [prev_y, impact_y], 'ro-')
	#plt.savefig("./temp/fig_" + str(i) + ".png")

	if abs(impact_x) <= 0.01 and impact_y > 0:
		break

	i += 1

print(i)

plt.show()