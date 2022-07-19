package example

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	EmailApi
}
var (
	emailService = service.ServiceGroupApp.ExampleServiceGroup.EmailService
)