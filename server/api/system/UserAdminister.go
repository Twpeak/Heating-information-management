package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/dto"
	"github.com/gin-gonic/gin"
)

func UserInformation(c *gin.Context) {
	all, err := userService.QueryUserAll()
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

	h, err := hospital.QueryBoosId(id.Id)
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
