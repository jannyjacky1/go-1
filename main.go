package main

import (
	hw1 "github.com/jannyjacky1/go-1/hw-1"
	hw2 "github.com/jannyjacky1/go-1/hw-2"
	hw6 "github.com/jannyjacky1/go-1/hw-6"
	"log"
)

func runHomework(num int) {
	switch num {
	case 1:
		hw1.Run()
	case 2:
		str, err := hw2.UnpackString("mystring7")
		if err != nil {
			log.Fatalf("Ошибка в задании %d: %s", num, err)
		}
		println(str)
	case 6:
		hw6.Check()
	default:
		log.Fatalf("Некорректный номер ДЗ: %d", num)
	}
}

func main() {
	runHomework(2)
}
