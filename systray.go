package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/atotto/clipboard"
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
		for {
			<-mAbout.ClickedCh
			open.Start("https://github.com/linexjlin/inputGPT")
		}
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

	mManager := systray.AddMenuItem(UText("Manager Prompts"), UText("Manager the prompts"))
	go func() {
		for {
			<-mManager.ClickedCh
			open.Start("prompts")
		}
	}()

	mImport := systray.AddMenuItem(UText("Import"), UText("Import"))

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
	//	systray.SetTitle(UText("InputGPT"))
	systray.SetTooltip(UText("InputGPT a Helpful input Assistant"))

	systray.AddSeparator()

	var maskMenus []*systray.MenuItem
	masks := getMasks()

	masks = append(masks, UText("Default"))
	maskCnt := 0

	for i, msk := range masks {
		m := systray.AddMenuItemCheckbox(fmt.Sprintf("%s", msk), "Check Me", false)
		filepath := fmt.Sprintf("prompts/%s.json", msk)
		mk := msk
		if i == len(masks)-1 {
			m.Check()
			filepath = ""
		}
		idx := i
		maskCnt = i
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
							initUserSetting() // reset all user settings
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

	// Handling import events
	go func() {
		for {
			select {
			case <-mImport.ClickedCh:
				fmt.Println(UText("Import"))
				clipboardContent, err := clipboard.ReadAll()
				if err != nil {
					fmt.Println("Failed to read clipboard content:", err)
				} else {
					fmt.Println(clipboardContent)
					if p, e := parsePrompt(clipboardContent); e != nil {
						fmt.Println(e)
					} else {
						if p.Name != "" {
							if e = savePrompt(p, fmt.Sprintf("prompts/%s.json", p.Name)); e == nil {
								m := systray.AddMenuItemCheckbox(fmt.Sprintf("%s", p.Name), "Check Me", false)
								maskCnt++
								idx := maskCnt
								go func() {
									for {
										<-m.ClickedCh
										initUserSetting() // reset all user settings
										g_userSetting.headMessages = p.HeadMessages
										if p.Model != "" {
											g_userSetting.model = p.Model
										}

										if p.MaxContext != 0 {
											g_userSetting.maxConext = p.MaxContext
										}
										updateClearContextTitle(0)
										g_userSetting.mask = p.Name

										for ii, mm := range maskMenus {
											if ii == idx {
												mm.Check()
											} else {
												mm.Uncheck()
											}
										}
									}
								}()
								maskMenus = append(maskMenus, m)
							}
						}
					}
				}
			}
		}
	}()

	updateClearContextTitle(0)
	updateHotKeyTitle(fmt.Sprintf(UText("Copy the question then click \"%s\" to query GPT"), strings.ToUpper(strings.Join(getGPTHotkeys(), "+"))))
}
