package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/akif999/gocr"
	"github.com/lileio/lile"
)

var s = GocrServer{}
var cli gocr.GocrClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		gocr.RegisterGocrServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = gocr.NewGocrClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
