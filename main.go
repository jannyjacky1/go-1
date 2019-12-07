package main

import (
	hw2 "github.com/jannyjacky1/go-1/hw-2"
	"github.com/labstack/gommon/log"
)

func runHomework(num int) {
	switch num {
	case 2:
		str, err := hw2.UnpackString("mystring7")
		if err != nil {
			log.Fatalf("Ошибка в задании %d: %s", num, err)
		}
		println(str)
	default:
		log.Fatalf("Некорректный номер ДЗ: %d", num)
	}
}

func main() {
	runHomework(2)
}
