package Fizz_Buzz_Cuckoo_Clock

func FizzBuzzCuckooClock(time string) string {
	if len(time) != 5 {
		panic("time is bad")
	}

	minutes := int(((time[3]- 48) * 10) + (time[4]- 48))
	hours := int(((time[0] - 48) * 10) +  (time[1] - 48))

	if minutes % 3 == 0 && minutes % 5 == 0 && minutes != 0 && minutes != 30 {
		return "Fizz Buzz"
	} else if minutes % 3 == 0 && minutes != 0 && minutes != 30    {
		return "Fizz"
	} else if minutes % 5 == 0 && minutes != 0 && minutes != 30 {
		return "Buzz"
	} else if  minutes == 0 {
		if 12 == hours || 0 == hours {
			var text string
			for i := 0; i < 12; i++ {
				text += "Cuckoo "
			}
			return text[:len(text)-1]
		}
		if hours >= 12 {
			hours = hours - 12
		}
		var text string
		for i := 0; i < hours; i++ {
			text += "Cuckoo "
		}
		return text[:len(text)-1]
	} else if minutes == 30 {
		return "Cuckoo"
	}

	return "tick"
}