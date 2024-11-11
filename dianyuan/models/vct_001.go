package models

import (
	"dianyuan/dao"
	"time"
)

type Vct_001 struct {
	Time_id     uint      `json:"time_id"`
	Start_time  time.Time `json:"start_time"`
	End_time    time.Time `json:"end_time"`
	Voltage     string    `json:"voltage"`
	Current     string    `json:"current"`
	Temperature string    `json:"temperature"`
}

func CreateAVct(vct_001 *Vct_001) (err error) {
	err = dao.DB.Create(&vct_001).Error
	if err != nil {
		return err
	}
	return
}

func GetAVctlist() (vctlist []*Vct_001, err error) {
	err = dao.DB.Find(&vctlist).Error
	if err != nil {
		return nil, err
	}
	return
}

func GetAVct(time_id string) (vct_001 *Vct_001, err error) {
	vct_001 = new(Vct_001)
	err = dao.DB.Where("time_id=?", time_id).First(&vct_001).Error
	if err != nil {
		return nil, err
	}
	return
}

func UpdateAVct(vct_001 *Vct_001) (err error) {
	err = dao.DB.Save(&vct_001).Error
	return

}

func DeleteAVct(time_id string) (err error) {
	err = dao.DB.Where("time_id=?", time_id).Delete(Vct_001{}).Error
	return

}
