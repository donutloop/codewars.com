package Make_the_Deadfish_swim

func Parse(data string) []int{
	var v int
	res := make([]int, 0)
	for _, c := range data {
		if c == 'i' {
			v++
		} else if c == 'd' {
			v--
		} else if c == 's' {
			v = v * v
		} else if c == 'o' {
			res = append(res, v)
		}
	}
	return res
}