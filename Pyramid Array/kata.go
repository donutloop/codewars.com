package Pyramid_Array

func Pyramid(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}

	r := make([][]int, 0)
	tpl := []int{}
	for i := 0; i < n; i++ {

		j := len(tpl) - i + 1
		for ;0 < j; j-- {
			tpl = append(tpl, 1)
		}

		p := make([]int, len(tpl))
		copy(p, tpl)

		r = append(r, p)
	}

	return r
}
