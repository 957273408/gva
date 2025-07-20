package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement/model"
	"gorm.io/gen"
	"path/filepath"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "announcement", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(new(model.Info), new(model.Hero), new(model.Team), new(model.Player), //go:generate go mod tidy
		//go:generate go mod download
		//go:generate go run gen.go

		new(model.Bplist),
	)
	g.Execute()
}
