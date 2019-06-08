package cpu

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/third_party/gopsutil/cpu"
	cpu2 "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/cpu"
)

func (p *Pusher) pushInfo() (err error) {

	vv, err := cpu.Info()
	if err != nil {
		return fmt.Errorf("[pushInfo] get infos error: %s", err)
	}

	data := make([]*cpu2.InfoStat, len(vv))
	t := ptypes.TimestampNow()

	for _, v := range vv {
		data = append(data, &cpu2.InfoStat{
			Timestamp:  t,
			Cpu:        v.CPU,
			VendorId:   v.VendorID,
			Family:     v.Family,
			Model:      v.Model,
			Stepping:   v.Stepping,
			PhysicalId: v.PhysicalID,
			CoreId:     v.CoreID,
			Cores:      v.Cores,
			ModelName:  v.ModelName,
			Mhz:        v.Mhz,
			CacheSize:  v.CacheSize,
			Flags:      v.Flags,
			Microcode:  v.Microcode,
		})
	}

	req := &cpu2.CPURequest{
		Timestamp: t,
		IP:        p.IP,
		NodeName:  p.NodeName,
		InfoStat:  data,
	}

	_, err = p.cpuClient.PushCPUInfoStat(context.Background(), req)
	if err != nil {
		return fmt.Errorf("[pushInfo] push error: %s", err)
	}

	return
}
