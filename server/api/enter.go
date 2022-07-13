package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
}
var ApiGroupApp = new(ApiGroup)