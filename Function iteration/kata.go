package Function_iteration

func CreateIterator(fn func(int) int,n int) func(int) int {
	return func(v int) int {
		for i := 0; i < n; i++ {
			v = fn(v)
		}
		return v
	}
}