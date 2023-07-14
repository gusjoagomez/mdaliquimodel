package models

import "github.com/shopspring/decimal"

type F572Retperpagos struct {
	//gorm.Model       // adds ID, created_at etc.
	Id                int              `json:"id" gorm:"primaryKey;not null;size:32;type:int4;column:id"`
	Presentacion_id   uint             `json:"presentacion_id" gorm:"type:int4;column:presentacion_id"`
	F572Presentacion1 F572Presentacion `gorm:"foreignKey:Presentacion_id;references:Id"`
	Tipo              string           `json:"tipo" gorm:"type:varchar(2);column:tipo"`
	Tipodoc           uint             `json:"tipodoc" gorm:"type:int4;column:tipodoc"`
	Nrodoc            string           `json:"nrodoc" gorm:"type:varchar(15);column:nrodoc"`
	Denominacion      string           `json:"denominacion" gorm:"type:varchar(200);column:denominacion"`
	Descbasica        string           `json:"descbasica" gorm:"type:varchar(300);column:descbasica"`
	Descadicional     string           `json:"descadicional" gorm:"type:varchar(300);column:descadicional"`
	Montototal        decimal.Decimal  `json:"montototal" gorm:"type:numeric(13,2);column:montototal"`
}

func (e *F572Retperpagos) TableName() string {
	return "siradig.f572_retperpagos"
}
