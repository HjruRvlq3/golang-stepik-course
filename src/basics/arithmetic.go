package basics

import "strconv"

func Square(num int) int {
	return num * num
}

func LastDigit(num int) int {
	return num % 10
}

func TensDigit(num int) int {
	return num % 100 / 10
}

// ClockFromDegrees вычисляет часы и минуты
// по градусам часовой стрелки
func ClockFromDegrees(d int) string {
	h := d / 30
	m := d % 30 * 2

	s := "It is " + strconv.Itoa(h) +
		" hours " + strconv.Itoa(m) + " minutes."

	return s
}

// SignMessage возвращает сообщение о знаке числа n
func SignMessage(num int) string {
	var strResult string

	switch {
	case num > 0:
		strResult = "Число положительное"
	case num < 0:
		strResult = "Число отрицательное"
	case num == 0:
		strResult = "Ноль"
	}

	return strResult
}

// AllDigitsDifferent проверяет,
// все ли цифры трехзначного числа различны
func AllDigitsDifferent(num int) string {
	var (
		a         int
		b         int
		c         int
		strResult string
	)

	a = num % 10
	b = num % 100 / 10
	c = num / 100

	if a != b && a != c && b != c {
		strResult = "YES"
	} else {
		strResult = "NO"
	}

	return strResult
}

// FirstDigit возвращает первую цифру неотрицательного числа n
func FirstDigit(num int) int {
	for num >= 10 {
		num = num / 10
	}
	return num
}

// IsLuckyTicket проверяет, является ли шестизначный номер билета счастливым
func IsLuckyTicket(ticketNum int) string {
	var (
		sum1      int // сумма первых трех цифр
		sum2      int // сумма последних трех цифр
		strResult string
	)

	sum2 = ticketNum % 10
	ticketNum = ticketNum / 10
	sum2 = sum2 + ticketNum%10
	ticketNum = ticketNum / 10
	sum2 = sum2 + ticketNum%10

	ticketNum = ticketNum / 10
	sum1 = ticketNum % 10
	ticketNum = ticketNum / 10
	sum1 = sum1 + ticketNum%10
	ticketNum = ticketNum / 10
	sum1 = sum1 + ticketNum%10

	if sum2 == sum1 {
		strResult = "YES"
	} else {
		strResult = "NO"
	}

	return strResult
}

// IsLeapYear проверяет, является ли год високосным
func IsLeapYear(year int) string {
	var strResult string

	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		strResult = "YES"
	} else {
		strResult = "NO"
	}

	return strResult
}

// SumRange возвращает сумму всех чисел от a до b включительно (a <= b)
func SumRange(a, b int) int {
	var sumRange int

	if a > b {
		sumRange = 0
	} else {
		for a <= b {
			sumRange += a
			a++
		}
	}

	return sumRange
}

// SumTwoDigitMultiplesOf8 возвращает сумму всех двузначных чисел из среза nums, кратных 8
func SumTwoDigitMultiplesOf8(nums []int) int {
	sum := 0

	for _, a := range nums {
		if a >= 10 && a <= 99 && a%8 == 0 {
			sum += a
		}
	}

	return sum
}

// CountMaxElements возвращает количество элементов в nums, равных максимальному
func CountMaxElements(nums []int) int {
	var max, count int

	for _, i := range nums {
		if i > max {
			max = i
			count = 1
		} else if i == max {
			count++
		}
	}

	return count

}

// FirstMultipleButNot возвращает первое число от 1 до n,
// кратное c, но не кратное d. Если такого числа нет, возвращает 0.
func FirstMultipleButNot(n, c, d int) int {
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			return i
		}
	}
	return 0
}
