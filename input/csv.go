package input

import (
	"encoding/csv"
	"errors"
	"fmt"
)

func CSVRead(filename string) error {
	fmt.Println("Begin parse filename: ", filename)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	fields, err := reader.Read()
	if err != nil {
		return err
	}

}
