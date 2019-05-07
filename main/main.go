package main

import (
	"github.com/sirupsen/logrus"
	"github.com/zbindenren/logrus_mail"
	"os"
	"time"
)

func logTest() {
	var log = logrus.New()

	//输入到文件
	file , err := os.OpenFile("acess.log",os.O_CREATE|os.O_WRONLY,0666)

	if err == nil {
		//default ,  log.Out = os.Stdout
		log.Out = file
	} else {
		log.Info("Failed to log to file , using default stderr")
	}
	//Fields 里面的字段是我们自己定义的，Info 才是等级
	log.WithFields(logrus.Fields{
		"animal":"walrus",
		"size":10,
		"BB":"allal",
	}).Info("hello log !!!!")

	log.WithFields(logrus.Fields{
		"AA":"bilibili",
	}).Warn("this is a warn!!!")
}

func Demo() {
	//生成一个logrus对象
	logger := logrus.New()

	//定义Email hook
	hook , err := logrus_mail.NewMailAuthHook(
		"logrus_mail",
		"smtp.exmail.qq.com",
		465,
		"xxxxxxxxxxxxx",
		"xxxxxxxxxx",
		"xxxxxxxxxxxxxxx",
		"xxxxxxxxxxxxxxx",
	)

	//如果注册hook
	if err == nil {
		logger.Hooks.Add(hook)
	}

	//生成*Entry
	var fileName = "access.log"

	contextLogger := logger.WithFields(logrus.Fields{
		"file":	fileName,
		"conteng": "ha ha ha ~~~",
	})

	//设置时间戳
	contextLogger.Time = time.Now()
	contextLogger.Message = "这是一个hook发送来的邮件"
	//只发送Error, Fatal , Panic 级别的log
	contextLogger.Level = logrus.ErrorLevel

	//使用Fire发送 ， 包含时间戳， message
	err2 := hook.Fire(contextLogger)
	print(err2)
	print("sss")
}


func main() {
	//ioapply.Scanf()
	//ioapply.Scanln()
	//ioapply.Scanf()
	//ioapply.Sscanf()
	//ioapply.Fscanf()
	//ioapply.FlagBool()
	//ioapply.FlagVar()
	//ioapply.FlagArgs()
	//logTest()
	//print("dasdfaf")
	//Demo()
}




