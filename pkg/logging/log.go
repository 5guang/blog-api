package logging

import (
	"blog/pkg/file"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix = ""
	DefaultCallerDepath = 2

	logger *log.Logger
	logPrefix = ""
	LevelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR","FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Init() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{})  {
	setPrefix(DEBUG)
	logger.Println(v)
}
func Info(v ...interface{})  {
	setPrefix(INFO)
	logger.Println(v)
}
func Warn(v ...interface{})  {
	setPrefix(WARNING)
	logger.Println(v)
}
func Error(v ...interface{})  {
	setPrefix(ERROR)
	logger.Println(v)
}
func Fatal(v ...interface{})  {
	setPrefix(FATAL)
	logger.Println(v)
}

func setPrefix(level Level)  {
	_, file, line, ok := runtime.Caller(DefaultCallerDepath)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", LevelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", LevelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}