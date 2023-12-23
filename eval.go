package tm

import (
	"strings"
	"time"

	"github.com/winebarrel/tm/internal/util"
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

type Result time.Duration

func (r Result) String() string {
	return util.ColonNotation(time.Duration(r))
}

func Eval(str string) (Result, error) {
	str = strings.TrimSpace(str)
	expr, err := parser.ParseString("", str)

	if err != nil {
		return 0, err
	}

	return Result(expr.Eval()), nil
}
