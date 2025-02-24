package grpc

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type App struct {
	log        *logrus.Logger
	gRPCServer *grpc.Server
	port       int
}

// App constructor with logger and Secer
func New(log *logrus.Logger, country_Service Country, port int) *App {

}
