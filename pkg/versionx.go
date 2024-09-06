package pkg

import (
	"errors"
	"os"
	"path/filepath"
)

func ListVersion() ([]string, error) {
	javaPath, err := ReadConfig()
	if err != nil {
		return nil, errors.New("java目录未配置，请先配置, 示例: javav config D:/programs/Java")
	}

	items := make([]string, 0)
	dirs, err := os.ReadDir(javaPath)
	if err != nil {
		return items, err
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		name := d.Name()
		_, err := os.Stat(filepath.Join(javaPath, name, "bin/java.exe"))
		if os.IsNotExist(err) {
			continue
		}
		items = append(items, name)
	}

	return items, nil
}
