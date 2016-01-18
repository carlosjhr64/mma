package mma

import "testing"

func TestMMA(test *testing.T) {
  bad := test.Error
  var s string

  a := New(2.0, 3.0, 4.0)

  a.Init(1.0)

  s = a.String()
  if s != "mma[ 2:1 3:1 4:1 ]{ count:1 momentum:0 pivot:false }" { bad("Init or String.") }

  a.Add(1.0)

  s = a.String()
  if s != "mma[ 2:1 3:1 4:1 ]{ count:2 momentum:0 pivot:false }" { bad("Add or String.") }

  a.Add(2.0)

  twoMma,   _ := a.Find(2.0)
  threeMma, _ := a.Find(3.0)
  fourMma,  _ := a.Find(4.0)
  if twoMma   != (1.0*1.0 + 2.0)/2.0 || // 1.5
     threeMma != (1.0*2.0 + 2.0)/3.0 || // 1.333...
     fourMma  != (1.0*3.0 + 2.0)/4.0 {  // 1.25
     bad("Add or Find")
  }
  if _, e := a.Find(5.0); e==nil { bad("Find error.") }

  if a.Momentum != 1.5 - 1.0 { bad("Momentum.") }
  if a.Count != 3 { bad("Count.") }

  if a.Pivot != false { bad("A: Pivot.") }
  a.Add(0.0)
  if a.Pivot != true { bad("B: Pivot.") }
}
