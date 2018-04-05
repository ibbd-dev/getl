// 参考Python的DictReader
package csv

import (
	"encoding/csv"
	"fmt"
	"io"
)

// A Reader reads records from a CSV-encoded file.
type MapReader struct {
	Reader     *csv.Reader
	fieldnames []string
}

// NewMapReader returns a new Reader that reads from r.
func NewMapReader(r io.Reader) *MapReader {
	return &MapReader{
		Reader: csv.NewReader(r),
	}
}

func (r *MapReader) SetFieldnames(fieldnames []string) {
	r.fieldnames = fieldnames
}

func (r *MapReader) Read() (record map[string]string, err error) {
	if len(r.fieldnames) == 0 {
		// 如果没有设置字段名，则默认为csv的第一行为字段名
		r.fieldnames, err = r.Reader.Read()
		if err != nil {
			return nil, err
		}
	}

	var rawRecord []string
	rawRecord, err = r.Reader.Read()
	if err != nil {
		return nil, err
	}

	length := len(r.fieldnames)
	record = make(map[string]string)
	for index := 0; index < length; index++ {
		field := r.fieldnames[index]
		if _, exists := record[field]; exists {
			return nil, fmt.Errorf("Multiple indices with the same name '%s'", field)
		}
		record[field] = rawRecord[index]
	}
	return record, err
}

// ReadAll reads all the remaining records from r. Each record is a map of column name to field value.
func (r *MapReader) ReadAll() (records []map[string]string, err error) {
	var record map[string]string
	for record, err = r.Read(); err == nil; record, err = r.Read() {
		records = append(records, record)
	}
	if err != nil && err != io.EOF {
		return nil, err
	}
	return records, nil
}
