package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jeandeaual/go-locale"
)

//go:embed languages/*
var langDatas embed.FS

var MYLANG = "en"

type Language struct {
	Data map[string]map[string]string
}

func NewLanguage() *Language {
	var l = Language{Data: make(map[string]map[string]string)}
	l.SetLang()
	l.Load()
	return &l
}

func (l *Language) UText(langText string) string {
	translated := langText
	if lang, ok := l.Data[langText]; ok {
		if text, ok := lang[MYLANG]; ok {
			translated = text
		}
	} else {
		if MYLANG != "en" {
			fmt.Printf("unmatch:\"%s\" \"%s\"\n", langText, MYLANG)
		}
	}

	return translated
}

func (l *Language) UTextWithLangCode(langText, lcode string) string {
	translated := langText
	if lang, ok := l.Data[langText]; ok {
		if text, ok := lang[lcode]; ok {
			translated = text
		}
	} else {
		if lcode != "en" {
			fmt.Printf("unmatch:\"%s\" \"%s\"\n", langText, lcode)
			return ""
		}
	}

	return translated
}

func (l *Language) SetLang() {
	if lang := os.Getenv("MYLANG"); lang != "" {
		fmt.Println("using user lang:", lang)
		MYLANG = lang
		return
	}
	userLanguage, err := locale.GetLanguage()
	if err == nil {
		fmt.Println("Current Language:", userLanguage)
		MYLANG = userLanguage
	}
}

func (l *Language) Load() {
	files, err := langDatas.ReadDir("languages")
	if err != nil {
		log.Fatal(err)
	}

	languages := make([]map[string]map[string]string, 0)

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			fmt.Println("load language:", file.Name())
			data, err := langDatas.ReadFile("languages/" + file.Name())
			if err != nil {
				log.Println("Error reading file:", err)
				continue
			}

			languageJson := make(map[string]map[string]string)
			err = json.Unmarshal(data, &languageJson)
			if err != nil {
				log.Println("Error unmarshaling JSON:", err)
				continue
			}

			languages = append(languages, languageJson)
		}
	}

	l.combineLanguages(languages)
}

func (l *Language) combineLanguages(languages []map[string]map[string]string) {
	for _, language := range languages {
		for key, value := range language {
			if _, ok := l.Data[key]; !ok {
				l.Data[key] = make(map[string]string)
			}

			for k, v := range value {
				l.Data[key][k] = v
			}
		}
	}
}
