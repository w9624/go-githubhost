package runner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	GithubPath = "./domain/github.txt"

	MaxSize = 1024 * 1024 * 3
)

func LoadFile(path string) (domains []string, err error) {
	if path == "" {
		path = GithubPath
	}

	fr, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("打开文件异常：%s", err.Error())
	}

	br := bufio.NewReaderSize(fr, MaxSize)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		domains = append(domains, strings.TrimSpace(string(line)))
	}

	return
}
