#!/bin/bash
xxd -i $(which ffmpeg) > ffmpeg.go
go build -o flac2mp3.out run.go
