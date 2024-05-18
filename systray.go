package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/joho/godotenv"
	icons "github.com/linexjlin/systray-icons/dynamic/ball-triangle"
	icon "github.com/linexjlin/systray-icons/enter-the-keyboard"
	"github.com/skratchdot/open-golang/open"
)

type SysTray struct {
	userCore          *UserCore
	updateHotKeyTitle func(string)
}

func (st *SysTray) Run() {
	systray.Run(st.onReady, st.onExit)
}

func (st *SysTray) ShowRunningIcon(ctx context.Context, done <-chan struct{}) {
	defer systray.SetIcon(icon.Data)
	i := 0
	for {
		select {
		case <-ctx.Done():
			return
		case <-done:
			return
		default:
			systray.SetIcon(icons.Datas[i])
			time.Sleep(time.Millisecond * 120)
			i++
			if i == len(icons.Datas)-1 {
				i = 0
			}
		}
	}
}

func (st *SysTray) getMasks() (masks []string) {
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

func (st *SysTray) onExit() {
	now := time.Now()
	fmt.Println("exit", now)
}

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

func (st *SysTray) onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)

	mQuitOrig := systray.AddMenuItem(UMenuText("Exit"), UMenuText("Quit the whole app"))
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	mHotKey := systray.AddMenuItem("", UMenuText("Click to active GPT"))
	st.updateHotKeyTitle = mHotKey.SetTitle

	mAbout := systray.AddMenuItem(UMenuText("About"), UMenuText("Open the project page"))
	go func() {
		for {
			<-mAbout.ClickedCh
			open.Start("https://github.com/linexjlin/inputGPT")
		}
	}()

	systray.AddSeparator()

	mSetting := systray.AddMenuItem(UMenuText("Setting"), UMenuText(""))

	mSetKey := mSetting.AddSubMenuItem(UMenuText("Set API KEY"), UMenuText("Set the OpenAI KEY, baseurl etc.."))
	go func() {
		for {
			<-mSetKey.ClickedCh
			open.Start("env.txt")
			go monitorFileModification("env.txt")
		}
	}()

	mManager := mSetting.AddSubMenuItem(UMenuText("Manage Prompts"), UMenuText("Modify, Delete prompts"))
	go func() {
		for {
			<-mManager.ClickedCh
			open.Start("prompts")
		}
	}()

	mImport := mSetting.AddSubMenuItem(UMenuText("Import"), UMenuText("Import a prompt from clipboard"))

	systray.AddSeparator()

	var modesMenus []*systray.MenuItem
	modes := getModeList()
	for i, mode := range modes {
		m := systray.AddMenuItem(UMenuText(mode), UMenuText("Chose "+mode))
		go func(i int, mode string) {
			for {
				<-m.ClickedCh
				fmt.Println(i, mode)
				st.userCore.SetDefaultMode(mode)

				for ii, mm := range modesMenus {
					if ii == i {
						mm.Check()
					} else {
						mm.Uncheck()
					}
				}
			}
		}(i, mode)
		if i == 0 {
			m.Check()
		}
		modesMenus = append(modesMenus, m)
	}

	systray.AddSeparator()
	mClearContext := systray.AddMenuItem(UMenuText("Clear Context"), UMenuText("Clear Context"))
	st.userCore.AddSetContextMenuFunc(mClearContext.SetTitle)

	go func() {
		for {
			select {
			case <-mClearContext.ClickedCh:
				fmt.Println("Clear Context")
				st.userCore.ClearContext()
			}
		}
	}()

	systray.SetTemplateIcon(icon.Data, icon.Data)
	//	systray.SetTitle(UMenuText("InputGPT"))
	systray.SetTooltip(UMenuText("InputGPT a Helpful input Assistant"))

	systray.AddSeparator()

	var maskMenus []*systray.MenuItem
	masks := st.getMasks()

	masks = append(masks, UMenuText("Default"))
	maskCnt := 0

	for i, msk := range masks {
		m := systray.AddMenuItemCheckbox(UMenuText(fmt.Sprintf("%s", msk)), UMenuText("Select this prompt"), false)
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
						st.userCore.initUserCore()
					} else {
						if p, e := loadModePrompt(filepath); e != nil {
							fmt.Println(e)
							continue
						} else {
							st.userCore.SetMask(mk)
							st.userCore.SetModePrompt(p)
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
										st.userCore.SetMask(p.Name)
										st.userCore.SetModePrompt(p)

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

	st.updateHotKeyTitle(fmt.Sprintf(UMenuText("Copy the question then click \"%s\" to query GPT"), strings.ToUpper(strings.Join(getGPTHotkeys(), "+"))))
}
