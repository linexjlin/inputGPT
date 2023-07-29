//go:build darwin
// +build darwin

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func OSDepCheck() {
	// change working directory
	userDir, _ := os.UserConfigDir()
	APP_SUPPORT_DIR := filepath.Join(userDir, "InputGPT")
	os.Chdir(APP_SUPPORT_DIR)
	promptsDir := filepath.Join(APP_SUPPORT_DIR, "prompts")

	if _, err := os.Stat(promptsDir); os.IsNotExist(err) {
		err := os.MkdirAll(promptsDir, 0755)
		if err != nil {
			fmt.Printf("unable to create：%s\n", err)
		}
		fmt.Println("promptsDir created")
	}

	envFile := filepath.Join(APP_SUPPORT_DIR, "env.txt")
	if _, err := os.Stat(envFile); os.IsNotExist(err) {
		exePath, err := os.Executable()
		if err == nil {
			appPath := filepath.Dir(filepath.Dir(exePath))
			sourceEnvFile := filepath.Join(appPath, "Resources", "env.txt")
			copyFile(sourceEnvFile, envFile)
			godotenv.Load(envFile)
		}
	}

	godotenv.Load("env.txt")
	return
}

// 复制文件
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	err = dstFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
