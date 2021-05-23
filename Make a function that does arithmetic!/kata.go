package Make_a_function_that_does_arithmetic_

type operator string

const (
	ADD operator = "add"
	SUBSTRACT operator = "subtract"
	MULTIPLY operator = "multiply"
	DIVIDE operator = "divide"
)

func Arithmetic(a int, b int, op string) (result int) {

	switch operator(op) {
	case ADD:
		result = a + b
	case SUBSTRACT:
		result = a - b
	case MULTIPLY:
		result = a * b
	case DIVIDE:
		result = a / b
	default:
		panic("operation is not supported")
	}

	return
}