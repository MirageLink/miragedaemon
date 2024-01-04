package dbspm

import (
	"bufio"
	"embed"
	"fmt"
	"gopkg.in/toast.v1"
	"math/rand"
	"time"
)

//go:embed motd.txt
var content embed.FS

func Deploy() {
	for {
		guh, _ := getRandomLine()
		notification := toast.Notification{
			AppID: "mirage <3",
			Title: guh,
		}
		notification.Push()
		time.Sleep(time.Hour / 2)
	}
}

func getRandomLine() (string, error) {
	source := rand.NewSource(time.Now().UnixNano())

	randomGenerator := rand.New(source)

	file, err := content.Open("motd.txt")
	if err != nil {
		return "", fmt.Errorf("error opening embedded file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading embedded file: %v", err)
	}

	if len(lines) == 0 {
		return "", fmt.Errorf("the embedded file is empty")
	}

	randomIndex := randomGenerator.Intn(len(lines))
	return lines[randomIndex], nil
}
