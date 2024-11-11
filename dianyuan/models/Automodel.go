package models

import "dianyuan/dao"

func Automodel() {
	dao.DB.AutoMigrate(&Eis_001{}, &Vct_001{})
}
