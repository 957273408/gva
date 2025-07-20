
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model/request"
)

var Hero = new(hero)

type hero struct {}
// CreateHero 创建英雄记录
// Author [yourname](https://github.com/yourname)
func (s *hero) CreateHero(ctx context.Context, hero *model.Hero) (err error) {
	err = global.GVA_DB.Create(hero).Error
	return err
}

// DeleteHero 删除英雄记录
// Author [yourname](https://github.com/yourname)
func (s *hero) DeleteHero(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&model.Hero{},"id = ?",ID).Error
	return err
}

// DeleteHeroByIds 批量删除英雄记录
// Author [yourname](https://github.com/yourname)
func (s *hero) DeleteHeroByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.Hero{},"id in ?",IDs).Error
	return err
}

// UpdateHero 更新英雄记录
// Author [yourname](https://github.com/yourname)
func (s *hero) UpdateHero(ctx context.Context, hero model.Hero) (err error) {
	err = global.GVA_DB.Model(&model.Hero{}).Where("id = ?",hero.ID).Updates(&hero).Error
	return err
}

// GetHero 根据ID获取英雄记录
// Author [yourname](https://github.com/yourname)
func (s *hero) GetHero(ctx context.Context, ID string) (hero model.Hero, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&hero).Error
	return
}
// GetHeroInfoList 分页获取英雄记录
// Author [yourname](https://github.com/yourname)
func (s *hero) GetHeroInfoList(ctx context.Context, info request.HeroSearch) (list []model.Hero, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Hero{})
    var heros []model.Hero
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
  
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&heros).Error
	return  heros, total, err
}

func (s *hero)GetHeroPublic(ctx context.Context) {

}
