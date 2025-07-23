package models

import (
	"time"
)

type Computer struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"type:varchar(100);not null"`
	Description  string `gorm:"type:text"`
	MacAddress   string `gorm:"type:varchar(17);uniqueIndex;not null"`
	IPAddress    string `gorm:"type:varchar(15);not null"`
	EmployeeCode string `gorm:"type:varchar(10)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
