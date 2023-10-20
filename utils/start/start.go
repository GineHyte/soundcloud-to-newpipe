package start

import (
	"fmt"
	"log"

	models "github.com/GineHyte/sc_to_np/models"
)

func Init(args models.Args) {
	fmt.Printf(models.CLEAR)
	log.Printf("Starting sc_to_np with parameters: %v", args)
}
