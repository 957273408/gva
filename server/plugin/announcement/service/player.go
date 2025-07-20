
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
)

var Player = new(player)

type player struct {}
// CreatePlayer 创建选手记录
// Author [yourname](https://github.com/yourname)
func (s *player) CreatePlayer(ctx context.Context, player *model.Player) (err error) {
	err = global.GVA_DB.Create(player).Error
	return err
}

// DeletePlayer 删除选手记录
// Author [yourname](https://github.com/yourname)
func (s *player) DeletePlayer(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Player{},"id = ?",ID).Error
	return err
}

// DeletePlayerByIds 批量删除选手记录
// Author [yourname](https://github.com/yourname)
func (s *player) DeletePlayerByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Player{},"id in ?",IDs).Error
	return err
}

// UpdatePlayer 更新选手记录
// Author [yourname](https://github.com/yourname)
func (s *player) UpdatePlayer(ctx context.Context, player model.Player) (err error) {
	err = global.GVA_DB.Model(&model.Player{}).Where("id = ?",player.ID).Updates(&player).Error
	return err
}

// GetPlayer 根据ID获取选手记录
// Author [yourname](https://github.com/yourname)
func (s *player) GetPlayer(ctx context.Context, ID string) (player model.Player, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&player).Error
	return
}
// GetPlayerInfoList 分页获取选手记录
// Author [yourname](https://github.com/yourname)
func (s *player) GetPlayerInfoList(ctx context.Context, info request.PlayerSearch) (list []model.Player, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Player{})
    var players []model.Player
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
    if info.Team != nil && *info.Team != "" {
        db = db.Where("team = ?", *info.Team)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&players).Error
	return  players, total, err
}

func (s *player)GetPlayerPublic(ctx context.Context) {

}
