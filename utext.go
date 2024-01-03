package main

import "fmt"

// give the right lanuage with query text when no match return the langText itself
func UText(langText string) string {
	translated := langText
	if lang, ok := g_languages.Data[langText]; ok {
		if text, ok := lang[MYLANG]; ok {
			translated = text
		}
	} else {
		if MYLANG != "en" {
			fmt.Printf("unmatch:\"%s\" \"%s\"\n", langText, MYLANG)
		}
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
