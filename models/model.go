package models

import (
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type BaseModel struct {
	ID         int       `gorm:"primary_key" json:"id"`
	IsValid    int       `gorm:"default:1" json:"is_valid"`
	CreateTime time.Time `gorm:"column:create_time;autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"update_time"`
}
type ProductInfo struct {
	BaseModel
	Code              string `json:"code"`
	ProductName       string `json:"product_name"`
	Specification     string `json:"specification"`
	Remark            string `json:"remark"`
	ProductUnitInfoID int
}

func (P *ProductInfo) TableName() string {
	return "productinfo"
}

type ProductUnitInfo struct {
	BaseModel
	Code     string `json:"code"`
	UnitName string `json:"unit_name"`
	Remark   string `json:"remark"`
	//ProductInfoID int
	ProductInfo ProductInfo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (P *ProductUnitInfo) TableName() string {
	return "productunitinfo"
}

type LineInfo struct {
	BaseModel
	Code            string `json:"code"`
	LineName        string `json:"line_name"`
	Reamrk          string `json:"reamrk"`
	WorkStationInfo WorkStationInfo
}

func (L *LineInfo) TableName() string {
	return "lineinfo"

}

type WorkProcessInfo struct {
	BaseModel
	Code            string `json:"code"`
	WorkProcessName string `json:"work_process_name"`
	Remark          string `json:"remark"`
	WorkStationInfo WorkStationInfo
	WorkCraftInfos  *[]WorkCraftInfo `json:"WorkCraftInfos" gorm:"many2many:work_craft_process;"`
}

func (W *WorkProcessInfo) TableName() string {
	return "workprocessinfo"

}

type WorkCraftInfo struct {
	BaseModel
	Code             string             `json:"code"`
	WorkCraftName    string             `json:"work_craft_name"`
	Remark           string             `json:"remark"`
	WorkProcessInfos *[]WorkProcessInfo `json:"WorkProcessInfos" gorm:"many2many:work_craft_process;"`
}

func (W *WorkCraftInfo) TableName() string {
	return "workcraftinfo"

}

type WorkStationInfo struct {
	BaseModel
	Code              string `json:"code"`
	WorkStationName   string `json:"work_station_name"`
	Remark            string `json:"remark"`
	WorkProcessInfoID int    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	LineInfoID        int    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (W *WorkStationInfo) TableName() string {
	return "workstationinfo"

}

type Users struct {
	BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Users
	jwt.StandardClaims
}
