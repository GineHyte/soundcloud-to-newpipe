package start

import (
	"fmt"
	"log"

	models "github.com/GineHyte/sc_to_np/models"
	requests "github.com/GineHyte/sc_to_np/utils/requests"
	storage "github.com/GineHyte/sc_to_np/utils/storage"
)

func Init(args models.Args) {
	fmt.Printf(models.CLEAR)
	log.Printf("Starting sc_to_np with parameters: %v", args)

	// get user data
	userData := requests.GetUserData()
	storage.UserData = userData

	println("User data:")
	fmt.Printf("  UserId: %v\n", userData.UserId)
	fmt.Printf("  Fullname: %s\n", userData.FullName)
	fmt.Printf("  LikesCount: %v\n", userData.LikesCount)
	fmt.Printf("  PlaylistsLikesCount: %v\n", userData.PlaylistsLikesCount)

	// get likes

}
