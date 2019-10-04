package options

import (
	"bufio"
	"os"
	"strconv"
)

type Options struct {
	Now  int64
	Test bool
}

func NewOptionsFromFile(filePath string) (*Options, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	s := scanner.Text()
	now, _ := strconv.ParseInt(s, 10, 64)

	scanner.Scan()
	s = scanner.Text()
	test := s == "0"

	return &Options{
		Now:  now,
		Test: test,
	}, nil
}