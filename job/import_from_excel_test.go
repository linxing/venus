package job

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"venus/env/envtest"
	jobintf "venus/job/intf"
	"venus/model"
)

func Test_ImportFromExcel(t *testing.T) {

	now := time.Now()
	r := rand.New(rand.NewSource(now.UnixNano()))

	ctx := context.Background()
	servlet := NewServlet()

	t.Run("should succeed", func(t *testing.T) {

		me := envtest.NewMockEnv(t, envtest.MockDBOption(model.User{}), envtest.MockRedisOption())
		defer me.Close()

		expectedUserName := fmt.Sprintf("user_name_%d", r.Int())
		expectedPhoneNumber := fmt.Sprintf("phone_number_%d", r.Int())
		url := ""

		{
			mockHttpService := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()

				f := excelize.NewFile()
				f.SetCellValue("Sheet1", "A1", "username")
				f.SetCellValue("Sheet1", "A2", expectedUserName)
				f.SetCellValue("Sheet1", "B1", "phone_number")
				f.SetCellValue("Sheet1", "B2", expectedPhoneNumber)

				_, err := f.WriteTo(w)
				require.NoError(t, err)
			}))

			defer mockHttpService.Close()

			url = mockHttpService.URL
		}

		err := servlet.ImportFromExcel(ctx, &jobintf.ImportFromExcelReq{
			URL: url,
		})
		require.NoError(t, err)

		gotItem, err := model.UserStatic.GetByName(ctx, expectedUserName)
		require.NoError(t, err)
		assert.Equal(t, expectedPhoneNumber, gotItem.PhoneNumber)
	})
}
