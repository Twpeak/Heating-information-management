package request

//医院展示信息中，有三表关联，所以参数最多需要三个id
type HospitalReq struct {
		HospitalId	uint	`json:"hospital_id"`
	UserId   	uint	`json:"user_id"`
	DistrictId 	uint	`json:"district_id"`
}
