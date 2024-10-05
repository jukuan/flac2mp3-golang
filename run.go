package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func convertFlacToMp3(filePath string) {
	outputFile := strings.TrimSuffix(filePath, ".flac") + ".mp3"
	cmd := exec.Command("ffmpeg", "-i", filePath, outputFile)
	err := cmd.Run()                    
	if err != nil {
		fmt.Printf("Error converting %s: %v\n", filePath, err)
	}
}

func readFlacFilesAndConvert(sourceDir string) {
	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".flac") {
			fmt.Printf("Converting %s...\n", path)
			convertFlacToMp3(path)
		}       
		return nil
	})
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the source directory as the first argument.")
		return
	}
	sourceDir := os.Args[1]
	readFlacFilesAndConvert(sourceDir)
}