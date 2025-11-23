// pkg/logger/logger.go
package logger

import (
    "os"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

func InitLogger() {
    // 配置编码器
    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.TimeKey = "time"
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    encoderConfig.LevelKey = "level"
    encoderConfig.MessageKey = "msg"
    
    // 设置日志级别
    level := zap.NewAtomicLevel()
    level.SetLevel(zap.InfoLevel) // 根据环境变量调整
    
    // 创建核心
    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),
        zapcore.AddSync(os.Stdout),
        level,
    )
    
    // 创建Logger
    Logger = zap.New(core, zap.AddCaller())
    SugarLogger = Logger.Sugar()
    
    defer Logger.Sync()
}

// 不同环境的初始化
func InitDevelopmentLogger() {
    config := zap.NewDevelopmentConfig()
    config.EncoderConfig.TimeKey = "time"
    config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    
    var err error
    Logger, err = config.Build()
    if err != nil {
        panic(err)
    }
    SugarLogger = Logger.Sugar()
}