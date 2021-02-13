package models

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"meetingRecord/utils"
)

type Txt struct {
	Project      string
	User         string
	Achievements string
	Goal         string
}

type Text struct {
	gorm.Model
	Txt
}

func QueryTextByProject(project string) (textList []Text) {
	utils.DB.Table("texts").Where("project = ?", project).Find(textList)
	return textList
}

func QueryTextById(id int) (textList Text) {
	utils.DB.Table("texts").Where("id = ?", id).Find(textList)
	return textList
}

func QueryTextByPage(page int) (textList []Text) {
	pageSize, _ := beego.AppConfig.Int("textsperpage")
	utils.DB.Table("texts").Offset((page - 1) * pageSize).Limit(pageSize).Find(&textList)
	return textList
}

func QueryTextSum() (sum int) {
	utils.DB.Table("texts").Count(&sum)
	return sum
}

func InsertText(text Text) error {
	err := utils.DB.Table("texts").Create(&text).Error
	return err
}

func CountTextByPara(para string) []string {
	var textList []Text
	var txtList []string
	utils.DB.Table("texts").Select("project").Find(&textList)
	for _, txt := range textList {
		txtList = append(txtList, txt.Project)
	}
	return txtList
}

func UpdateText(text Text) error {
	err := utils.DB.Table("texts").Update(text).Error
	return err
}

func DeleteTextById(id int) error {
	err := utils.DB.Table("texts").Delete(id).Error

	return err
}
