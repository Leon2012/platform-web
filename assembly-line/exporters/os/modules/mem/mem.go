package mem

import (
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
	"github.com/micro/go-micro/util/log"
)

type Mem struct {
	modules.BaseModule
	opts      *modules.MemOptions
	memClient proto.MemService
}

func (m *Mem) Init(opts *modules.Options) {
	log.Logf("[INFO] [Init] init memory module")
	m.opts = opts.Mem
	m.InitB()
	m.memClient = proto.NewMemService(opts.Collector.Name, opts.Collector.Client)
	return
}

func (m *Mem) Push() (err error) {
	err = m.pushMemoryStat()
	if err != nil {
		return
	}

	return err
}

func (m *Mem) Interval() time.Duration {
	return m.opts.Interval
}

func (m *Mem) String() string {
	return "mem"
}
