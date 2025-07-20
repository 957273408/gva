
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
)

var Bplist = new(bplist)

type bplist struct {}
// CreateBplist 创建bp记录记录
// Author [yourname](https://github.com/yourname)
func (s *bplist) CreateBplist(ctx context.Context, bplist *model.Bplist) (err error) {
	err = global.GVA_DB.Create(bplist).Error
	return err
}

// DeleteBplist 删除bp记录记录
// Author [yourname](https://github.com/yourname)
func (s *bplist) DeleteBplist(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Bplist{},"id = ?",ID).Error
	return err
}

// DeleteBplistByIds 批量删除bp记录记录
// Author [yourname](https://github.com/yourname)
func (s *bplist) DeleteBplistByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Bplist{},"id in ?",IDs).Error
	return err
}

// UpdateBplist 更新bp记录记录
// Author [yourname](https://github.com/yourname)
func (s *bplist) UpdateBplist(ctx context.Context, bplist model.Bplist) (err error) {
	err = global.GVA_DB.Model(&model.Bplist{}).Where("id = ?",bplist.ID).Updates(&bplist).Error
	return err
}

// GetBplist 根据ID获取bp记录记录
// Author [yourname](https://github.com/yourname)
func (s *bplist) GetBplist(ctx context.Context, ID string) (bplist model.Bplist, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&bplist).Error
	return
}
// GetBplistInfoList 分页获取bp记录记录
// Author [yourname](https://github.com/yourname)
func (s *bplist) GetBplistInfoList(ctx context.Context, info request.BplistSearch) (list []model.Bplist, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Bplist{})
    var bplists []model.Bplist
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.ScheduleId != nil && *info.ScheduleId != "" {
        db = db.Where("schedule_id = ?", *info.ScheduleId)
    }
    if info.TeamId != nil && *info.TeamId != "" {
        db = db.Where("team_id = ?", *info.TeamId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&bplists).Error
	return  bplists, total, err
}

func (s *bplist)GetBplistPublic(ctx context.Context) {

}
