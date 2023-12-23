package tm_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/tm"
)

func TestTmCapture(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	tt := []struct {
		str      string
		expected string
	}{
		{str: "01:23:45", expected: "1h23m45s"},
		{str: "01:23", expected: "1h23m"},
		{str: "123:456", expected: "130h36m0s"},
		{str: ":23", expected: "23m"},
		{str: "::45", expected: "45s"},
		{str: "1:", expected: "1h"},
		{str: ":", expected: "0s"},
	}

	for _, t := range tt {
		var v tm.Tm
		err := v.Capture([]string{t.str})
		require.NoError(err)
		d := time.Duration(v)
		e, err := time.ParseDuration(t.expected)
		require.NoError(err)
		assert.Equal(e, d)
	}
}

func TestDurCapture(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	tt := []struct {
		str      string
		expected string
	}{
		{str: "1s", expected: "1s"},
		{str: "23m", expected: "23m0s"},
		{str: "456h", expected: "456h0m0s"},
		{str: "1h2m3s", expected: "1h2m3s"},
		{str: "1h2m", expected: "1h2m"},
		{str: "1h3s", expected: "1h3s"},
		{str: "2m3s", expected: "2m3s"},
	}

	for _, t := range tt {
		var v tm.Dur
		err := v.Capture([]string{t.str})
		require.NoError(err)
		d := time.Duration(v)
		e, err := time.ParseDuration(t.expected)
		require.NoError(err)
		assert.Equal(e, d)
	}
}
