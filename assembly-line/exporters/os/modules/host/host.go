package host

import (
	"sync"
	"time"

	"github.com/micro-in-cn/platform-web/assembly-line/exporters/os/modules"
	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/host"
)

var (
	once sync.Once
)

type Host struct {
	hostClient proto.HostService
	modules.BaseModule
	opts *modules.HostOptions
}

func (h *Host) Init(opts *modules.Options) {
	h.opts = opts.Host
	h.InitB()
	h.hostClient = proto.NewHostService(opts.Collector.Name, opts.Collector.Client)

	return
}

func (h *Host) Push() (err error) {
	err = h.pushInfo()
	return err
}

func (h *Host) Interval() time.Duration {
	return h.opts.Interval
}

func (h *Host) String() string {
	return "host"
}
