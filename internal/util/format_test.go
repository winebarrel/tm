package util_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/tm/internal/util"
)

func TestColonNotation(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		str      string
		expected string
	}{
		{str: "3s", expected: "0:00:03"},
		{str: "4m13s", expected: "0:04:13"},
		{str: "5h41m13s", expected: "5:41:13"},
		{str: "123h34m45s", expected: "123:34:45"},
		{str: "-123h34m45s", expected: "-123:34:45"},
		{str: "123h34m45s3ms", expected: "123:34:45.003"},
		{str: "123h34m45s300ms", expected: "123:34:45.3"},
		{str: "123h34m45s3us", expected: "123:34:45.000003"},
		{str: "123h34m45s3ns", expected: "123:34:45.000000003"},
	}

	for _, t := range tt {
		d, _ := time.ParseDuration(t.str)
		cn := util.ColonNotation(d)
		assert.Equal(t.expected, cn)
	}
}
