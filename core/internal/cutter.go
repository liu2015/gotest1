package internal

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Cutter struct {
	level    string
	format   string
	Director string
	file     *os.File
	mutex    *sync.RWMutex
}

type CutterOption func(*Cutter)

func WithCutterFormat(format string) CutterOption {
	return func(c *Cutter) {
		c.format = format
	}
}

func NewCutter(director string, level string, options ...CutterOption) *Cutter {
	rotate := &Cutter{
		level:    level,
		Director: director,
		mutex:    new(sync.RWMutex),
	}
	for i := 0; i < len(options); i++ {
		options[i](rotate)
	}
	return rotate
}

func (c *Cutter) Write(bytes []byte) (n int, err error) {
	c.mutex.Lock()
	defer func() {
		if c.file != nil {
			_ = c.file.Close()
			c.file = nil
		}
		c.mutex.Unlock()
	}()
	var business string
	if strings.Contains(string(bytes), "business") {
		var compile *regexp.Regexp
		compile, err = regexp.Compile(`{"business":"([^,]+)"}`)
		if err != nil {
			return 0, err
		}
		if compile.Match(bytes) {
			finds := compile.FindSubmatch(bytes)
			business = string(finds[len(finds)-1])
			bytes = compile.ReplaceAll(bytes, []byte(""))
		}
		compile, err = regexp.Compile(`"business": "([^,]+)"`)
		if err != nil {
			return 0, err
		}
		if compile.Match(bytes) {
			filnds := compile.FindSubmatch(bytes)
			business = string(filnds[len(filnds)-1])
			bytes = compile.ReplaceAll(bytes, []byte(""))
		}
	}

	format := time.Now().Format(c.format)
	formats := make([]string, 0, 4)
	formats = append(formats, c.Director)
	if format != "" {
		formats = append(formats, format)
	}
	if business != "" {
		formats = append(formats, business)
	}
	formats = append(formats, c.level+".log")
	filename := filepath.Join(formats...)
	dirname := filepath.Dir(filename)
	err = os.MkdirAll(dirname, 0755)

	if err != nil {
		return 0, err
	}
	c.file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		return 0, err
	}
	return c.file.Write(bytes)
}
