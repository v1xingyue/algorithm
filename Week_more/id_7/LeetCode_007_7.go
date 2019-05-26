package geekcode 
import (
    "math"
)
func reverse(x int) int {
    s := 0 
    fh := 1 
    if x < 0 {
        x = -x
        fh = -1
    }
    
    for {
        m := x % 10
        x = x / 10
        s = s * 10 + m 
        if x == 0 {
            break
        }
    }
    
    if s >= math.MaxInt32 {
        return 0
    }
    return fh * s
}
