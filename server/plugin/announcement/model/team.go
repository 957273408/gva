
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Team 战队 结构体
type Team struct {
    global.GVA_MODEL
  Team  *string `json:"team" form:"team" gorm:"primarykey;column:team;" binding:"required"`  //战队名
  Logo  string `json:"logo" form:"logo" gorm:"column:logo;"`  //战队LOGO
}


// TableName 战队 Team自定义表名 team
func (Team) TableName() string {
    return "team"
}







