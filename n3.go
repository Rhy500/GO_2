package main
import "fmt"
func main() {
    type klub struct {
        nama string
        point, selisihGol int
    }
    var k1, k2, k3 klub
    
    fmt.Scan(&k1.nama, &k1.point, &k1.selisihGol)
    fmt.Scan(&k2.nama, &k2.point, &k2.selisihGol)
    fmt.Scan(&k3.nama, &k3.point, &k3.selisihGol)
    
    if k1.point < k2.point && k1.point < k3.point {
        fmt.Println(k1.nama)
    } else if k2.point < k1.point && k2.point < k3.point {
        fmt.Println(k2.nama)
    } else {
       fmt.Println(k3.nama) 
    }
    
    if k1.selisihGol < k2.selisihGol && k1.selisihGol < k3.selisihGol{
        fmt.Println(k1.nama)
    } else if k2.selisihGol < k1.selisihGol && k2.selisihGol < k3.selisihGol {
        fmt.Println(k2.nama)
    } else {
       fmt.Println(k3.nama) 
    }
}