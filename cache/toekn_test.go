package cache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"venus/env/envtest"
)

func Test_GetData(t *testing.T) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Run("should success in normal case", func(t *testing.T) {

		me := envtest.NewMockEnv(t, envtest.MockRedisOption())
		defer me.Close()

		k := fmt.Sprintf("key_%d", r.Int())
		v := fmt.Sprintf("value_%d", r.Int())

		success, err := TokenStatic.SetData(k, v, 100)
		require.NoError(t, err)

		assert.True(t, success)

		gotValue, err := TokenStatic.GetData(k)
		require.NoError(t, err)

		assert.Equal(t, v, gotValue)
	})
}
