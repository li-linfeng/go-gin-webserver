package utils

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"meta/pkg/setting"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type logger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	LogLevel              gormlogger.LogLevel
	Log                   *logrus.Logger
}

func GormLogger() *logger {
	file, fileName := GetGormLoggerFileSrc()
	log := NewLogrus()
	log.Out = file

	// 设置 rotatelogs
	Rotatelogs(fileName, log)

	l := &logger{
		SkipErrRecordNotFound: true,
		LogLevel:              gormlogger.Silent,
		SlowThreshold:         time.Microsecond,
		Log:                   log,
	}
	return l
}

func NewLogrus() *logrus.Logger {
	logger := logrus.New()
	return logger
}

func GetGormLoggerFileSrc() (*os.File, string) {
	fileName := path.Join(setting.AppSetting.LogPath, "/", setting.AppSetting.SqlDir, "/", setting.AppSetting.AppName)
	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("err", err)
	}

	return src, fileName
}

func GetFrameWorkLoggerFileSrc() (*os.File, string) {
	fileName := path.Join(setting.AppSetting.LogPath, "/", setting.AppSetting.DailyDir, "/", setting.AppSetting.AppName)
	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("err", err)
	}

	return src, fileName
}

func Rotatelogs(fileName string, logger *logrus.Logger) {
	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		// rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))
}

func (l *logger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *logger) Info(ctx context.Context, s string, args ...interface{}) {
	l.Log.WithContext(ctx).Infof(s, args)
}

func (l *logger) Warn(ctx context.Context, s string, args ...interface{}) {
	l.Log.WithContext(ctx).Warnf(s, args)
}

func (l *logger) Error(ctx context.Context, s string, args ...interface{}) {
	l.Log.WithContext(ctx).Errorf(s, args)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}

	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}

	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		l.Log.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.Log.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	l.Log.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}
