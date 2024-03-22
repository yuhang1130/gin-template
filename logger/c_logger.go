package logger

import (
	log "github.com/sirupsen/logrus"
)

func InitLogger(level string) *log.Logger {
	parseLevel, err := log.ParseLevel(level)
	if err != nil {
		panic(err.Error())
	}

	logger := log.New()

	// 设置日志级别
	logger.SetLevel(parseLevel)

	// 设置日志输出格式
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,                  // 展示日期
		TimestampFormat: "2006-01-02 15:04:05", //日期格式
		ForceColors:     true,                  // 颜色日志
	})

	// logger.Out = os.Stdout

	return logger
}
