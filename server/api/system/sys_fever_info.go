package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type FeverApi struct {
}


func (f *FeverApi)Test(c *gin.Context)  {
	response.Ok(c)
}