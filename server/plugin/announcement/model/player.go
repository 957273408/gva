
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Player 选手 结构体
type Player struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"column:name;"`  //选手名
  Team  *string `json:"team" form:"team" gorm:"column:team;"`  //战队
  Img  string `json:"img" form:"img" gorm:"column:img;"`  //定妆照
  BImg  string `json:"bImg" form:"bImg" gorm:"column:b_img;"`  //大屏定妆照
}


// TableName 选手 Player自定义表名 player
func (Player) TableName() string {
    return "player"
}







