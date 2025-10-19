package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Logger struct {
	Id   string
	Path string
}

func New() *Logger {
	//main, _ := os.Getwd()
	//path := "/logs"
	//fullPath := filepath.Join(main, path)

	fullPath := filepath.Join("/logs")
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		os.MkdirAll(fullPath, 0755)
	}

	fmt.Printf("path: %s", fullPath)
	return &Logger{
		Id:   uuid.New().String(),
		Path: fullPath,
	}
}

func (l *Logger) Write(message string) error {

	path := fmt.Sprintf(l.Path + "/" + l.Id)

	fmt.Printf("\npath: %s\n", path)

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Sprintf("\nErro ao abrir arquivo. Erro: %s", err)
		return err
	}

	now := time.Now()

	timestamp := now.Format("2006-01-02 15:04:05")

	newLine := fmt.Sprintf(timestamp + ": " + message + "\n")

	_, err = file.Write([]byte(newLine))
	if err != nil {
		return err
	}

	return nil
}

func (l *Logger) Batch(message []string) error {

	path := fmt.Sprintf(l.Path + "/" + l.Id)

	fmt.Printf("\npath: %s\n", path)

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("\nErro ao abrir arquivo. Erro: %s", err)
		return err
	}

	now := time.Now()

	timestamp := now.Format("2006-01-02 15:04:05")

	for _, msg := range message {

		newLine := fmt.Sprintf(timestamp + ": " + msg + "\n")

		_, err = file.Write([]byte(newLine))
		if err != nil {
			return err
		}

	}

	return nil
}
