# Audio Converter (Go)

This is a simple audio converter tool written in Go that converts FLAC and WAV audio files to MP3 format using ffmpeg.

Features

- Converts FLAC and WAV audio files to MP3 format
- Supports custom bitrate for MP3 conversion
- Skips conversion if the output MP3 file already exists
- Ability to specify input format (FLAC or WAV)

Usage

1. Make sure you have ffmpeg installed on your system.
2. Clone the repository and navigate to the project directory.
3 Run the program with the following command:

    go run main.go <source_directory> <bitrate> <input_format>

        <source_directory>: The directory containing the audio files to convert.
        <bitrate>: Optional. The bitrate for the MP3 output (default is 192k).
        <input_format>: Optional. The input audio format (default is "flac").

Example:
```
go run main.go /path/to/audio/files 256k wav
```

Notes

- This is my first GoLang project, created as an exercise to learn the language and work with audio conversion.
- Please ensure that you have the necessary permissions to run the ffmpeg command for audio conversion.

Feel free to provide feedback or suggestions for improvement!