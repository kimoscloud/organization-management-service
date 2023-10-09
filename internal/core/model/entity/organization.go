package entity

import (
	"gorm.io/gorm"
	"time"
)

type Organization struct {
	ID                         string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name                       string         `gorm:"column:name"`
	Email                      string         `gorm:"column:email"`
	AcceptTermsAndConditions   bool           `gorm:"column:accept_terms_and_conditions"`
	AcceptTermsAndConditionsAt time.Time      `gorm:"column:accept_terms_and_conditions_at"`
	PhotoUrl                   string         `gorm:"column:photo_url"`
	Phone                      string         `gorm:"column:phone"`
	Timezone                   string         `gorm:"column:timezone"`
	CreatedAt                  time.Time      `gorm:"column:created_at"`
	UpdatedAt                  time.Time      `gorm:"column:updated_at"`
	DeletedAt                  gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Organization) TableName() string {
	return "Organizations"
}
