package caculate

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var c CaculatorInterface = Caculator{}

func TestCaculateAdd(t *testing.T) {
	sumResult := c.sum(2.5, 3.5)
	assert.Equal(t, 6.0, sumResult, "Sum result incorrect")
}

func TestCaculateSub(t *testing.T) {
	subResult := c.sub(2.1, 1.9)
	subResult = math.Round(subResult*100) / 100
	assert.Equal(t, 0.2, subResult, "Sub result incorrect")
}

func TestCaculateMult(t *testing.T) {
	multResult := c.multiplication(3, 2)
	assert.Equal(t, 6.0, multResult, "Mult result ")
}

func TestCaculateDiv(t *testing.T) {
	divResult := c.division(6, 2)
	assert.Equal(t, 3.0, divResult, "Div result incorrect")
}
