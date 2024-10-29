package models

import (
	"gorm.io/gorm"
	"time"
)

type Quote struct {
	gorm.Model
	Quote		string		`json:"quote" gorm:"text;not null;default:null"`
	Author		string		`json:"author" gorm:"text;not null;default:null"`
	LastQOTD	time.Time	`json:"lastqotd" gorm:"column:lastqotd"`
}