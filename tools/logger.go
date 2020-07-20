package tools

import (
	config2 "Web_Api/tools/config"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func InitLogger()  {
	log.SetFormatter(&log.TextFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime: "@timestamp",
			log.FieldKeyLevel: "@level",
			log.FieldKeyMsg: "@message",
		},
	})

	switch Mode(config2.ApplicationConfig.Mode) {
	case ModeDev, ModeTest:
		log.SetOutput(os.Stdout)
		log.SetLevel(log.TraceLevel)
	case ModeProd:
		file, err := os.OpenFile(
			config2.LogConfig.Dir+"/api-"+time.Now().Format("2006-01-02")+".log",
			os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC,
			0600)
		if err != nil {
			log.Fatalln("log init failed")
		}
		var info os.FileInfo
		info, err = file.Stat()
		if err !=nil {
			log.Fatal(err)
		}
		fileWriter := logFileWriter{file: file,size: info.Size()}
		log.SetOutput(&fileWriter)
		log.SetLevel(log.WarnLevel)
	}
	log.SetReportCaller(true)
}

type logFileWriter struct {
	file *os.File
	size int64
}

func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}
	n, e := p.file.Write(data)
	p.size += int64(n)
	//每天一个文件
	if p.file.Name() != config2.LogConfig.Dir+"/api-"+time.Now().Format("2006-01-02")+".log" {
		p.file.Close()
		p.file, _ = os.OpenFile(
			config2.LogConfig.Dir+"/api-"+time.Now().Format("2006-01-02")+".log",
			os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC,
			0600)
		p.size = 0
	}
	return n, e
}