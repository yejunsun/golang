package newmath

import "testing"

func TestSqrt (t *testing.T) {
  in, out := 4.0, 2.0
  if x := Sqrt(in); x != out {
    t.Errorf("Sqrt(%v) = %v , want %v", in, x , out)
  }
}
