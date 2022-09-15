package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroups struct {
	System 	system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupsApp = new(RouterGroups)

