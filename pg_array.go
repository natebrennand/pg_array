package pg_array

import (
	"bytes"
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
		if err != nil && intBytes[0] == '}' {
			break
		}
		num, _ := strconv.ParseInt(string(intBytes[0:len(intBytes)-1]), 10, 64)
		s.Data = append(s.Data, num)
	}

	return nil
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
		intBytes, err := buf.ReadBytes(',')
		if err != nil && intBytes[0] == '}' {
			break
		}
		s.Data = append(s.Data, string(intBytes[0:len(intBytes)-1]))
	}

	return nil
}
