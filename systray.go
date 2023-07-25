package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func getMasks() (masks []string) {
	// Define the directory path and file extension
	dir := "prompts"
	ext := ".json"

	// Use the filepath package to get a list of all files with the specified extension
	files, err := filepath.Glob(filepath.Join(dir, "*"+ext))
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		filename := filepath.Base(file)
		filenameWithoutExt := filename[:len(filename)-len(ext)]
		fmt.Println(filenameWithoutExt)
		masks = append(masks, filenameWithoutExt)
	}
	return
}

func onExit() {
	now := time.Now()
	fmt.Println("exit", now)
}

// TODO: Set context cnt in clear context menu
func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	mQuitOrig := systray.AddMenuItem("Exit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	systray.AddSeparator()

	mClearContext := systray.AddMenuItem("Clear Context", "Clear Context")
	go func() {
		for {
			select {
			case <-mClearContext.ClickedCh:
				fmt.Println("Clear Context")
				g_userSetting.histMessages = g_userSetting.histMessages[:0]
			}
		}
	}()

	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("InputGPT")
	systray.SetTooltip("InputGPT a Helpful input Assistant")

	systray.AddSeparator()

	var maskMenus []*systray.MenuItem
	var masks = getMasks()

	masks = append(masks, "Default")

	for i, msk := range masks {
		fmt.Println(i)
		m := systray.AddMenuItemCheckbox(fmt.Sprintf("%s", msk), "Check Me", false)
		filepath := fmt.Sprintf("prompts/%s.json", msk)
		mk := msk
		if i == len(masks)-1 {
			m.Check()
			filepath = ""
		}
		var idx = i
		go func() {
			for {
				select {
				case <-m.ClickedCh:
					fmt.Println(idx, filepath, mk)
					if filepath == "" {
						initUserSetting()
					} else {
						if p, e := loadPrompt(filepath); e != nil {
							fmt.Println(e)
							continue
						} else {
							g_userSetting.headMessages = p.HeadMessages
							g_userSetting.model = p.Model
							if p.MaxContext != 0 {
								g_userSetting.maxConext = p.MaxContext
							} else {
								g_userSetting.maxConext = getMaxContext()
							}
						}
					}

					for ii, mm := range maskMenus {
						if ii == idx {
							mm.Check()
						} else {
							mm.Uncheck()
						}
					}

				}
			}
		}()
		maskMenus = append(maskMenus, m)
	}
}
