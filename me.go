package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"bitbucket.org/kardianos/osext"
)

const (
	logTsFormat = "2006-01-02T15-04-05Z0700"
)

var (
	processImagePath = ""
	processImageDir  = ""
)

func init() {
}

func CwdToMe() error {
	exeDir, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}
	err = os.Chdir(exeDir)
	if err != nil {
		return err
	}
	return nil
}

func GetLogFilePathForMe() (string, error) {
	processImagePath, err := osext.Executable()
	if err != nil {
		return "", err
	}
	_, imageName := filepath.Split(processImagePath)

	imageName = strings.TrimSuffix(imageName, ".exe")

	logsDir, err := CreateLogsDirForMe()
	if err != nil {
		return "", err
	}

	logFileName := fmt.Sprintf("%s.%s.%d.log", imageName, time.Now().Format(logTsFormat), os.Getpid())

	return filepath.Join(logsDir, logFileName), nil
}

func CreateLogFileForMe() (*os.File, error) {
	logFilePath, err := GetLogFilePathForMe()
	if err != nil {
		return nil, err
	}
	logFile, err := os.Create(logFilePath)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func GetLogsDirForMe() (string, error) {
	processImageDir, err := osext.ExecutableFolder()
	if err != nil {
		return "", err
	}
	return filepath.Join(processImageDir, "logs"), nil
}

func CreateLogsDirForMe() (string, error) {
	logsDir, err := GetLogsDirForMe()
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(logsDir, 0)
	if err != nil {
		return "", err
	}

	return logsDir, nil
}
