package mma

import "testing"
import "fmt"

var puts = fmt.Println

func TestMMA(test *testing.T) {
  bad := test.Error
  var s string

  a := New(2.0, 3.0, 4.0)
  a.Init(1.0)
  s = a.String()
  if s != "{ 2:1 3:1 4:1 }[ 1 ]" { bad("Init or String") }

  a.Add(1.0)
  s = a.String()
  if s != "{ 2:1 3:1 4:1 }[ 2 ]" { bad("Add or String") }

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

  if a.count != 3 { bad("count") }
}
