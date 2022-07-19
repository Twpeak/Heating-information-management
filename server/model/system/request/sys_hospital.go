package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

//医院展示信息中，有三表关联，所以参数最多需要三个id
type HospitalReq struct {
	request.PageInfo
	HospitalId	uint	`json:"hospital_id"`
	UserId   	uint	`json:"user_id"`
	DistrictId 	uint	`json:"district_id"`
}

//关键字
type KeyReq struct {
	request.PageInfo
	Key string	`json:"key"`
}
