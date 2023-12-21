package models

import (
	"time"
)

type Liquidacion struct {
	//gorm.Model       // adds ID, created_at etc.
	Activado                 bool       `json:"activado" gorm:"column:activado"`
	CreadoEn                 time.Time  `json:"creadoEn" gorm:"column:creadoEn"`
	ModificadoEn             *time.Time `json:"modificadoEn" gorm:"column:modificadoEn"`
	EliminadoEn              *time.Time `json:"eliminadoEn" gorm:"column:eliminadoEn"`
	ActivadoEn               time.Time  `json:"activadoEn" gorm:"type:timestamp;column:activadoEn"`
	ID                       uint       `json:"id" gorm:"primaryKey;not null;size:32;type:int4;column:id"`
	AnioMesLiquidador        string     `json:"anioMesLiquidador" gorm:"type:varchar(6);column:anioMesLiquidador"`
	AnioMesDevengado         string     `json:"anioMesDevengado" gorm:"type:varchar(6);column:anioMesDevengado"`
	AnioMesPago              string     `json:"anioMesPago" gorm:"type:varchar(6);column:anioMesPago"`
	SubtipoLiquidacion       uint       `json:"subtipoLiquidacion" gorm:"column:subtipoLiquidacion"`
	Anulacion                string     `json:"anulacion" gorm:"column:anulacion"`
	NumeroLegajoDesde        uint       `json:"numeroLegajoDesde" gorm:"type:int4;column:numeroLegajoDesde"`
	NumeroLegajoHasta        uint       `json:"numeroLegajoHasta" gorm:"type:int4;column:numeroLegajoHasta"`
	NumeroCargoDesde         uint       `json:"numeroCargoDesde" gorm:"type:int4;column:numeroCargoDesde"`
	NumeroCargoHasta         uint       `json:"numeroCargoHasta" gorm:"type:int4;column:numeroCargoHasta"`
	FechaDesde               time.Time  `json:"fechaDesde" gorm:"column:fechaDesde"`
	FechaPago                time.Time  `json:"fechaPago" gorm:"column:fechaPago"`
	FechaLiquidacion         string     `json:"fechaLiquidacion" gorm:"column:fechaLiquidacion"`
	FechaInicioProcedimiento time.Time  `json:"fechaInicioProcedimiento" gorm:"column:fechaInicioProcedimiento"`
	FechaFinProcedimiento    *time.Time `json:"fechaFinProcedimiento" gorm:"column:fechaFinProcedimiento"`
	Descripcion              string     `json:"descripcion" gorm:"column:descripcion"`
	TipoLiquidacionId        uint       `json:"tipoLiquidacionId" gorm:"column:tipoLiquidacionId"`
	EstadoId                 uint       `json:"estadoId" gorm:"column:estadoId"`
	EstadoProcesoId          uint       `json:"estadoProcesoId" gorm:"column:estadoProcesoId"`
	GrupoCargoId             uint       `json:"grupoCargoId" gorm:"type:int8;column:grupoCargoId"`
	SubgrupoCargoId          uint       `json:"subgrupoCargoId" gorm:"type:int8;column:subgrupoCargoId"`
	PlantaPersonalId         uint       `json:"plantaPersonalId" gorm:"type:int4;column:plantaPersonalId"`
	CreadoPorId              int        `json:"creadoPorId" gorm:"type:int4;column:creadoPorId"`
	ModificadoPorId          int        `json:"modificadoPorId" gorm:"type:int4;column:modificadoPorId"`
	EliminadoPorId           int        `json:"eliminadoPorId" gorm:"type:int4;column:eliminadoPorId"`
	ActivadoPorId            int        `json:"activadoPorId" gorm:"type:int4;column:activadoPorId"`
}

func (e *Liquidacion) TableName() string {
	return "migracion.liquidacion"
}
