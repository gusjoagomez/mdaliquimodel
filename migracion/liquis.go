package models

import (
	"time"
)

type Liquis struct {
	//gorm.Model       // adds ID, created_at etc.
	ID               uint       `json:"id" gorm:"primaryKey;not null;size:32;type:int4;column:id"`
	Created_at       time.Time  `json:"created_at" gorm:"column:created_at"`
	Updated_at       *time.Time `json:"updated_at" gorm:"column:updated_at"`
	Deleted_at       *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Id_agr           uint       `json:"id_agr" gorm:"type:int8;column:id_agr"`
	Id_plaper        uint       `json:"id_plaper" gorm:"type:int4;column:id_plaper"`
	Anio_liq         uint       `json:"anio_liq" gorm:"type:int4;column:anio_liq"`
	Mes_liq          uint       `json:"mes_liq" gorm:"type:int4;column:mes_liq"`
	Anio_dev         uint       `json:"anio_dev" gorm:"type:int4;column:anio_dev"`
	Mes_dev          uint       `json:"mes_dev" gorm:"type:int4;column:mes_dev"`
	Anio_gan         uint       `json:"anio_gan" gorm:"type:int4;column:anio_gan"`
	Mes_gan          uint       `json:"mes_gan" gorm:"type:int4;column:mes_gan"`
	Tipo_liq         string     `json:"tipo_liq" gorm:"column:tipo_liq"`
	Stipo_liq        uint       `json:"stipo_liq" gorm:"column:stipo_liq"`
	Fecha_pago       time.Time  `json:"fecha_pago" gorm:"column:fecha_pago"`
	Descripcion      string     `json:"descripcion" gorm:"column:descripcion"`
	Anulacion        string     `json:"anulacion" gorm:"column:anulacion"`
	Fecha_liq        string     `json:"fecha_liq" gorm:"column:fecha_liq"`
	Id_estado        uint       `json:"id_estado" gorm:"column:id_estado"`
	Nro_legajo_desde uint       `json:"nro_legajo_desde" gorm:"type:int4;column:nro_legajo_desde"`
	Nro_legajo_hasta uint       `json:"nro_legajo_hasta" gorm:"type:int4;column:nro_legajo_hasta"`
	Fec_ini_proc     time.Time  `json:"fec_ini_proc" gorm:"column:fec_ini_proc"`
	Fec_fin_proc     *time.Time `json:"fec_fin_proc" gorm:"column:fec_fin_proc"`
}

func (e *Liquis) TableName() string {
	return "migracion.liquis"
}
