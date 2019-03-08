package doubler

import (
  "doubler"
  "testing"
)

func TestDouble(t *testing.T) {
  v := doubler.Double(5)
  if v != 10 {
    t.Error("Expected 10, got ", v)
  }
}

type testdata struct {
  value int
  result int
}

var testvalues = []testdata{
  {1,2},
  {2,5},
  {3,6},
}

func TestMultipleValues(t *testing.T) {
  for _, integer := range testvalues {
    v := doubler.Double(integer.value)
    if v != integer.result {
      t.Error(
        "For", integer.value,
        "expect", integer.result,
        "got", v,
      )
    }
  }
}

// The above mentioned returns the following

// --- FAIL: TestMultipleValues (0.00s)
// 	testing1_test.go:30: For 2 expect 5 got 4
// FAIL
// exit status 1
// FAIL	_/home/x/x/Introducing-Go/chapter-9	0.001s



/*
Sample fail
--- FAIL: TestDouble (0.00s)
	testing1_test.go:11: Expected 11, got  10
FAIL
exit status 1
FAIL	_/home/x/x/Introducing-Go/chapter-9	0.001s

Sample pass
PASS
ok  	_/home/x/x/Introducing-Go/chapter-9	0.001s
*/
