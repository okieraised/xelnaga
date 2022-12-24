package postgres

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DCMObject struct {
	bun.BaseModel `bun:"table:dcm_objects"`

	ID                uuid.UUID `bun:",pk,type:uuid"`
	StudyInstanceUID  string    `bun:",type:string,notnull"`
	SeriesInstanceUID string    `bun:",type:string,notnull"`
	SOPInstanceUID    string    `bun:",type:string,notnull"`
	FileLocation      string    `bun:",type:string,notnull"`
	Created           int64     `bun:"type:integer"`
	Updated           int64     `bun:"type:integer"`
}

type DCMStudy struct {
}

type DCMSeries struct {
}

type PatientIOD struct {
	PatientName               string
	PatientID                 string
	PatientIDType             string
	Birthdate                 string
	Gender                    string
	QualityControlSubject     string
	ReferencedPatientSequence []string
}
