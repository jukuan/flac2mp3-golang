package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func isFFmpegInstalled() bool {
	cmd := exec.Command("ffmpeg", "-version")
	err := cmd.Run()
	return err == nil
}

func convertFlacToMp3(filePath string, bitrate string) {
	outputFile := strings.TrimSuffix(filePath, ".flac") + ".mp3"

	// Check if the output MP3 file already exists
	if _, err := os.Stat(outputFile); err == nil {
		fmt.Printf("Warning: %s already exists. Skipping conversion for %s\n", outputFile, filePath)
		return
	}

	cmd := exec.Command("ffmpeg", "-i", filePath, "-b:a", bitrate, outputFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error converting %s: %v\n", filePath, err)
	}
}

func readFlacFilesAndConvert(sourceDir string, bitrate string) {
	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".flac") {
			fmt.Printf("Converting %s with bitrate %s...\n", path, bitrate)
			convertFlacToMp3(path, bitrate)
		}
		return nil
	})
}

func main() {
	sourceDir := ""
	bitrate := "192k"

	if len(os.Args) > 1 {
		sourceDir = os.Args[1]
	}

	if len(os.Args) > 2 {
		bitrate = os.Args[2]
	}

	if sourceDir == "" {
		fmt.Println("Please provide the source directory as the first argument.")
		return
	}

	if !isFFmpegInstalled() {
		fmt.Println("ffmpeg is not installed. Please install ffmpeg to proceed with the conversion.")
		return
	}

	readFlacFilesAndConvert(sourceDir, bitrate)
}