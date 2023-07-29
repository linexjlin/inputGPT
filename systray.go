package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/getlantern/systray"
	icon "github.com/linexjlin/systray-icons/enter-the-keyboard"
	"github.com/skratchdot/open-golang/open"
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

var _mClearContextSetTitle func(string)

func updateClearContextTitle(n int) {
	_mClearContextSetTitle(fmt.Sprintf(UText("Clear Context %d/%d"), n, g_userSetting.maxConext))
}

var updateHotKeyTitle func(string)

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	mAbout := systray.AddMenuItem(UText("About"), UText("About the App"))
	go func() {
		<-mAbout.ClickedCh
		open.Start("https://github.com/linexjlin/inputGPT")
	}()
	mQuitOrig := systray.AddMenuItem(UText("Exit"), UText("Quit the whole app"))
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	mHotKey := systray.AddMenuItem("", UText("Click to active GPT"))
	updateHotKeyTitle = mHotKey.SetTitle

	systray.AddSeparator()

	mClearContext := systray.AddMenuItem(UText("Clear Context"), UText("Clear Context"))
	_mClearContextSetTitle = mClearContext.SetTitle
	go func() {
		for {
			select {
			case <-mClearContext.ClickedCh:
				fmt.Println("Clear Context")
				g_userSetting.histMessages = g_userSetting.histMessages[:0]
				updateClearContextTitle(0)
			}
		}
	}()

	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle(UText("InputGPT"))
	systray.SetTooltip(UText("InputGPT a Helpful input Assistant"))

	systray.AddSeparator()

	var maskMenus []*systray.MenuItem
	var masks = getMasks()

	masks = append(masks, UText("Default"))

	for i, msk := range masks {
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
							initUserSetting() //reset all user settings
							g_userSetting.headMessages = p.HeadMessages
							if p.Model != "" {
								g_userSetting.model = p.Model
							}

							if p.MaxContext != 0 {
								g_userSetting.maxConext = p.MaxContext
							}
							updateClearContextTitle(0)
							g_userSetting.mask = mk
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
	updateClearContextTitle(0)
	updateHotKeyTitle(fmt.Sprintf(UText("Copy the question then click \"%s\" to query GPT"), strings.ToUpper(strings.Join(getGPTHotkeys(), "+"))))
}
