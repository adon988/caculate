package caculate

type Caculator struct {
}

type CaculatorInterface interface {
	Sum(x, y float64) float64
	Sub(x, y float64) float64
	Multiplication(x, y float64) float64
	Division(x, y float64) float64
}

func (Caculator) Sum(x, y float64) float64 {
	n := x + y
	return n
}

func (Caculator) Sub(x, y float64) float64 {
	n := x - y
	return n
}

func (Caculator) Multiplication(x, y float64) float64 {
	n := x * y
	return n
}

func (Caculator) Division(x, y float64) float64 {
	n := x / y
	return n
}

func Hello() string {
	return "hi"
}
