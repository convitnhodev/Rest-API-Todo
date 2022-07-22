package component

import (
	 "gorm.io/gorm"
)

type AppContext interface {
	 GetMainDbConnection() *gorm.DB
}

type appCtx struct {
	 db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	 return &appCtx{db}
}

func (ctx *appCtx) GetMainDbConnection() *gorm.DB {
	 return ctx.db
}
