package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/joho/godotenv"
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

var updateHotKeyTitle func(string)

func monitorFileModification(filepath string) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	modTime := fileInfo.ModTime()
	fmt.Println("Initial modification time:", modTime)

	// Calculate the end time after 30 minutes
	endTime := time.Now().Add(30 * time.Minute)

	for time.Now().Before(endTime) {
		fileInfo, err = os.Stat(filepath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if modTime != fileInfo.ModTime() {
			fmt.Println("Modification time changed:", fileInfo.ModTime())
			modTime = fileInfo.ModTime()
			godotenv.Load(filepath)
		}

		time.Sleep(time.Second * 3) // Sleep for 3 seconds before checking again
	}

	fmt.Println("Monitoring completed.")
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)

	mQuitOrig := systray.AddMenuItem(UText("Exit"), UText("Quit the whole app"))
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	mAbout := systray.AddMenuItem(UText("About"), UText("Open the project page"))
	go func() {
		for {
			<-mAbout.ClickedCh
			open.Start("https://github.com/linexjlin/inputGPT")
		}
	}()

	mSetKey := systray.AddMenuItem(UText("Set API KEY"), UText("Set the OpenAI KEY, baseurl etc.."))
	go func() {
		for {
			<-mSetKey.ClickedCh
			open.Start("env.txt")
			go monitorFileModification("env.txt")
		}
	}()

	mHotKey := systray.AddMenuItem("", UText("Click to active GPT"))
	updateHotKeyTitle = mHotKey.SetTitle

	systray.AddSeparator()

	mManager := systray.AddMenuItem(UText("Manage Prompts"), UText("Modify, Delete prompts"))
	go func() {
		for {
			<-mManager.ClickedCh
			open.Start("prompts")
		}
	}()

	mImport := systray.AddMenuItem(UText("Import"), UText("Import a prompt from clipboard"))

	mClearContext := systray.AddMenuItem(UText("Clear Context"), UText("Clear Context"))
	g_userCore.AddSetContextMenuFunc(mClearContext.SetTitle)

	go func() {
		for {
			select {
			case <-mClearContext.ClickedCh:
				fmt.Println("Clear Context")
				g_userCore.ClearContext()
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
		m := systray.AddMenuItemCheckbox(UText(fmt.Sprintf("%s", msk)), UText("Select this prompt"), false)
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
						g_userCore.initUserCore()
					} else {
						if p, e := loadModePrompt(filepath); e != nil {
							fmt.Println(e)
							continue
						} else {
							g_userCore.SetMask(mk)
							g_userCore.SetModePrompt(p)
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
					if p, e := parseModePrompt(clipboardContent); e != nil {
						fmt.Println(e)
					} else {
						if p.Name != "" {
							promptFilePath := fmt.Sprintf("prompts/%s.json", p.Name)
							if _, err := os.Stat(promptFilePath); err == nil {
								fmt.Println("File exists.", promptFilePath)
							} else {
								fmt.Println("create new", promptFilePath)
								m := systray.AddMenuItemCheckbox(fmt.Sprintf("%s", p.Name), "Check Me", false)
								maskCnt++
								idx := maskCnt
								go func() {
									for {
										<-m.ClickedCh
										g_userCore.SetMask(p.Name)
										g_userCore.SetModePrompt(p)

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
							saveModePrompt(p, promptFilePath)
						}
					}
				}
			}
		}
	}()

	updateHotKeyTitle(fmt.Sprintf(UText("Copy the question then click \"%s\" to query GPT"), strings.ToUpper(strings.Join(getGPTHotkeys(), "+"))))
}
