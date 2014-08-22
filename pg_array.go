package pg_array

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type SqlIntArray struct {
	Data []int64
}

func (s SqlIntArray) String() string {
	return fmt.Sprintf("%v, ", s.Data)
}

func (s *SqlIntArray) Scan(src interface{}) error {
	buf := bytes.NewBuffer(src.([]byte))
	s.Data = make([]int64, 0)

	b, err := buf.ReadByte()
	if err != nil || b != '{' {
		return fmt.Errorf("Failed to read in first byte of array")
	}

	for buf.Len() > 0 { // greater than 1 number left
		intBytes, err := buf.ReadBytes(',')
		if err != nil && len(intBytes) == 1 {
			break
		}
		num, _ := strconv.ParseInt(string(intBytes[0:len(intBytes)-1]), 10, 64)
		s.Data = append(s.Data, num)
	}

	return nil
}

func (s SqlIntArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Data)
}

type SqlStringArray struct {
	Data []string
}

func (s SqlStringArray) String() string {
	return fmt.Sprintf("%v, ", s.Data)
}

func (s *SqlStringArray) Scan(src interface{}) error {
	buf := bytes.NewBuffer(src.([]byte))
	s.Data = make([]string, 0)

	b, err := buf.ReadByte()
	if err != nil || b != '{' {
		return fmt.Errorf("Failed to read in first byte of array")
	}

	for buf.Len() > 0 { // greater than 1 number left
		var stringBytes []byte
		for {
			bufBytes, err := buf.ReadBytes(',')
			if err != nil && len(bufBytes) == 1 { // EOF w/ '}'
				return nil
			}
			stringBytes = append(stringBytes, bufBytes...)

			// break if empty or actual word delimeter
			if len(buf.Bytes()) == 0 || buf.Bytes()[0] == '"' {
				break
			}
		}

		// offset to account for quotations
		s.Data = append(s.Data, string(stringBytes[1:len(stringBytes)-2]))
	}

	return nil
}

func (s SqlStringArray) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Data)
}
