package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	logFolder = "log"
	logFile   = "go.log"
)

var (
	logger  *zap.Logger
	syncOne sync.Once
)

func init() {
	_, err := os.Stat(logFolder)
	if os.IsNotExist(err) {
		if err := os.Mkdir(logFolder, os.ModePerm); err != nil {
			panic(err)
		}
	}

	// if logger == nil {
	// 	NewLogger(level.Debug)
	// }
}
func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	g := gin.Default()
	g.Use(LoggingMiddleware(logger))
	//h := handlers.NewUser()
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	http.ListenAndServe(":8080", g)
}
func GetInforUser(c *gin.Context) {
	fmt.Println("hello")
}
func LoggingMiddleware(logger *zap.Logger) func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Since(start)

		logger.Info("Completed api call",
			zap.String("address", c.Request.RequestURI),
			zap.String("duration", end.String()))
	}
}

// var sugarLogger *zap.SugaredLogger

// func main() {
// 	InitLogger()
// 	defer sugarLogger.Sync()
// 	simpleHttpGet("www.sogo.com")
// 	simpleHttpGet("http://www.sogo.com")
// }

// func InitLogger() {
// 	writeSyncer := getLogWriter()
// 	encoder := getEncoder()
// 	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

// 	logger := zap.New(core, zap.AddCaller())
// 	sugarLogger = logger.Sugar()
// }

// func getEncoder() zapcore.Encoder {
// 	encoderConfig := zap.NewProductionEncoderConfig()
// 	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
// 	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
// 	return zapcore.NewConsoleEncoder(encoderConfig)
// }

// func getLogWriter() zapcore.WriteSyncer {
// 	lumberJackLogger := &lumberjack.Logger{
// 		Filename:   "./test.log",
// 		MaxSize:    1,
// 		MaxBackups: 5,
// 		MaxAge:     30,
// 		Compress:   false,
// 	}
// 	return zapcore.AddSync(lumberJackLogger)
// }

// func simpleHttpGet(url string) {
// 	sugarLogger.Debugf("Trying to hit GET request for %s", url)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
// 	} else {
// 		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
// 		resp.Body.Close()
// 	}
// }
