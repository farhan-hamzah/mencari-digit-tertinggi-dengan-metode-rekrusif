package main
import "fmt"

func digitPuncak(n int) int {
	if n < 10 {
		return n
	}
	akhir := n % 10
	sisa := digitPuncak(n / 10)

	if akhir > sisa {
		return akhir
	}
	return sisa
}

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(digitPuncak(num))
}