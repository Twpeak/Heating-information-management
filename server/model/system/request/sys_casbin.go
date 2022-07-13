package request

//封装 所有访问鉴权的请求

// CasbinInfo Casbin信息结构
type CasbinInfo struct {
	Path	string	`json:"path"`	//路由
	Method 	string	`json:"method"`	//方法
}

// CasbinInReceive 输入参数的Casbin结构
type CasbinInReceive struct {
	RoleId	uint			`json:"roleId"`//权限id（角色id）
	CasbinInfos	[]CasbinInfo	`json:"casbinInfos" `
}
