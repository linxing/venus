package job

import (
	"context"
	"errors"
	"net/http"
	"venus/model"

	"github.com/360EntSecGroup-Skylar/excelize"

	jobintf "venus/job/intf"
)

func (*servlet) ImportFromExcel(ctx context.Context, req *jobintf.ImportFromExcelReq) error {

	if ctx == nil {
		ctx = context.Background()
	}

	resp, err := http.Get(req.URL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	r, err := excelize.OpenReader(resp.Body)
	if err != nil {
		return err
	}

	sheets := make([][][]string, 0)
	for _, name := range r.GetSheetMap() {

		one := r.GetRows(name)
		if len(one) == 0 {
			continue
		}

		sheets = append(sheets, one)
	}

	if len(sheets) == 0 {
		err = errors.New("file not data")
		return err
	}

	for _, sheet := range sheets {

		totalRow := len(sheet)

		for row := 1; row < totalRow; row++ {

			err := model.UserStatic.Insert(ctx, &model.User{
				UserName:    sheet[row][0],
				PhoneNumber: sheet[row][1],
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
