package first

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/first"
    firstReq "github.com/flipped-aurora/gin-vue-admin/server/model/first/request"
    "gorm.io/gorm"
)

type KhinfoService struct {}

// CreateKhinfo 创建客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (khinfoService *KhinfoService) CreateKhinfo(khinfo *first.Khinfo) (err error) {
	err = global.GVA_DB.Create(khinfo).Error
	return err
}

// DeleteKhinfo 删除客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (khinfoService *KhinfoService)DeleteKhinfo(ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&first.Khinfo{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&first.Khinfo{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteKhinfoByIds 批量删除客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (khinfoService *KhinfoService)DeleteKhinfoByIds(IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&first.Khinfo{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&first.Khinfo{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateKhinfo 更新客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (khinfoService *KhinfoService)UpdateKhinfo(khinfo first.Khinfo) (err error) {
	err = global.GVA_DB.Model(&first.Khinfo{}).Where("id = ?",khinfo.ID).Updates(&khinfo).Error
	return err
}

// GetKhinfo 根据ID获取客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (khinfoService *KhinfoService)GetKhinfo(ID string) (khinfo first.Khinfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&khinfo).Error
	return
}

// GetKhinfoInfoList 分页获取客户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (khinfoService *KhinfoService)GetKhinfoInfoList(info firstReq.KhinfoSearch) (list []first.Khinfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&first.Khinfo{})
    var khinfos []first.Khinfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Ddtype != "" {
        db = db.Where("ddtype = ?",info.Ddtype)
    }
    if info.Name != "" {
        db = db.Where("name = ?",info.Name)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&khinfos).Error
	return  khinfos, total, err
}