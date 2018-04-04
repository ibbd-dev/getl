package common

import (
	"encoding/json"
)

type PipeData map[string]interface{}

// 从stdio中读取数据
func (d PipeData) Read() {

	bts, err := reader.ReadString('\n')
	if err == io.EOF {
		break
	} else if err != nil {
		panic(err)
	}

}

// 写入数据到stdout
func (d PipeData) Write() {

}
