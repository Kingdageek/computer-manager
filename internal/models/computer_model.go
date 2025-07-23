package models

import (
	"computer-manager/internal/dtos"
	"time"
)

type Computer struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	Name         string  `gorm:"type:varchar(100);not null"`
	Description  *string `gorm:"type:text"`
	MacAddress   string  `gorm:"type:varchar(17);uniqueIndex;not null"`
	IPAddress    string  `gorm:"type:varchar(15);not null"`
	EmployeeCode *string `gorm:"type:varchar(10)"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (cmp *Computer) ToDto() *dtos.ComputerDto {
	return &dtos.ComputerDto{
		ID:           cmp.ID,
		Name:         cmp.Name,
		Description:  cmp.Description,
		MacAddress:   cmp.MacAddress,
		IPAddress:    cmp.IPAddress,
		EmployeeCode: cmp.EmployeeCode,
		CreatedAt:    cmp.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    cmp.UpdatedAt.Format(time.RFC3339),
	}
}
