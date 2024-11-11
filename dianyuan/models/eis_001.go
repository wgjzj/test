package models

import "dianyuan/dao"

// 定义数据库实体
type Eis_001 struct {
	Eis_id     uint   `json:"eis_id"`
	Time_id    int    `json:"time_id"`
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
	Real_imp   string `json:"real_imp"`
	Imag_imp   string `json:"imag_imp"`
	Impedance  string `json:"impedance"`
	Phase      string `json:"phase"`
	Frequency  string `json:"frequency"`
}

func CreateAEis(eis_001 *Eis_001) (err error) {
	err = dao.DB.Create(&eis_001).Error
	if err != nil {
		return err
	}
	return
}

func GetAEislist() (eislist []*Eis_001, err error) {
	err = dao.DB.Find(&eislist).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetAEis(eis_id string) (eis_001 *Eis_001, err error) {
	eis_001 = new(Eis_001)
	err = dao.DB.Where("eis_id=?", eis_id).First(&eis_001).Error
	if err != nil {
		return nil, err
	}
	return
}

func UpdateAEis(eis *Eis_001) (err error) {
	err = dao.DB.Save(&eis).Error
	return

}

func DeleteAEis(eis_id string) (err error) {
	err = dao.DB.Where("eis_id=?", eis_id).Delete(Eis_001{}).Error
	return

}
