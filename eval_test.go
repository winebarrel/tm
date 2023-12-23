package tm_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/tm"
)

func TestEval(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	tt := []struct {
		str      string
		expected string
	}{
		{str: "1h + 2m + 3s", expected: "1h2m3s"},
		{str: "1h+2m+3s", expected: "1h2m3s"},
		{str: "\t1h\t+\t2m\t+\t3s\t", expected: "1h2m3s"},
		{str: "1:00 + 2:34 + 00:01:56", expected: "3h35m56s"},
		{str: "3", expected: "3s"},
		{str: "1:23 + 3m - 1h + 0:0:5", expected: "26m5s"},
		{str: "5m - 2:23", expected: "-2h18m0s"},
	}

	for _, t := range tt {
		d, err := tm.Eval(t.str)
		require.NoError(err)
		e, err := time.ParseDuration(t.expected)
		require.NoError(err)
		assert.Equal(tm.Result(e), d)
	}
}
