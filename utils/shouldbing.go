package utils

type LineCreateBind struct {
	LineCode string `form:"line_code" binding:"required"`
	LineName string `form:"line_name" binding:"required"`
	Remark   string `form:"remark"`
}

type LineGetBind struct {
	LineCode string `form:"line_code"`
	LineName string `form:"line_name"`
}
type LineUpdateBind struct {
	LineId   int    `form:"id"`
	LineCode string `form:"line_code"`
	LineName string `form:"line_name"`
	Remark   string `form:"remark"`
}
type StationBind struct {
	StationId   int    `form:"id"`
	StationCode string `form:"station_code"`
	StationName string `form:"station_name"`
	Remark      string `form:"remark"`
	LineId      int    `form:"line_id"`
	ProcessId   int    `form:"process_id"`
}
