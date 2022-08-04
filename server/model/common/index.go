package common

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"server/utils"
	"strings"
)

type DeletedAt sql.NullTime
type Model struct {
	ID         string           `json:"id" gorm:"primarykey;size:32;comment:主键"`
	CreateTime *utils.LocalTime `json:"create_time" gorm:"type:timestamp;default:current_timestamp;<-:create;comment:创建时间"`
	UpdateTime *utils.LocalTime `json:"update_time" gorm:"type:timestamp;default:current_timestamp on update current_timestamp;comment:更新时间"`
	DeleteTime gorm.DeletedAt   `json:"delete_time" gorm:"type:timestamp;comment:删除时间"`
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == "" {
		id := uuid.NewString()
		id = strings.ReplaceAll(id, "-", "")
		m.ID = id
	}

	m.CreateTime = nil
	m.UpdateTime = nil
	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	m.CreateTime = nil
	m.UpdateTime = nil
	return
}
