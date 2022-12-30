package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/djimenez/iconv-go"
)

type SubtitleConverter struct {
	converter *iconv.Converter
}

func NewSubtitleConverter() *SubtitleConverter {
	converter, err := iconv.NewConverter("iso-8859-9", "utf-8")
	if err != nil {
		panic(err)
	}
	return &SubtitleConverter{converter: converter}
}

func main() {
	converter := NewSubtitleConverter()
	defer converter.Close()

	subtitlesPath := flag.String("path", ".", "path that contain subtitles")
	flag.Parse()
	if subtitlesPath == nil || *subtitlesPath == "" {
		panic("path is missing")
	}

	err := filepath.Walk(*subtitlesPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || filepath.Ext(path) != ".srt" {
				return nil
			}

			return converter.Covert(path)
		})

	if err != nil {
		panic(err)
	}
	log.Println("All subtitles fixed!")
}

func (sc *SubtitleConverter) Covert(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	inputBuffer, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	outputBuffer := make([]byte, len(inputBuffer)*2)
	_, _, err = sc.converter.Convert(inputBuffer, outputBuffer)
	if err != nil {
		return err
	}

	cleanedOutputBuffer := bytes.Trim(outputBuffer, "\x00")
	_, err = file.WriteAt(cleanedOutputBuffer, 0)
	return err
}

func (sc *SubtitleConverter) Close() {
	_ = sc.converter.Close()
}
