
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Hero 英雄 结构体
type Hero struct {
    global.GVA_MODEL
  HeroName  *string `json:"heroName" form:"heroName" gorm:"index;column:hero_name;" binding:"required"`  //英雄名
  HeroImg  string `json:"heroImg" form:"heroImg" gorm:"column:hero_img;"`  //英雄原画
  HeroHd  string `json:"heroHd" form:"heroHd" gorm:"column:hero_hd;"`  //英雄头像
}


// TableName 英雄 Hero自定义表名 hero
func (Hero) TableName() string {
    return "hero"
}







