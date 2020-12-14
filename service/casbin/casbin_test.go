package casbin

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"venus/env/envtest"
)

func Test_CheckPermission(t *testing.T) {

	me := envtest.NewMockEnv(t, envtest.MockDBOption())
	defer me.Close()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	sub := fmt.Sprintf("sub_%d", r.Int31())
	obj := fmt.Sprintf("obj_%d", r.Int31())
	act := fmt.Sprintf("act_%d", r.Int31())

	casbinAuth, err := NewCasbinAuth()
	require.NoError(t, err)
	got, err := casbinAuth.AddCasbinPolicies([][]string{{sub, obj, act}})
	require.NoError(t, err)

	err = casbinAuth.Save()
	require.NoError(t, err)

	assert.True(t, got)

	defer func() {
		_, err := casbinAuth.DeleteCasbinPolicies([][]string{{sub, obj, act}})
		if err != nil {
			require.NoError(t, err)
		}
	}()

	permission, err := casbinAuth.CheckPermission(sub, obj, act)
	require.NoError(t, err)
	assert.True(t, permission)

	permission2, err := casbinAuth.CheckPermission("fake_sub", obj, act)
	require.NoError(t, err)
	assert.False(t, permission2)
}

func Test_Delete(t *testing.T) {

	me := envtest.NewMockEnv(t, envtest.MockDBOption())
	defer me.Close()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	sub := fmt.Sprintf("sub_%d", r.Int31())
	obj := fmt.Sprintf("obj_%d", r.Int31())
	act := fmt.Sprintf("act_%d", r.Int31())

	casbinAuth, err := NewCasbinAuth()
	require.NoError(t, err)
	got, err := casbinAuth.AddCasbinPolicies([][]string{{sub, obj, act}})
	require.NoError(t, err)

	assert.True(t, got)

	_, err = casbinAuth.DeleteCasbinPolicies([][]string{{sub, obj, act}})
	require.NoError(t, err)

	permission, err := casbinAuth.CheckPermission(sub, obj, act)
	require.NoError(t, err)
	assert.False(t, permission)

	permission2, err := casbinAuth.CheckPermission("fake_sub", obj, act)
	require.NoError(t, err)
	assert.False(t, permission2)

}
