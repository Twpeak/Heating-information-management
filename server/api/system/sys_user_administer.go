package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/dto"
	"github.com/gin-gonic/gin"
)

func UserInformation(c *gin.Context) {
	page := c.Query("page")
	offset := c.Query("offset")
	all, err := userService.QueryUserAll(page, offset)
	if err != nil {
		global.G_LOG.Error("接口:UserInformation,获取用户所有数据失败,error:" + err.Error())
		response.FailWithMessage("查询数据失败", c)
		return
	}
	response.OkWithData(all, c)
}

func UserAdd(c *gin.Context) {
	var user dto.UserCreateDto
	_ = c.ShouldBindJSON(&user)
	u, err := userService.IsUsername(user.Username)
	if err != nil {
		global.G_LOG.Error("接口:UserAdd,添加用户失败,error:" + err.Error())
		response.FailWithMessage("添加用户失败", c)
		return
	}
	if u.Username != "" {
		response.FailWithMessage("用户名已存在", c)
		return
	}
	err = userService.CreateUser(user)
	if err != nil {
		global.G_LOG.Error("接口:UserAdd,添加用户失败,error:" + err.Error())
		response.FailWithMessage("添加用户失败", c)
		return
	}
	response.Ok(c)
}
func UserUpdateDisplay(c *gin.Context) {
	id := c.Query("id")
	byId, err := userService.QueryUserById(id)
	if err != nil {
		global.G_LOG.Error("接口:UserUpdateDisplay,查询用户失败,error:" + err.Error())
		response.FailWithMessage("查询用户失败", c)
		return
	}
	response.OkWithData(byId, c)
}

func UserUpdate(c *gin.Context) {
	var user dto.UserUpdateDto
	_ = c.ShouldBindJSON(&user)
	err := userService.UpdateUser(user)
	if err != nil {
		global.G_LOG.Error("接口:UserUpdate,修改用户失败,error:" + err.Error())
		response.FailWithMessage("修改用户失败", c)
		return
	}
	response.Ok(c)
}

func UserDelete(c *gin.Context) {
	var id dto.IdDto
	_ = c.ShouldBindJSON(&id)

	h, err := hospitalService.QueryBoosId(id.Id)
	if err != nil {
		global.G_LOG.Error("接口:UserDelete,删除用户失败,error:" + err.Error())
		response.FailWithMessage("删除用户失败", c)
		return
	}
	if h.HospitalName != "" {
		response.FailWithMessage("请先修改医院负责人", c)
		return
	}
	err = userService.DeleteUser(id.Id)
	if err != nil {
		global.G_LOG.Error("接口:UserDelete,删除用户失败,error:" + err.Error())
		response.FailWithMessage("删除用户失败", c)
		return
	}
	response.Ok(c)
}

func MyUpdateText(c *gin.Context) {
	id := c.Query("id")
	text, err := userService.QueryUserText(id)
	if err != nil {
		global.G_LOG.Error("接口:MyUpdateText,查询用户失败,error:" + err.Error())
		response.FailWithMessage("查询用户失败", c)
		return
	}
	response.OkWithData(dto.UserTextDto{
		Id:           text.ID,
		Name:         text.Name,
		IdentityCard: text.IdentityCard,
		Phone:        text.Phone,
	}, c)

}
func MyUpdate(c *gin.Context) {
	var d dto.UserTextDto
	_ = c.ShouldBindJSON(&d)
	err := userService.UpdateUserText(d)
	if err != nil {
		global.G_LOG.Error("接口:MyUpdate,修改用户失败,error:" + err.Error())
		response.FailWithMessage("修改用户失败", c)
		return
	}
	response.Ok(c)
}

func MyUpdatePwd(c *gin.Context) {
	var u dto.MyPwdDto
	_ = c.ShouldBindJSON(&u)
	if u.Pwd1 != u.Pwd2 {
		response.FailWithMessage("新密码不一致", c)
		return
	}
	pwd, err := userService.QueryUserByIdPwd(u)
	if err != nil {
		global.G_LOG.Error("接口:MyUpdatePwd,修改密码失败,error:" + err.Error())
		response.FailWithMessage("修改密码失败", c)
		return
	}
	if pwd.Name == "" {
		response.FailWithMessage("原密码错误", c)
		return
	}

	err = userService.UpdatePwd(u)
	if err != nil {
		global.G_LOG.Error("接口:MyUpdatePwd,修改密码失败,error:" + err.Error())
		response.FailWithMessage("修改密码失败", c)
		return
	}
	response.Ok(c)

}
