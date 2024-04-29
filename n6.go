package main
import "fmt"
const NMax int = 10000
type tabInt [NMax]int
func InputArray(Tab *tabInt, n int){
    /*I.S. Terdefinisi nilai bilangan bulat n*/
    /*Proses : Menerima masukan bilangan bulat. Proses input akan berhenti 
               ketika sudah mencapai kapasitas maksimum yaitu n*/
    /*F.S. Array Tab berisi data yang diberikan*/
    var i int
    for i=0; i<n; i++{
        fmt.Scan(&Tab[i])
    }
}

func jumlahArrGenap(Tab tabInt, n int)int{
    /*Mengembalikan hasil dari penjumlahan array yang bernilai genap 
    yang terdapat didalam array dengan tipe data integer*/
    var i, sum int
    sum = 0
    for i=0; i<n; i++{
        if Tab[i]%2 == 0 {
			sum = sum + Tab[i]
		}
    }
    return sum
}

func main() {
	var Tab tabInt
	var n int
	fmt.Scan(&n)
	InputArray(&Tab,n)
	fmt.Println(jumlahArrGenap(Tab,n))
}