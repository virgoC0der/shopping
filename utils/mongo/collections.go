package mongo

const (
	KDBMonitor       = "monitor"
	KCollMonitorData = "monitor_data"
	KCollUserLog     = "user_log"
)

type MonitorData struct {
	Id        string       `json:"id"         bson:"_id"`
	AgentName string       `json:"agent_name" bson:"agent_name"`
	System    string       `json:"system"     bson:"system"`
	Network   *NetworkInfo `json:"network"    bson:"network"`
	Cpu       *CpuInfo     `json:"cpu"        bson:"cpu"`
	Mem       *MemInfo     `json:"mem"        bson:"mem"`
	Disk      *DiskInfo    `json:"disk"       bson:"disk"`
	Time      int64        `json:"time"       bson:"time"`
}

type NetworkInfo struct {
	Ips []string `json:"ips"      bson:"ips"`
	Ios int      `json:"ios"      bson:"ios"`
}

type CpuInfo struct {
	Models          []string `json:"models"            bson:"models"`
	LogicCoreNum    int      `json:"logic_core_num"    bson:"logic_core_num"`
	PhysicalCoreNum int      `json:"physical_core_num" bson:"physical_core_num"`
	Usage           float64  `json:"usage"             bson:"usage"`
}

type MemInfo struct {
	Total  uint64 `json:"total"  bson:"total"`
	Actual uint64 `json:"actual" bson:"actual"`
	Swap   uint64 `json:"swap"   bson:"swap"`
}

type DiskInfo struct {
	Total uint64 `json:"total" bson:"total"`
	Used  uint64 `json:"used"  bson:"used"`
}

type UserLog struct {
	UserId      string      `json:"user_id"      bson:"user_id"`
	Time        int64       `json:"time"         bson:"time"`
	QueryParams interface{} `json:"query_params" bson:"query_params"`
}
