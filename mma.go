package mma

import "fmt"
import "errors"

const VERSION string = "0.0.0"

type MMA struct {
  length int
  n []float64
  avg []float64
  // Side Calculations
  count int
  momentum float64
  pivot bool
}

func New(n ...float64) *MMA {
  length := len(n)
  avg := make([]float64, length)
  return &MMA{length, n, avg, 0, 0.0, false}
}

func (mma *MMA) Init(p float64) {
  for i:=0; i<mma.length; i++ {
    mma.avg[i] = p
  }
  // Side Calculations
  mma.count = 1
  mma.momentum = 0.0
  mma.pivot = false
}

func pivoted(a float64, b float64) bool {
  x, y := 0 , 0
  if a > 0.0 { x=1 } else if a < 0.0 { x=-1}
  if b > 0.0 { y=1 } else if b < 0.0 { y=-1}
  return x != y
}

func (mma *MMA) Add(p float64) {
  var a, n float64
  p0 := mma.momentum
  a0 := mma.avg[0]
  for i:=0; i<mma.length; i++ {
    a = mma.avg[i]
    n = mma.n[i]
    a = ((n - 1.0)*a + p) / n
    mma.avg[i] = a
  }
  mma.count++
  mma.momentum = mma.avg[0] - a0
  if float64(mma.count) > mma.n[0] {
    mma.pivot = pivoted(mma.momentum, p0)
  }
}

func (mma *MMA) Find(n float64) (float64, error) {
  for i:=0; i<mma.length; i++ {
    if mma.n[i] == n { return mma.avg[i], nil }
  }
  return 0.0, errors.New("Not found.")
}

func (mma *MMA) String() string {
  s := "mma["
  for i:=0; i<mma.length; i++ {
    s += fmt.Sprintf(" %v:%v", mma.n[i], mma.avg[i])
  }
  s += fmt.Sprintf(" ]{ count:%v momentum:%v pivot:%v }", mma.count, mma.momentum, mma.pivot)
  return s
}
