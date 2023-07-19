package models

type DetallesHab struct {
	TipoReg                     int    //1=haberes, 2=SAC
	Mes                         int    //1-12
	Anio                        int    //1995-2023
	TipoEmp                     string //D: Docentes, H: HorasCátedras, P: Prejubilados, B: Becarios, L: Limpieza (camioneros), A: Resto de empleados (planta, mensualizados, médicos)
	Legajo                      int
	NombreApellido              string
	Documento                   int
	Cuil                        int
	GrupoLiquidacion            int
	DescripcionGrupoLiquidacion string
	CodigoTipoPlanta            int
	DescripcionTipoPlanta       string
	CodigoTipoPuesto            int
	DescripcionTipoPuesto       string
	CodigoGrupoPersonal         int
	DescripcionGrupoPersonal    string
	CodigoFuncion               int
	DescripcionFuncion          string
	Codep                       string
	DescripcionCodep            string
	CodigoConcepto              int
	DescripcionConcepto         string
	CodigoEntidad               int
	DescripcionEntidad          string
	Unidades                    float64
	Importe                     float64
	AnomesRetroactivo           int
	Escalafon                   int
	Categoria                   int
	Descripcion                 string
}

func (e *DetallesHab) TableName() string {
	return "migracion.DetallesHab"
}
