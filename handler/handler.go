package handler

import (
	"github.com/fernandobloedorn/Go-Lang-Dev-Ops/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializaHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
