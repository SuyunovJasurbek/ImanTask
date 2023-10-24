package main
// Console ishladi , endi shuni API ga boglash kerak :)
import (
	"fmt"
	"time"
)

func main() {
	var Sum int
	t_now := time.Now()
	t_day := t_now.Day()
	t_month := t_now.Month()
	t_year := t_now.Year()

	var month_before = [13]int{
		0,
		31,                                    //yanvar
		31 + 28,                               //fevral
		31 + 28 + 31,                          //mart
		31 + 28 + 31 + 30,                     //aprel
		31 + 28 + 31 + 30 + 31,                //may
		31 + 28 + 31 + 30 + 31 + 30,           //iyun
		31 + 28 + 31 + 30 + 31 + 30 + 31,      //iyul
		31 + 28 + 31 + 30 + 31 + 30 + 31 + 31, //avgust
		31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,                //sentabr
		31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,           //oktabr
		31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,      //noyabr
		31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31, //dekabr
	}
	for i := t_year; i <= 2025; i++ {
		fmt.Println(i)
		if i == t_year {
			if !Leap(t_year) {
				Sum = (month_before[12] - (month_before[int(t_month)-1]) - t_day) + Sum
				fmt.Println("{{{{{{{{{{{{{{1", Sum, "}}}}}}}}}}}}}}")
			} else {
				Sum = (month_before[12] - (month_before[int(t_month)-1]) - t_day) + Sum + 1
				fmt.Println("{{{{{{{{{{{{{{2", Sum, "}}}}}}}}}}}}}}")
			}
		} else if i > t_year && i < 2025 {
			if Leap(i) {
				Sum = Sum + 366
				fmt.Println("{{{{{{{{{{{{{{3", Sum, "}}}}}}}}}}}}}}")
			} else {
				Sum = Sum + 365
				fmt.Println("{{{{{{{{{{{{{{4", Sum, "}}}}}}}}}}}}}}")
			}
		} else if i == 2025 {
			Sum = Sum + 31 // Februar
			fmt.Println("{{{{{{{{{{{{{{5", Sum, "}}}}}}}}}}}}}}")
			break
		}
	}
	fmt.Println(Sum)
}

func Leap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
