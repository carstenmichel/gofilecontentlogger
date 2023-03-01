package main

import (
	"fmt"
	"io/fs"
	"os"
	"time"

	instana "github.com/instana/go-sensor"
	"github.com/instana/go-sensor/logger"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type directories struct {
	checkDirectory string
	logDirectory   string
}

func checkDirectory(dirs directories) {
	logger := log.WithFields(log.Fields{
		"check directory": dirs.checkDirectory,
		"log directory":   dirs.logDirectory,
	})
	logger.Infoln("Start checking Directory")
	files, err := os.ReadDir(dirs.checkDirectory)
	if err != nil {
		logger.Errorln("Failed to read Directory")
	}
	for _, file := range files {
		log.WithFields(log.Fields{"filename": file.Name()}).Infoln("found following File")
	}
	logFilesToLogDir(dirs, files)
}

func logFilesToLogDir(dirs directories, files []fs.DirEntry) {
	f, err := os.OpenFile(dirs.logDirectory, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.WithField("logDir", dirs.logDirectory).Errorln("Could not open logfile for writing")
	}
	defer f.Close()
	for _, file := range files {
		f.WriteString(fmt.Sprintf("%s was found\n", file.Name()))
	}
}

func setupEnv() {
	viper.SetEnvPrefix("PREFIX")
	viper.BindEnv("checkdir")
	viper.BindEnv("logdir")
	viper.SetDefault("checkdir", ".")
	viper.SetDefault("logdir", "./logfile.txt")
}

/*
A small application to test volume mounts on Openshift.
Will scan one directory,
write some more or less usefull data to another directory
and will log to the console.
*/
func main() {
	instanalogger := logger.New(log.New())
	instana.SetLogger(instanalogger)
	setupEnv()
	dirs := directories{
		checkDirectory: viper.GetString("checkdir"),
		logDirectory:   viper.GetString("logdir"),
	}
	instana.InitSensor(&instana.Options{
		Service:           "gofilelogger",
		LogLevel:          instana.Debug,
		EnableAutoProfile: true})

	log.Warn("Hello")
	for {
		go checkDirectory(dirs)
		time.Sleep(time.Second * 5)
	}

}
