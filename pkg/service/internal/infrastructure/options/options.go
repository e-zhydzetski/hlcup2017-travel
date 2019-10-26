package options

import (
	"bufio"
	"os"
	"strconv"

	"github.com/e-zhydzetski/hlcup2017-travel/pkg/service/internal/app"
)

func NewOptionsFromFile(filePath string) (*app.Options, error) {
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

	return &app.Options{
		Now:  now,
		Test: test,
	}, nil
}