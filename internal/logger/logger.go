package logger

import (
	"io"
	"log"
	"os"
)

func Init(logFile string) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("打开日志文件失败：", err)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
