package model

import (
	"github.com/google/uuid"
)

type DCMObject struct {
	ID                uuid.UUID `gorm:"->"`
	StudyInstanceUID  string    `gorm:"->,unique,not null"`
	SeriesInstanceUID string    `gorm:"->,unique,not null"`
	SOPInstanceUID    string    `gorm:"->,unique,not null"`
	FileLocation      string    `gorm:"->,unique,not null"`
	Created           int64     `gorm:"autoCreateTime"`
	Updated           int64     `gorm:"autoCreateTime"`
}
