package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("fistname: %s, age:%d", "hung", 20)

	// //Owner
	// logger := zap.NewExample()
	// logger.Info("Helo", zap.String("name", "hung"), zap.Int("age", 20))

	// //Develop
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hung Test + 1")

	// //Production
	// logger, _ = zap.NewProduction()
	// logger.Info("Hung Test + 2")

	//3. Custom
	encodeder := GetEncoder()
	syncer := GetWriteSyncer()
	core := zapcore.NewCore(encodeder, syncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Hung Test + 1111")
	logger.Error("Hung Test + 4")

}

// Format log
func GetEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// Write log to file
func GetWriteSyncer() zapcore.WriteSyncer {
	// return zapcore.AddSync(&lumberjack.Logger{
	// 	Filename:   "logs/app.log",
	// 	MaxSize:    1, // megabytes
	// 	MaxBackups: 3,
	// 	MaxAge:     28, // days
	// })
	file, _ := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
