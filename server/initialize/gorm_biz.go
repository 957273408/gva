package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/first"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(first.Khinfo{})
	if err != nil {
		return err
	}
	return nil
}
