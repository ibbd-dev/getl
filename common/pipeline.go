package common

import (
	"bufio"
	"encoding/json"
	"os"
)

type PipelineIO struct {
	Reader *bufio.Reader
	Writer *bufio.Writer
}

func NewPipelineIO() *PipelineIO {
	return &PipelineIO{
		Reader: bufio.NewReader(os.Stdin),
		Writer: bufio.NewWriter(os.Stdout),
	}
}

// 从stdio中读取数据
func (p *PipelineIO) Read() (row map[string]interface{}, err error) {
	bts, err := p.Reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bts, row)
	return row, err
}

// 写入数据到stdout
func (p *PipelineIO) Write(row map[string]interface{}) (err error) {
	bts, err := json.Marshal(row)
	if err != nil {
		return err
	}

	_, err = p.Writer.Write(bts)
	return err
}

// 写入数据到stdout
func (p *PipelineIO) WriteMapString(row map[string]string) (err error) {
	bts, err := json.Marshal(row)
	if err != nil {
		return err
	}

	_, err = p.Writer.Write(bts)
	return err
}
