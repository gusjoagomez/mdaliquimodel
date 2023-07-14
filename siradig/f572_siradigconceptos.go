package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type F572Siradigconceptos struct {
	//gorm.Model       // adds ID, created_at etc.
	Id        int             `json:"id" gorm:"primaryKey;not null;size:32;type:int4;column:id"`
	Legajo    uint            `json:"legajo" gorm:"not null;type:int8;default:0;column:legajo"`
	Cuit      uint            `json:"cuit" gorm:"not null;type:int8;column:cuit"`
	Anio      uint            `json:"anio" gorm:"not null;type:int4;column:anio"`
	Mes       uint            `json:"mes" gorm:"not null;type:smallint;column:mes"`
	Nrover    uint            `json:"nrover" gorm:"not null;type:int4;column:nrover"`
	Atributo  string          `json:"atributo" gorm:"not null;type:varchar(30);column:atributo"`
	Valor     decimal.Decimal `json:"valor" gorm:"type:numeric(10,2);column:valor"`
	CreatedAt time.Time       `json:"createdat" gorm:"type:timestamp;column:createdat;not null"`
	UpdatedAt time.Time       `json:"updatedat" gorm:"type:timestamp;column:updatedat;not null"`
	Estadoliq string          `json:"estadoliq" gorm:"type:varchar(1);column:estadoliq"`
}

func (e *F572Siradigconceptos) TableName() string {
	return "siradig.f572_siradigconceptos"
}
