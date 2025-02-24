package main

import "github.com/sirupsen/logrus"

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

//TODO start microservice
func main() {
	//TODO init config obj

	//TODO init logger

	//TODO init app

	//TODO start grpc-Server

	//TODO graceful shutdown
}

func setupLogger(env string) *logrus.Logger {
	//TODO init logger in local, dev, prod level
}
