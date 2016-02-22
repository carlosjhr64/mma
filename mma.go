package mma

import "fmt"
import "math"
import "errors"

const VERSION string = "0.2.0"

type MMA struct {
  Length int
  N []float64
  Avg []float64
  // Side Calculations
  Count int
  Momentum float64
  Pivot bool
  UseLog bool
}

func New(n ...float64) *MMA {
  length := len(n)
  avg := make([]float64, length)
  return &MMA{length, n, avg, 0, 0.0, false, false}
}

func (mma *MMA) Init(p float64) {
  for i:=0; i<mma.Length; i++ {
    mma.Avg[i] = p
  }
  // Side Calculations
  mma.Count = 1
  mma.Momentum = 0.0
  mma.Pivot = false
}

func pivoted(a float64, b float64) bool {
  x, y := 1 , 1 // Assume >= 0.0
  if a < 0.0 { x=-1 }
  if b < 0.0 { y=-1 }
  return x != y
}

func (mma *MMA) Add(p float64) {
  var a, n float64
  p0 := mma.Momentum
  a0 := mma.Avg[0]
  for i:=0; i<mma.Length; i++ {
    a = mma.Avg[i]
    n = mma.N[i]
    if mma.UseLog {
      a = math.Exp(((n - 1.0)*math.Log(a) + math.Log(p)) / n)
    } else {
      a = ((n - 1.0)*a + p) / n
    }
    mma.Avg[i] = a
  }
  mma.Count++
  mma.Momentum = mma.Avg[0] - a0
  if float64(mma.Count) > mma.N[0] {
    mma.Pivot = pivoted(mma.Momentum, p0)
  }
}

func (mma *MMA) Find(n float64) (float64, error) {
  for i:=0; i<mma.Length; i++ {
    if mma.N[i] == n { return mma.Avg[i], nil }
  }
  return 0.0, errors.New("Not found.")
}

func (mma *MMA) String() string {
  s := "mma["
  for i:=0; i<mma.Length; i++ {
    s += fmt.Sprintf(" %v:%v", mma.N[i], mma.Avg[i])
  }
  s += fmt.Sprintf(
    " ]{ %v, %v, %v }",
    mma.Count, mma.Momentum, mma.Pivot)
  return s
}
