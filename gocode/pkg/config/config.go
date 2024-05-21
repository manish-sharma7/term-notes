package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	NotesDir string
}

func getNotesDir() string {
	// TODO: Should give support to change notes dir to user other than the HOME Dir
	val := os.Getenv(homeEnv)
	if len(val) == 0 {
		fmt.Println("HOME Env is not present, using / path")
		val = "/"
	}
	return filepath.Join(val, dirName)
}

func GetConfig() Config {
	return Config{
		NotesDir: getNotesDir(),
	}
}
