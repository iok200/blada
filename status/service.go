package status

import "github.com/iok200/blada/pb"

var (
	OK = &ErrorInfo{0, 0, ""}
)

type ErrorInfo struct {
	Service int32
	Code    int32
	Message string
}

func (this *ErrorInfo) ToResultState() *pb.ResultState {
	return &pb.ResultState{
		Service: this.Service,
		Code:    this.Code,
		Message: this.Message,
	}
}
