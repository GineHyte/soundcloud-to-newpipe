package main

import (
	"os"

	args "github.com/GineHyte/sc_to_np/utils/args"
	start "github.com/GineHyte/sc_to_np/utils/start"
	storage "github.com/GineHyte/sc_to_np/utils/storage"
)

func main() {
	ar := os.Args[1:]
	if len(ar) == 0 {
		args.Help()
	} else {
		//pase all agrs
		for i := 0; i < len(ar)-1; i++ {
			switch ar[i] {
			case "-h", "--help":
				args.Help()
				return
			case "-v", "--version":
				args.Version()
				return
			case "-u", "--userid":
				args.UserId(ar[i+1])
			case "-c", "--clientid":
				args.ClientId(ar[i+1])
			case "-t", "--token":
				args.Token(ar[i+1])
			}
		}
		args.Output(ar[len(ar)-1])
	}

	start.Init(storage.Args)
}
