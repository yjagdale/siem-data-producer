package files

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/utils/networkUtils"
	"math/rand"
	"net"
	"os"
)

func ReadFileLineByLine(filePath string) []string {
	file := readFile(filePath)
	scanner := bufio.NewScanner(file)
	var logsLines []string
	for scanner.Scan() {
		logsLines = append(logsLines, scanner.Text())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logrus.Error(err.Error())
		}
	}(file)
	return logsLines
}

func ReadAndPublishInChunk(filePath string, connection net.Conn) {
	file := readFile(filePath)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logrus.Error("Error while closing file", err.Error())
		}
	}(file)
	scanner := bufio.NewScanner(file)
	min := 1
	max := 10
	maxGoroutines := rand.Intn(max-min) + min
	guard := make(chan struct{}, maxGoroutines)

	for scanner.Scan() {
		guard <- struct{}{}
		go func(conn net.Conn, data string) {
			networkUtils.ProduceLog(conn, data)
			<-guard
		}(connection, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		logrus.Errorln(err)
	}
}

func readFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		logrus.Error(err)
	}
	return file
}
