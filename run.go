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

func convertToMp3(inputFile string, outputFile string, bitrate string) {
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-b:a", bitrate, outputFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error converting %s: %v\n", inputFile, err)
	}
}

func convertAudioFiles(sourceDir string, bitrate string, inputFormat string) {
	filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), "."+inputFormat) {
			outputFile := strings.TrimSuffix(path, "."+inputFormat) + ".mp3"
			if _, err := os.Stat(outputFile); err == nil {
				fmt.Printf("Warning: %s already exists. Skipping conversion for %s\n", outputFile, path)
				return nil
			}
			fmt.Printf("Converting %s to MP3 with bitrate %s...\n", path, bitrate)
			convertToMp3(path, outputFile, bitrate)
		}
		return nil
	})
}


func removeAudioFiles(dirPath string) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (strings.HasSuffix(strings.ToLower(path), ".flac") || strings.HasSuffix(strings.ToLower(path), ".wav")) {
			err := os.Remove(path)
			if err != nil {
				fmt.Printf("Error removing %s: %v\n", path, err)
			} else {
				fmt.Printf("Removed %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through directory: %v\n", err)
	}
}

func main() {
	sourceDir := ""
	bitrate := "192k"
	inputFormat := "flac" // Default input format

	if len(os.Args) > 1 {
		sourceDir = os.Args[1]
	}

	if len(os.Args) > 2 {
		bitrate = os.Args[2]
	}

    if ("rm" == bitrate) {
        removeAudioFiles(sourceDir)
        return
    }

	if len(os.Args) > 3 {
		inputFormat = os.Args[3]
	}

	if sourceDir == "" {
		fmt.Println("Please provide the source directory as the first argument.")
		return
	}

	if !isFFmpegInstalled() {
		fmt.Println("ffmpeg is not installed. Please install ffmpeg to proceed with the conversion.")
		return
	}

	convertAudioFiles(sourceDir, bitrate, inputFormat)
}