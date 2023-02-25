package models

import (
	"time"

	"gorm.io/gorm"
)

type Quest struct {
	gorm.Model
	Title       string        `json:"title" gorm:"text;not null;default:null`
	Type        string        `json:"type" gorm:"text;not null;default:null`
	Information string        `json:"information" gorm:"text;not null;default:null`
	Duration    time.Duration `json:"duration"`
	Difficulty  string        `json:"difficulty" gorm:"text;not null;default:null`
}
