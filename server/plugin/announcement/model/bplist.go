
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Bplist bp记录 结构体
type Bplist struct {
    global.GVA_MODEL
  ScheduleId  *string `json:"scheduleId" form:"scheduleId" gorm:"column:schedule_id;"`  //赛程ID
  TeamId  *string `json:"teamId" form:"teamId" gorm:"column:team_id;"`  //队伍名
  Round  *int `json:"round" form:"round" gorm:"column:round;"`  //小局
  Sort  *string `json:"sort" form:"sort" gorm:"column:sort;"`  //红蓝
  Ban1  *string `json:"ban1" form:"ban1" gorm:"column:ban1;"`  //Ban1
  Ban2  *string `json:"ban2" form:"ban2" gorm:"column:ban2;"`  //Ban2
  Ban3  *string `json:"ban3" form:"ban3" gorm:"column:ban3;"`  //Ban3
  Ban4  *string `json:"ban4" form:"ban4" gorm:"column:ban4;"`  //Ban4
  Ban5  *string `json:"ban5" form:"ban5" gorm:"column:ban5;"`  //ban5
  Player1  *string `json:"player1" form:"player1" gorm:"column:player1;"`  //选手1
  Hero1  *string `json:"hero1" form:"hero1" gorm:"column:hero1;"`  //英雄1
  HeroSkill1  *string `json:"heroSkill1" form:"heroSkill1" gorm:"column:hero_skill1;"`  //召唤师技能1
  Player2  *string `json:"player2" form:"player2" gorm:"column:player2;"`  //选手2
  Hero2  *string `json:"hero2" form:"hero2" gorm:"column:hero2;"`  //英雄2
  HeroSkill2  *string `json:"heroSkill2" form:"heroSkill2" gorm:"column:hero_skill2;"`  //召唤师技能2
  Player3  *string `json:"player3" form:"player3" gorm:"column:player3;"`  //选手3
  Hero3  *string `json:"hero3" form:"hero3" gorm:"column:hero3;"`  //英雄3
  HeroSkill3  *string `json:"heroSkill3" form:"heroSkill3" gorm:"column:hero_skill3;"`  //召唤师技能3
  Player4  *string `json:"player4" form:"player4" gorm:"column:player4;"`  //选手4
  Hero4  *string `json:"hero4" form:"hero4" gorm:"column:hero4;"`  //英雄4
  HeroSkill4  *string `json:"heroSkill4" form:"heroSkill4" gorm:"column:hero_skill4;"`  //召唤师技能4
  Player5  *string `json:"player5" form:"player5" gorm:"column:player5;"`  //选手5
  Hero5  *string `json:"hero5" form:"hero5" gorm:"column:hero5;"`  //英雄5
  HeroSkill5  *string `json:"heroSkill5" form:"heroSkill5" gorm:"column:hero_skill5;"`  //召唤师技能5
}


// TableName bp记录 Bplist自定义表名 bplist
func (Bplist) TableName() string {
    return "bplist"
}







