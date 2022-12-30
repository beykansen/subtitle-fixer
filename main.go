package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"unicode/utf8"

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

	count := 0
	err := filepath.Walk(*subtitlesPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || filepath.Ext(path) != ".srt" {
				return nil
			}
			fixed, err := converter.Covert(path)
			if fixed {
				count++
			}
			return err
		})

	if err != nil {
		panic(err)
	}
	log.Printf("%d Subtitles Fixed!", count)
}

func (sc *SubtitleConverter) Covert(filename string) (bool, error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = file.Close()
	}()
	inputBuffer, err := io.ReadAll(file)
	if err != nil {
		return false, err
	}

	alreadyFixed := utf8.Valid(inputBuffer)
	if alreadyFixed {
		return false, nil
	}

	outputBuffer := make([]byte, len(inputBuffer)*2)
	_, _, err = sc.converter.Convert(inputBuffer, outputBuffer)
	if err != nil {
		return false, err
	}

	cleanedOutputBuffer := bytes.Trim(outputBuffer, "\x00")
	_, err = file.WriteAt(cleanedOutputBuffer, 0)
	if err != nil {
		return false, err
	}
	log.Printf("%s fixed.\n", filename)
	return true, nil
}

func (sc *SubtitleConverter) Close() {
	_ = sc.converter.Close()
}
