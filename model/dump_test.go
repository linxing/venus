package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_DumpAllToFile(t *testing.T) {
	t.Run("testing xorm dump all to file", func(t *testing.T) {
		db, err := NewDB()
		require.NoError(t, err)

		defer db.Close()
		//db.Engine().DumpAllToFile("dump.sql")
	})
}
