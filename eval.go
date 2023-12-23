package tm

import (
	"strings"
	"time"
)

func (expr *Expr) Eval() time.Duration {
	sum := expr.Value.Duration()

	for _, opVal := range expr.OpValues {
		switch opVal.Op {
		case "+":
			sum += opVal.Value.Duration()
		case "-":
			sum -= opVal.Value.Duration()
		}
	}

	return sum
}

func Eval(str string) (time.Duration, error) {
	str = strings.TrimSpace(str)
	expr, err := parser.ParseString("", str)

	if err != nil {
		return 0, err
	}

	return expr.Eval(), nil
}
