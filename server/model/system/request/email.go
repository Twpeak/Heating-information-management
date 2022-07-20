package request

//存储发送email所需的数据
type EmailReq struct {
	//所需要的信息：boss名，boss邮箱，医院名，该医院所有医生名，各个医生本周登记数，本周登记总数
	BossName 		string
	BossEmail		string
	HospitalName 	string
	DoctorNames 	[] string
	DoctorCount		int
	SumCount		int
}
