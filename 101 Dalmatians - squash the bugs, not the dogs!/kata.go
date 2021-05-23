package _01_Dalmatians___squash_the_bugs__not_the_dogs_

func HowManyDalmatians(number int) string {

	dogs := []string{"Hardly any", "More than a handful!", "Woah that's a lot of dogs!", "101 DALMATIONS!!!"}

	if number <= 10 {
		return  dogs[0]
	} else if number <= 50 {
		return dogs[1]
	} else if number <= 100 {
		return dogs[2]
	}
	return dogs[3]
}