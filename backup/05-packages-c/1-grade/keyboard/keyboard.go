package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	str = strings.TrimSpace(str)
	num, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return 0, err
	}

	return num, nil
}