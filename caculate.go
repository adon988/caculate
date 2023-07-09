package caculate

type Caculator struct {
}

type CaculatorInterface interface {
	sum(x, y float64) float64
	sub(x, y float64) float64
	multiplication(x, y float64) float64
	division(x, y float64) float64
}

func (Caculator) sum(x, y float64) float64 {
	n := x + y
	return n
}

func (Caculator) sub(x, y float64) float64 {
	n := x - y
	return n
}

func (Caculator) multiplication(x, y float64) float64 {
	n := x * y
	return n
}

func (Caculator) division(x, y float64) float64 {
	n := x / y
	return n
}
