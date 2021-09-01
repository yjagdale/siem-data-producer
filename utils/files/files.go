package files

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"github.com/yjagdale/siem-data-producer/utils/networkUtils"
	"net"
	"os"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ReadFileLineByLine(filePath string) []string {
	file := readFile(filePath)
	scanner := bufio.NewScanner(file)
	var logsLines []string
	for scanner.Scan() {
		logsLines = append(logsLines, scanner.Text())
	}
	defer file.Close()
	return logsLines
}

func ReadAndPublishInChunk(filePath string, connection net.Conn) {
	file := readFile(filePath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		networkUtils.ProduceLog(connection, scanner.Text())
	}
}

func readFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		logrus.Error(err)
	}
	return file
}
