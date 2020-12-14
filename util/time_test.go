package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetZeroDayTime(t *testing.T) {

	layout := "2006-01-02 15:04:05"
	str := "2020-08-11 10:10:33"

	expectedStartAt := "2020-08-11 00:00:00"
	expectedEndAt := "2020-08-12 00:00:00"

	inputTime, err := time.Parse(layout, str)
	require.NoError(t, err)

	s, e := GetZeroDayTime(inputTime)

	assert.Equal(t, expectedStartAt, s.Format(layout))
	assert.Equal(t, expectedEndAt, e.Format(layout))
}
