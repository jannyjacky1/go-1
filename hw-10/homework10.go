package main

import (
	"errors"
	"fmt"
	flag "github.com/spf13/pflag"
	"io"
	"os"
)

func Copy(from string, to string, limit int, offset int) error {
	if from == "" {
		return errors.New("required variable --from is missed")
	}
	if to == "" {
		return errors.New("required variable --to is missed")
	}

	source, err := os.OpenFile(from, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer source.Close()

	target, err := os.OpenFile(to, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()

	source.Seek(int64(offset), io.SeekStart)

	if limit > 0 {
		_, err = io.CopyN(target, source, int64(limit))
	} else {
		_, err = io.Copy(target, source)
	}

	if err != nil && err != io.EOF {
		return err
	}

	return nil
}

func main() {
	var from, to string
	var limit, offset int

	flag.StringVar(&from, "from", "", "")
	flag.StringVar(&to, "to", "", "")
	flag.IntVar(&limit, "limit", 0, "")
	flag.IntVar(&offset, "offset", 0, "")

	flag.Parse()

	err := Copy(from, to, limit, offset)

	if err == nil {
		fmt.Println("Copy completed")
	} else {
		fmt.Println(err)
	}
}
