package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction() // 生产环境
	//logger, _ := zap.NewDevelopment() // 开发环境
	defer logger.Sync() // flushes buffer, if any
	url := "https://google.com"

	// sugar 用法  通常系统用这个就够了
	//sugar := logger.Sugar()
	//sugar.Infow("failed to fetch URL",
	//	// Structured context as loosely typed key-value pairs.
	//	"url", url,
	//	"attempt", 3,
	//)
	//sugar.Infof("Failed to fetch URL: %s", url)

	// logger 用法 追求极致性能用logger
	logger.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("nums", 3),
	)
}
