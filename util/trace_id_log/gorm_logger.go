package trace_id_log

import (
	"TikTokLite_v2/util/trace_id_log/loggers"
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

var _ = logger.Interface(&GormLogger{})

type GormLogger struct {
	logLevel                                           logrus.Level
	slowThreshold                                      time.Duration
	ignoreRecordNotFoundErr                            bool
	infoStr, warnStr, errStr                           string
	traceStr, traceWarnStr, traceErrStr                string
	colorInfoStr, colorWarnStr, colorErrStr            string
	colorTraceStr, colorTraceWarnStr, colorTraceErrStr string
}

func NewGormLogger(level logger.LogLevel) logger.Interface {
	var (
		infoStr      = "%s\n[info] "
		warnStr      = "%s\n[warn] "
		errStr       = "%s\n[error] "
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)
	colorInfoStr := logger.Green + "%s\n" + logger.Reset + logger.Green + "[info] " + logger.Reset
	colorWarnStr := logger.BlueBold + "%s\n" + logger.Reset + logger.Magenta + "[warn] " + logger.Reset
	colorErrStr := logger.Magenta + "%s\n" + logger.Reset + logger.Red + "[error] " + logger.Reset
	colorTraceStr := logger.Green + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	colorTraceWarnStr := logger.Green + "%s " + logger.Yellow + "%s\n" + logger.Reset + logger.RedBold + "[%.3fms] " + logger.Yellow + "[rows:%v]" + logger.Magenta + " %s" + logger.Reset
	colorTraceErrStr := logger.RedBold + "%s " + logger.MagentaBold + "%s\n" + logger.Reset + logger.Yellow + "[%.3fms] " + logger.BlueBold + "[rows:%v]" + logger.Reset + " %s"
	return &GormLogger{
		logLevel:          logrus.Level(level),
		infoStr:           infoStr,
		warnStr:           warnStr,
		errStr:            errStr,
		traceStr:          traceStr,
		traceWarnStr:      traceWarnStr,
		traceErrStr:       traceErrStr,
		colorInfoStr:      colorInfoStr,
		colorErrStr:       colorErrStr,
		colorWarnStr:      colorWarnStr,
		colorTraceStr:     colorTraceStr,
		colorTraceErrStr:  colorTraceErrStr,
		colorTraceWarnStr: colorTraceWarnStr,
	}
}
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.logLevel = logrus.Level(level)
	return l
}
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logrus.InfoLevel {
		l.printf(logrus.InfoLevel, msg, data...)
	}
}
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logrus.WarnLevel {
		l.printf(logrus.WarnLevel, msg, data...)
	}
}
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logrus.ErrorLevel {
		l.printf(logrus.ErrorLevel, msg, data...)
	}
}
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && l.logLevel >= logrus.ErrorLevel && (!errors.Is(err, logger.ErrRecordNotFound) || !l.ignoreRecordNotFoundErr):
		sql, rows := fc()
		if rows == -1 {
			l.printf(logrus.ErrorLevel, l.colorTraceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.printf(logrus.ErrorLevel, l.colorTraceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.slowThreshold && l.slowThreshold != 0 && l.logLevel >= logrus.WarnLevel:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.slowThreshold)
		if rows == -1 {
			l.printf(logrus.WarnLevel, l.colorTraceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.printf(logrus.WarnLevel, l.colorTraceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.logLevel >= logrus.InfoLevel:
		sql, rows := fc()
		if rows == -1 {
			l.printf(logrus.InfoLevel, l.colorTraceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.printf(logrus.InfoLevel, l.colorTraceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
func (l *GormLogger) printf(level logrus.Level, msg string, data ...interface{}) {
	switch level {
	case logrus.InfoLevel:
		loggers.Logger.Infof(msg, data...)
	case logrus.WarnLevel:
		loggers.Logger.Warnf(msg, data...)
	case logrus.ErrorLevel:
		loggers.Logger.Errorf(msg, data...)
	case logrus.TraceLevel:
		loggers.Logger.Tracef(msg, data...)
	case logrus.DebugLevel:
		loggers.Logger.Debugf(msg, data)
	}
}
