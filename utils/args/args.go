package args

import (
	"fmt"

	storage "github.com/GineHyte/sc_to_np/utils/storage"
)

func Help() {
	fmt.Println("Usage: sc_to_np [options] [output_file]")
	fmt.Println("Options:")
	fmt.Println("  -h, --help\t\t\t\tShow this help message and exit")
	fmt.Println("  -v, --version\t\t\t\tShow version number and exit")
	fmt.Println("  -u, --userid\t\t\t\tUser ID (numbers)")
	fmt.Println("  -t, --token\t\t\t\tToken (X-XXXXXX-XXXXXXXXX-...)")
}

func Version() {
	fmt.Println("sc_to_np 0.1.0")
}

func UserId(user_id string) {
	storage.Args.UserId = user_id
}

func Token(token string) {
	storage.Args.Token = token
}

func Output(output_file string) {
	storage.Args.Output = output_file
}
