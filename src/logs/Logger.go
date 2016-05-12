package logs

import (
	"common"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"task"
	"time"
)

type logLevel int

var checkTimeOnce sync.Once

const oneDay int64 = 60 * 60 * 24
const (
	Debug_Level logLevel = iota //0
	Info_Level                  //1
	Warn_Level                  //2
	Error_Level                 //3
)

var curLevel logLevel = Debug_Level
var logFilePath string
var ticker bool = false
var writeFile *os.File
var createErr error

type Logger struct {
}

func (log Logger) SetLevel(level logLevel) {
	curLevel = level
}
func (log Logger) SetLogFilePath(path string) {
	checkTimeOnce.Do(func() {
		fmt.Println("创建文件" + path)
		logFilePath = path
		err := os.MkdirAll(logFilePath, os.ModeDir)
		if err != nil {
			panic(err.Error())
		}
		createFile()

		t := time.Now()
		sche := time.Date(t.Year(), t.Month(), t.Day(), 00, 00, 01, 0, time.Local)
		init := sche.Unix() - t.Unix()
		task.NewSchedule(init, oneDay, func() {
			createFile()
		})
	})
}
func createFile() {
	if writeFile != nil {
		writeFile.Close()
	}
	curDate := time.Now().Format("2006-01-02")
	storePath := fmt.Sprint(logFilePath, "/", curDate+".log")
	if writeFile, createErr = common.Append_CreateFile(storePath); createErr != nil {
		writeFile.Close()
	}
}
func Info(v ...interface{}) {
	if curLevel >= Debug_Level && curLevel < Warn_Level {
		write("INFO", v...)
	}
}
func Debug(v ...interface{}) {
	if curLevel >= Debug_Level && curLevel < Info_Level {
		write("Debug", v...)
	}
}
func Error(v ...interface{}) {
	if curLevel >= Debug_Level {
		write("Error", v...)
	}
}
func Warn(v ...interface{}) {
	if curLevel >= Debug_Level && curLevel < Error_Level {
		write("Warn", v...)
	}
}
func write(levelStr string, v ...interface{}) {
	//2006-01-02 15:04:05
	if _, file, line, ok := runtime.Caller(2); ok {
		//2006-01-02 15:04:05
		if writeFile != nil {
			_, err := writeFile.WriteString(time.Now().Format("15:04:05") + " " + levelStr + " " + filepath.Base(file) + ":" + strconv.Itoa(line) + " - " + fmt.Sprint(v...) + "\n")
			if err != nil {
				writeFile.Close()
			}
		} else {
			fmt.Println(time.Now().Format("15:04:05"), levelStr, filepath.Base(file)+":"+strconv.Itoa(line), "-", fmt.Sprint(v...))
		}
		if common.Equals(levelStr,"Error"){
			buf := make([]byte, 1 << 20)
			runtime.Stack(buf, false)
			fmt.Printf("呵呵\n%s", buf)
		}
	}
}
