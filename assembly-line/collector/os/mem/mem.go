package mem

import (
	"context"
	"database/sql"

	proto "github.com/micro-in-cn/platform-web/assembly-line/protobuf/go/mem"
	"github.com/micro-in-cn/platform-web/internal/db"
	"github.com/micro/cli"
	"github.com/micro/go-micro/server"
)

var (
	o  *sql.DB
	ms *storage
)

type collector struct {
}

func (c *collector) PushVirtualMemoryStat(ctx context.Context, req *proto.MemRequest, rsp *proto.MemResponse) (err error) {
	err = ms.saveVirtualMemoryStat(req.VirtualMemoryStat, req.IP, req.NodeName)
	return
}

func (c *collector) PushSwapMemoryStat(ctx context.Context, req *proto.MemRequest, rsp *proto.MemResponse) (err error) {
	err = ms.saveSwapMemoryStat(req.SwapMemoryStat, req.IP, req.NodeName)
	return
}

func Init(server server.Server, ctx *cli.Context) {
	proto.RegisterMemServiceHandler(server, new(collector))
	o = db.GetPG()
	ms = new(storage)
}
