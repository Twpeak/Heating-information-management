package system

type ServiceGroup struct {
	CasbinService
	JwtService
	UserService
	HospitalService
	DistrictService
}
