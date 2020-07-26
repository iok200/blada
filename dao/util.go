package dao

import "github.com/iok200/blada/pb"

func GetPageInfo(request *pb.PageRequest) (page int64, rows int64) {
	if request.GetPage() <= 0 || request.GetRows() <= 0 {
		page = 0
	} else {
		page = (request.GetPage() - 1) * request.GetRows()
	}
	if request.GetRows() <= 0 {
		rows = 10
	} else {
		rows = request.GetRows()
	}
	return
}
