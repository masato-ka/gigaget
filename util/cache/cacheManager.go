package cache

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io"
	"os"
	"path/filepath"
	"time"
)

type CacheManager struct {
	atCreate time.Time
	lines    []string
}

func (cm *CacheManager) Load() error {
	filepath, err := cm.getCacheFilePath()
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0644)

	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	reader := bufio.NewReaderSize(file, 4096)
	start, _, err := reader.ReadLine()
	cm.atCreate, _ = time.Parse("2019-05-23 21:58:56", string(start))

	for {

		line, _, err := reader.ReadLine()
		if string(line) != "" {
			cm.lines = append(cm.lines, string(line))
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

	}

	return nil
}

func (cm *CacheManager) Save() error {
	now := time.Now()
	filepath, err := cm.getCacheFilePath()
	file, err := os.Create(filepath)

	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, now.String()+"\n")

	for _, line := range cm.lines {
		fmt.Fprintf(file, line+"\n")
	}

	return nil
}

func (cm *CacheManager) GetLastModify() string {
	return cm.atCreate.String()
}

func (cm *CacheManager) Append(line string) {
	cm.lines = append(cm.lines, line)
}

func (cm *CacheManager) getCacheFilePath() (string, error) {
	root, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	cachedir := filepath.FromSlash(root + "/.gigaget")
	if _, err := os.Stat(cachedir); os.IsNotExist(err) {
		err := os.Mkdir(cachedir, 0666)
		if err != nil {
			return "", err
		}
	}
	cachefile := filepath.FromSlash(cachedir + "/cache")

	return cachefile, nil
}
