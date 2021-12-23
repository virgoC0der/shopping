package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	gonet "net"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"go.uber.org/zap"

	. "shopping/utils/log"
	"shopping/utils/mongo"
	"shopping/utils/webbase"
)

const (
	kCollectApi = "http://127.0.0.1:8081/api/v1/monitor/collect"
)

func Run() {
	t := time.NewTicker(time.Minute * 1)
	for {
		err := CollectData()
		if err != nil {
			Logger.Warn("collect data err", zap.Error(err))
		}
		<-t.C
	}

}

func CollectData() error {
	hostName, err := host.HostID()
	if err != nil {
		Logger.Warn("get host id err", zap.Error(err))
		return err
	}

	system, _, _, err := host.PlatformInformation()
	if err != nil {
		Logger.Warn("get system info err", zap.Error(err))
		return err
	}

	addrs, err := gonet.InterfaceAddrs()
	if err != nil {
		Logger.Warn("get ip err", zap.Error(err))
		return err
	}
	ips := make([]string, len(addrs))
	for _, addr := range addrs {
		ips = append(ips, addr.String())
	}

	ios, err := net.IOCounters(true)
	if err != nil {
		Logger.Warn("get io counters err", zap.Error(err))
		return err
	}
	ioNum := len(ios)

	cpuInfo, err := cpu.Info()
	if err != nil {
		Logger.Warn("get cpu info err", zap.Error(err))
		return err
	}

	models := make([]string, len(cpuInfo))
	for _, c := range cpuInfo {
		models = append(models, c.ModelName)
	}

	logicCores, err := cpu.Counts(true)
	if err != nil {
		Logger.Warn("get logic cores err", zap.Error(err))
		return err
	}

	phicalCores, err := cpu.Counts(false)
	if err != nil {
		Logger.Warn("get phical cores err", zap.Error(err))
		return err
	}

	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		Logger.Warn("get cpu percent err", zap.Error(err))
		return err
	}
	usage := float64(0)
	for _, p := range cpuPercent {
		usage += p
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		Logger.Warn("get mem info err", zap.Error(err))
		return err
	}

	diskUsage, err := disk.Usage("/")
	if err != nil {
		Logger.Warn("get disk usage err", zap.Error(err))
		return err
	}

	diskTotal := diskUsage.Total
	diskUsed := diskUsage.Used

	monitorData := &mongo.MonitorData{
		AgentName: hostName,
		System:    system,
		Network: &mongo.NetworkInfo{
			Ips: ips,
			Ios: ioNum,
		},
		Cpu: &mongo.CpuInfo{
			Models:          models,
			LogicCoreNum:    logicCores,
			PhysicalCoreNum: phicalCores,
			Usage:           usage,
		},
		Mem: &mongo.MemInfo{
			Total:  memInfo.Total,
			Actual: memInfo.Available,
			Swap:   memInfo.SwapTotal,
		},
		Disk: &mongo.DiskInfo{
			Total: diskTotal,
			Used:  diskUsed,
		},
		Time: time.Now().Unix(),
	}

	data, err := json.Marshal(monitorData)
	if err != nil {
		Logger.Warn("json marshal err", zap.Error(err))
		return err
	}
	req := bytes.NewReader(data)
	request, err := http.NewRequest(http.MethodPost, kCollectApi, req)
	if err != nil {
		Logger.Warn("new request err", zap.Error(err))
		return err
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		Logger.Warn("do request err", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Logger.Warn("read response err", zap.Error(err))
		return err
	}

	response := &webbase.CommonResp{}
	err = json.Unmarshal(b, response)
	if err != nil {
		Logger.Warn("json unmarshal err", zap.Error(err))
		return err
	}

	if response.Code != 0 {
		Logger.Warn("collect data err", zap.Error(err))
		return err
	}

	return nil
}
