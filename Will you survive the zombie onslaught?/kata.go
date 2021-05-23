package Will_you_survive_the_zombie_onslaught_

import "fmt"

func Zombie_shootout(zombies, initial_range, ammo int) string {
	if zombies == 0 {
		return fmt.Sprintf("You shot all 0 zombies.")
	}

	var c int
	j := float64(initial_range)
	for i := j; i > .0; i = i - 0.5 {
		diff := ammo - 1
		if diff >= 0 {
			ammo = ammo - 1
			zombies = zombies - 1
			c++
		} else {
			break
		}
		if zombies == 0 {
			return fmt.Sprintf("You shot all %d zombies.", c)
		}
		j = j - 0.5
	}
	var t string
	if  ammo == 0 && zombies >= 1 && j == 0.0 {
		t = "overwhelmed."
	} else if ammo == 0 {
		t = "ran out of ammo."
	} else {
		t = "overwhelmed."
	}
	return fmt.Sprintf("You shot %d zombies before being eaten: %s", c, t)
}