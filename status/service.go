package status

import (
	"github.com/iok200/blada/pb/iok200_blada_pb"
)

var (
	OK = &ErrorInfo{0, 0, ""}
)

type ErrorInfo struct {
	Service int32
	Code    int32
	Message string
}

func (this *ErrorInfo) ToResultState() *iok200_blada_pb.ResultState {
	return &iok200_blada_pb.ResultState{
		Service: this.Service,
		Code:    this.Code,
		Message: this.Message,
	}
}
