package first

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ KhinfoApi }

var khinfoService = service.ServiceGroupApp.FirstServiceGroup.KhinfoService
