package Functional_Addition

func Add(a int) func(int)int {
	return func(b int) int {
		return a + b
	}
}