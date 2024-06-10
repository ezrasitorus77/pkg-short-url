package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logger *Logger

type (
	Logger struct {
		File          *os.File
		InfoLogger    *log.Logger
		WarningLogger *log.Logger
		ErrorLogger   *log.Logger
		PanicLogger   *log.Logger
	}
)

func init() {
	var (
		err  error
		file *os.File
	)

	if _, err = os.Stat("./log"); os.IsNotExist(err) {
		if err = os.Mkdir("./log", 0644); err != nil {
			log.Panic(err)
		}
	}

	file, err = os.OpenFile(fmt.Sprintf("./log/%s.txt", time.Now().Format("01-02-2006")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger = &Logger{
		File:          file,
		InfoLogger:    log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		WarningLogger: log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger:   log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		PanicLogger:   log.New(file, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func INFO(text string) {
	logger.InfoLogger.Println(text)
}

func WARNING(text string) {
	logger.WarningLogger.Println(text)
}

func ERROR(text string) {
	logger.ErrorLogger.Println(text)
}

func PANIC(text string) {
	logger.PanicLogger.Panicln(text)
}
