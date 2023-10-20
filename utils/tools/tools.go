package tools

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/GineHyte/sc_to_np/models"
)

func Errors(err error, errorLevel int) {
	if err != nil {
		switch errorLevel {
		case 0:
			fmt.Println(models.RED, err, models.RESET)
		case 1:
			fmt.Println(models.RED, err, models.RESET)
			panic(err)
		}
	}
}

func JsonDecode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
