package main

import (
	_ "net/http/pprof"

	"github.com/akif999/gocr"
	"github.com/akif999/gocr/gocr/cmd"
	"github.com/akif999/gocr/server"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.GocrServer{}

	lile.Name("gocr")
	lile.Server(func(g *grpc.Server) {
		gocr.RegisterGocrServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
