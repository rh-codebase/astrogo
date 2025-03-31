package astrounit

type MonitorPointType int

const (
	FLOAT64 MonitorPointType = iota
	FLOAT32
	INT64
	INT32
	INT16
	INT8
	STRING
	BOOL
)
type Number interface {
    int64 | float64
}

type MonitorPoint struct {
	Ssid  int32
	Id    int32
	Type  MonitorPointType
	Value interface{}
}

type MonitorPointMeta struct {
  // subsystem id
	Ssid       int32
  // monitor point id within subsystem
	Id         int32
	Name       string
	ShortName  string
	Unit       MonitorPointType
	MaxWarning interface{}
	MinWarning interface{}
	MaxError   interface{}
	MinError   interface{}
}

func (mpm MonitorPointMeta) GetMaxWarning() interface{} {
	return mpm.MaxWarning
}

