package models

import (
	"time"
)

type LiquiCons struct {
	//gorm.Model       // adds ID, created_at etc.
	ID            uint       `json:"id" gorm:"primaryKey;not null;size:32;type:int4;column:id"`
	Created_at    time.Time  `json:"created_at" gorm:"column:created_at"`
	Updated_at    *time.Time `json:"updated_at" gorm:"column:updated_at"`
	Deleted_at    *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Liqui_id      uint       `json:"liqui_id" gorm:"type:int8;column:liqui_id"`
	Nro_legajo    uint       `json:"nro_legajo" gorm:"type:int4;column:nro_legajo"`
	Nro_cargo     uint       `json:"nro_cargo" gorm:"type:int4;column:nro_cargo"`
	Anio_liqui    uint       `json:"anio_liqui" gorm:"type:int4;column:anio_liqui"`
	Anio_dev      uint       `json:"anio_dev" gorm:"type:int4;column:anio_dev"`
	Mes_dev       uint       `json:"mes_dev" gorm:"type:int4;column:mes_dev"`
	Fecha_desde   time.Time  `json:"fecha_desde" gorm:"column:fecha_desde"`
	Concepto_id   uint       `json:"concepto_id" gorm:"type:int8;column:concepto_id"`
	Cantidad      string     `json:"cantidad" gorm:"column:cantidad"`
	Imp_unit      uint       `json:"imp_unit" gorm:"type:numeric(12,2);column:imp_unit"`
	Imp_item      uint       `json:"imp_item" gorm:"type:numeric(12,2);column:imp_item"`
	Tipo_concepto uint       `json:"tipo_concepto" gorm:"type:int8;column:tipo_concepto"`
	Tipo_con_gan  uint       `json:"tipo_con_gan" gorm:"type:int8;column:tipo_con_gan"`
}

func (e *LiquiCons) TableName() string {
	return "migracion.liqui_cons"
}
