// 自动生成模板Khinfo
package first

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 客户信息 结构体  Khinfo
type Khinfo struct {
    global.GVA_MODEL
    Ddtype  string `json:"ddtype" form:"ddtype" gorm:"column:ddtype;comment:;"`  //订单类型 
    Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`  //姓名 
    Sex  string `json:"sex" form:"sex" gorm:"column:sex;comment:;"`  //性别 
    Address  string `json:"address" form:"address" gorm:"column:address;comment:;"`  //地址 
    Birthday  string `json:"birthday" form:"birthday" gorm:"column:birthday;comment:;"`  //出生日期 
    Getuser  string `json:"getuser" form:"getuser" gorm:"column:getuser;comment:;"`  //收件人姓名 
    Gx  string `json:"gx" form:"gx" gorm:"column:gx;comment:;"`  //关系 
    Boxtype  string `json:"boxtype" form:"boxtype" gorm:"column:boxtype;comment:;"`  //安置类型 
    Number  string `json:"number" form:"number" gorm:"column:number;comment:;"`  //安置数量 
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 客户信息 Khinfo自定义表名 khinfo
func (Khinfo) TableName() string {
    return "khinfo"
}

