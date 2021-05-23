package Simple_Fun__3__Late_Ride

func LateRide(n int) int {
	var hours int
	for n >= 60 {
		hours = hours + 100
		n -= 60
	}

	hours += n

	var sum int
	for hours > 0 {
		k := hours % 10
		hours = hours / 10
		sum += k
	}

	return sum
}