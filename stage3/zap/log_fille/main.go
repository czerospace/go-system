package main

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./server.log",
	}
	return cfg.Build()
}

func main() {
	//logger, _ := zap.NewProduction()
	logger, err := NewLogger()
	if err != nil {
		panic(err)
		//panic("初始化logger失败")
	}
	defer logger.Sync()

	url := "https://google.com"
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("nums", 3),
	)
}
