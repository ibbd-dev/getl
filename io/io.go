package io

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func (f *IOParams) ReadAll() ([]byte, error) {
	if !terminal.IsTerminal(0) {
		return ioutil.ReadAll(os.Stdin)
	}

	if f.Input == "" {
		return nil, errors.New("输入文件名为空")
	}

	return ioutil.ReadFile(f.Input)
}

func (f *IOParams) Write(data []byte) error {
	if f.Output == "" {
		fmt.Println(string(data))
		return nil
	}

	return ioutil.WriteFile(f.Output, data, 0644)
}
