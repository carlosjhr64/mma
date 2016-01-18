package mma // import "github.com/carlosjhr64/mma"

const VERSION string = "0.0.0"
func New(n ...float64) *MMA
type MMA struct { ... }

Example:
  // Have a list of prices([]float64).
  mmas := mma.New(5.0, 20.0, 90.0)
  mmas.Init(price[0])
  for i:=1; i<len(price); i++ { mmas.Add(price[i]) }
  mma5 := mma.avg[0]
  mma20 := mma.avg[1]
  mma90 := mma.avg[2]
  // Also has...
  if mmas.Count == len(price) {
    // Yep, counted correctly.
  }
  if mmas.Momentum > 0.0 {
    // The week mma is rissing.
    // The momentum is based on the first average, mma.avg[0].
  }
  if mmas.Pivot {
    // The momentum switched sign.
    // That is if it was going up, it's now going down (and vice-versa).
  }
