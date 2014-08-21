package pg_array

import (
	"bytes"
	"fmt"
	"log"
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
	s.Data = make([]int64, buf.Len()/5)

	b, err := buf.ReadByte()
	if err != nil || b != '{' {
		log.Fatalf("Failed to read in first byte of array")
	}

	numIndex := 0
	for buf.Len() > 0 { // greater than 1 number left
		intBytes, err := buf.ReadBytes(',')
		if err != nil && intBytes[0] == '}' {
			break
		}
		num, _ := strconv.ParseInt(string(intBytes[0:len(intBytes)-1]), 10, 64)
		s.Data[numIndex] = num
		numIndex += 1
	}

	return nil
}
