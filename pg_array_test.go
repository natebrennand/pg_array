package pgarray

import "testing"

type stringTestCase struct {
	Wire     []byte
	Expected []string
}

var stringArrayTestData = []stringTestCase{
	{
		Wire:     []byte(`{"name"}`),
		Expected: []string{"name"},
	},
	{
		Wire:     []byte(`{"last, first"}`),
		Expected: []string{"last, first"},
	},
	{
		Wire:     []byte(`{"last1, first1","last2, first2"}`),
		Expected: []string{"last1, first1", "last2, first2"},
	},
}

func TestStringArrayScan(t *testing.T) {
	var arr SqlStringArray
	for _, s := range stringArrayTestData {
		arr.Scan(s.Wire)

		if len(s.Expected) != len(arr.Data) {
			t.Errorf("EXPECTED: %v", s.Expected)
			t.Errorf("FOUND: %v", arr.Data)
			t.Fatal("Scan did not yield enough values")
		}

		for i, d := range s.Expected {
			if d != arr.Data[i] {
				t.Errorf("Mismatch, expected %s, found %s", d, arr.Data[i])
			}
		}
	}
}

type intTestCase struct {
	Wire     []byte
	Expected []int64
}

var intArrayTestData = []intTestCase{
	{
		Wire:     []byte(`{1}`),
		Expected: []int64{1},
	},
	{
		Wire:     []byte(`{1,2,3,4,5}`),
		Expected: []int64{1, 2, 3, 4, 5},
	},
	{
		Wire:     []byte(`{1234,1234,12345}`),
		Expected: []int64{1234, 1234, 12345},
	},
}

func TestIntArrayScan(t *testing.T) {
	var arr SqlIntArray
	for _, s := range intArrayTestData {
		arr.Scan(s.Wire)

		if len(s.Expected) != len(arr.Data) {
			t.Errorf("EXPECTED: %v", s.Expected)
			t.Errorf("FOUND: %v", arr.Data)
			t.Fatal("Scan did not yield enough values")
		}

		for i, d := range s.Expected {
			if d != arr.Data[i] {
				t.Errorf("Mismatch, expected %d, found %d", d, arr.Data[i])
			}
		}
	}
}
