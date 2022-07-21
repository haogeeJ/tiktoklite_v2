package loggers

import (
	"github.com/jtolds/gls"
	"github.com/sirupsen/logrus"
	"os"
)

type LoggerContext struct {
	TraceId string
}

type GlsHook struct {
}

func (h *GlsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

var LoggerContextMgr = gls.NewContextManager()

//hook，可以添加trace_id,或者输出到磁盘
func (h *GlsHook) Fire(e *logrus.Entry) error {
	gls.EnsureGoroutineId(func(gid uint) {
		value, ok := LoggerContextMgr.GetValue(gid)
		if ok {
			loggerCtx := value.(*LoggerContext)
			if loggerCtx.TraceId != "" {
				e.Data["trace_id"] = loggerCtx.TraceId
			}
		}
	})
	//fileName := "D:\\awesomeProject\\demo\\logs\\sql.log"
	//f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	//if err != nil {
	//	logrus.Error(err)
	//	return err
	//}
	////defer f.Close()
	//if _, err := f.Write([]byte(e.Message)); err != nil {
	//	logrus.Error(err)
	//	return err
	//}
	return nil
}

var Logger *logrus.Logger

func InitLogger() {
	//fileName := "D:\\awesomeProject\\demo\\logs\\sql.log"
	//f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	//if err != nil {
	//	logrus.Error(err)
	//	return
	//}
	////defer f.Close()
	//writes := []io.Writer{
	//	os.Stdout,
	//	f,
	//}
	//fileAndStdoutWriter := io.MultiWriter(writes...)

	Logger = logrus.New()
	//可以设置输出文件
	//Logger.SetOutput(fileAndStdoutWriter)
	Logger.SetOutput(os.Stdout)
	Logger.AddHook(&GlsHook{})
	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		ForceQuote:       true,
		DisableTimestamp: true,
	})
}
