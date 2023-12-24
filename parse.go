package tm

import (
	"strconv"
	"strings"
	"time"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	tmLexer = lexer.MustSimple([]lexer.SimpleRule{
		{Name: `Tm`, Pattern: `\d*:\d*(:\d*)?`},
		{Name: `Dur1`, Pattern: `\d+h(\d+m)?(\d+s)?`},
		{Name: `Dur2`, Pattern: `\d+m(\d+s)?`},
		{Name: `Dur3`, Pattern: `\d+s?`},
		{Name: `Symbol`, Pattern: `[-+]`},
		{Name: `SP`, Pattern: `\s+`},
	})

	parser = participle.MustBuild[Expr](
		participle.Lexer(tmLexer),
	)
)

type Tm time.Duration

func (v *Tm) Capture(values []string) error {
	t := strings.SplitN(values[0], ":", 3)
	var hh, mm, ss int

	if t[0] != "" {
		hh, _ = strconv.Atoi(t[0])
	}

	if t[1] != "" {
		mm, _ = strconv.Atoi(t[1])
	}

	if len(t) == 3 {
		ss, _ = strconv.Atoi(t[2])
	}

	*v = Tm(
		time.Duration(hh)*time.Hour +
			time.Duration(mm)*time.Minute +
			time.Duration(ss)*time.Second,
	)

	return nil
}

type Dur time.Duration

func (v *Dur) Capture(values []string) error {
	t := values[0]
	u := t[len(t)-1]
	var d time.Duration

	if u == 'h' || u == 'm' || u == 's' {
		d, _ = time.ParseDuration(values[0])
	} else {
		n, _ := strconv.Atoi(t)
		d = time.Duration(n) * time.Second
	}

	*v = Dur(d)
	return nil
}

type Value struct {
	Tm  *Tm  `@Tm`
	Dur *Dur `| ( @Dur1 | @Dur2 | @Dur3 )`
}

func (v *Value) Duration() time.Duration {
	if v.Tm != nil {
		return time.Duration(*v.Tm)
	} else {
		return time.Duration(*v.Dur)
	}
}

type OpValue struct {
	Op    string `( @"+" | @"-" )`
	Value Value  `SP* @@`
}

type Expr struct {
	Value    Value     `@@`
	OpValues []OpValue `( SP* @@ )*`
}
