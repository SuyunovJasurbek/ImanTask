package service

// Console ishladi , endi shuni API ga boglash kerak :)
import (
	"fmt"
	"time"
)

func (s *Service) Time() int {
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

		// 2025 yildagi biror sanadan 1-fevralgacha bulgn kundagi xolat uchun
		if t_year == 2025 {
			Sum = t_day
			break
		}

		// 2025 yilgacha bulgan davrdagi xoxlagan vaqt uchun
		if i == t_year {
			if !Leap(t_year) {
				Sum = (month_before[12] - (month_before[int(t_month)-1]) - t_day) + Sum
			} else {
				Sum = (month_before[12] - (month_before[int(t_month)-1]) - t_day) + Sum + 1
			}
		} else if i > t_year && i < 2025 {
			if Leap(i) {
				Sum = Sum + 366
			} else {
				Sum = Sum + 365
			}
		} else if i == 2025 {
			Sum = Sum + 31 // Februar
			break
		}
	}
	fmt.Println(Sum)
	return Sum
}

func Leap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}