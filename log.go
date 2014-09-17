package utils

import "github.com/Sirupsen/logrus"

type SimpleLogrusHook struct {
	OnFire func(*logrus.Entry)
}

func (this *SimpleLogrusHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (this *SimpleLogrusHook) Fire(entry *logrus.Entry) error {
	this.OnFire(entry)
	return nil
}
