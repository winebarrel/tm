package tm

import (
	"strings"
	"time"

	"github.com/winebarrel/tm/internal/util"
)

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
