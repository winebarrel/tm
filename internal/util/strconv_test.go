package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/tm/internal/util"
)

func TestMustAtoiOk(t *testing.T) {
	assert := assert.New(t)
	n := util.MustAtoi("123")
	assert.Equal(123, n)
}

func TestMustAtoiErr(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	defer func() {
		e := recover()
		err, ok := e.(error)
		require.True(ok)
		assert.ErrorContains(err, `strconv.Atoi: parsing "xxx": invalid syntax`)
	}()

	util.MustAtoi("xxx")
}
