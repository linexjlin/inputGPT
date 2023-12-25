package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"github.com/jeandeaual/go-locale"
)

//go:embed languages/*
var langDatas embed.FS

var LANG = "en"

type Language struct {
	Data map[string]map[string]string
}

var g_languages = Language{Data: make(map[string]map[string]string)}

// give the right lanuage with query text when no match return the langText itself
func UText(langText string) string {
	translated := ""
	if lang, ok := g_languages.Data[langText]; ok {
		if text, ok := lang[LANG]; ok {
			translated = text
		}
	} else {
		if LANG != "en" {
			fmt.Printf("unmatch:\"%s\" \"%s\"\n", langText, LANG)
		}
		translated = langText
	}

	emoji := ""
	if lang, ok := g_languages.Data[langText]; ok {
		if text, ok := lang["emoji"]; ok {
			emoji = text
		}
	}
	if emoji != "" {
		translated = fmt.Sprintf("%s %s", translated, emoji)
	}
	return translated
}

func setLang() {
	userLanguage, err := locale.GetLanguage()
	if err == nil {
		fmt.Println("Current Language:", userLanguage)
		LANG = userLanguage
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
