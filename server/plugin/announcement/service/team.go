
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
)

var Team = new(team)

type team struct {}
// CreateTeam 创建战队记录
// Author [yourname](https://github.com/yourname)
func (s *team) CreateTeam(ctx context.Context, team *model.Team) (err error) {
	err = global.GVA_DB.Create(team).Error
	return err
}

// DeleteTeam 删除战队记录
// Author [yourname](https://github.com/yourname)
func (s *team) DeleteTeam(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Team{},"id = ?",ID).Error
	return err
}

// DeleteTeamByIds 批量删除战队记录
// Author [yourname](https://github.com/yourname)
func (s *team) DeleteTeamByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Team{},"id in ?",IDs).Error
	return err
}

// UpdateTeam 更新战队记录
// Author [yourname](https://github.com/yourname)
func (s *team) UpdateTeam(ctx context.Context, team model.Team) (err error) {
	err = global.GVA_DB.Model(&model.Team{}).Where("id = ?",team.ID).Updates(&team).Error
	return err
}

// GetTeam 根据ID获取战队记录
// Author [yourname](https://github.com/yourname)
func (s *team) GetTeam(ctx context.Context, ID string) (team model.Team, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&team).Error
	return
}
// GetTeamInfoList 分页获取战队记录
// Author [yourname](https://github.com/yourname)
func (s *team) GetTeamInfoList(ctx context.Context, info request.TeamSearch) (list []model.Team, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Team{})
    var teams []model.Team
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
	err = db.Find(&teams).Error
	return  teams, total, err
}

func (s *team)GetTeamPublic(ctx context.Context) {

}
