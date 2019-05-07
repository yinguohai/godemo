package logapplay

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Example1() {
	var log = logrus.New()
	//以JSON格式存放日志信息
	log.Formatter = new(logrus.JSONFormatter)


	//log.Formatter = new(logrus.TextFormatter)
	//剔除颜色显示
	//log.Formatter.(*logrus.TextFormatter).DisableColors = true
	//显示时间
	//log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true

	log.Level = logrus.InfoLevel

	file , err := os.OpenFile("access.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)

	if err == nil {
		log.Out = file
	} else {
		panic("open file error !")
	}

	defer func() {
		err := recover()
		if err != nil {
			entry := err.(*logrus.Entry)
			log.WithFields(logrus.Fields{
				"omg": true,
				"err_animal":entry.Data["animal"],
				"err_size" : entry.Data["size"],
				"err_message": entry.Message,
				"err_level": entry.Level,
			}).Error("The ice breaks")
		}
	}()


	log.WithFields(logrus.Fields{
		"animal":"walrus",
		"number":0,
	}).Trace("Went to the beach")

	log.WithFields(logrus.Fields{
		"animal":"dog",
		"number":8,
	}).Info("A group of dog from the ocean")
}

///////////////////////////////////////Hook Test////////////////////////////////////////////////////
var (
	mystring string
)

type GlobalHook struct {}

/**
	Levels 是 Hook 的前置条件，只有在Levels中定义的等级才能触发Hook
 */
func (h *GlobalHook) Levels() []logrus.Level {
	var levels = []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
	}

	return levels

	//logrus.AllLevels 监控所有事件
	//return logrus.AllLevels
}

/**
	Fire 是Hook的执行体，用于定义hook到后需要进行的哪些操作
 */
func (h *GlobalHook) Fire(e *logrus.Entry) error {
	e.Data["mystr"] = mystring
	e.Data["num"] = "007"
	return nil
}

func ExampleHook() {
	l := logrus.New()
	l.Out = os.Stdout

	l.AddHook(&GlobalHook{})

	mystring = "first Value"
	l.Info("first log")

	//只有Error级别的事件才能被hook到
	mystring = "another value"
	l.Error("second log")
}