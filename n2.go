package main
import (
    "fmt"
    "math"
)

func main() {
    type titik struct {
        x, y float64
    }
    var A, B, C, D titik
    
    fmt.Scan(&A.x, &A.y)
    fmt.Scan(&B.x, &B.y)
    fmt.Scan(&C.x, &C.y)
    fmt.Scan(&D.x, &D.y)
    
    dAB := math.Sqrt(math.Pow((A.x - B.x), 2) + math.Pow((A.y - B.y), 2))
    dBC := math.Sqrt(math.Pow((B.x - C.x), 2) + math.Pow((B.y - C.y), 2))
    dCD := math.Sqrt(math.Pow((C.x - D.x), 2) + math.Pow((C.y - D.y), 2))
    dDA := math.Sqrt(math.Pow((D.x - A.x), 2) + math.Pow((D.y - A.y), 2))
    
    fmt.Println(dAB + dBC + dCD + dDA)
}