package files

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func ReadFileLineByLine(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		logrus.Error(err)
	}

	scanner := bufio.NewScanner(file)
	var logsLines []string
	for scanner.Scan() {
		logsLines = append(logsLines, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		logrus.Fatal(err)
	}
	return logsLines
}

func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
