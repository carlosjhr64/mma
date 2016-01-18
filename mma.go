package mma

import "fmt"
import "errors"

const VERSION string = "0.0.0"

type MMA struct {
  length int
  n []float64
  avg []float64
  count int
}

func New(n ...float64) *MMA {
  length := len(n)
  avg := make([]float64, length)
  return &MMA{length, n, avg, 0}
}

func (mma *MMA) Init(p float64) {
  for i:=0; i<mma.length; i++ {
    mma.avg[i] = p
  }
  mma.count = 1
}

func (mma *MMA) Add(p float64) {
  var a, n float64
  for i:=0; i<mma.length; i++ {
    a = mma.avg[i]
    n = mma.n[i]
    a = ((n - 1.0)*a + p) / n
    mma.avg[i] = a
  }
  mma.count++
}

func (mma *MMA) Find(n float64) (float64, error) {
  for i:=0; i<mma.length; i++ {
    if mma.n[i] == n { return mma.avg[i], nil }
  }
  return 0.0, errors.New("Not found.")
}

func (mma *MMA) String() string {
  s := "{"
  for i:=0; i<mma.length; i++ {
    s += fmt.Sprintf(" %v:%v", mma.n[i], mma.avg[i])
  }
  s += fmt.Sprintf(" }[ %v ]", mma.count)
  return s
}
