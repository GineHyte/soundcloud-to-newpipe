package requests

import (
	"fmt"
	"net/http"

	models "github.com/GineHyte/sc_to_np/models"
	storage "github.com/GineHyte/sc_to_np/utils/storage"
	tools "github.com/GineHyte/sc_to_np/utils/tools"
)

func GetUserData() models.UserData {
	// get user data
	headers := http.Header{
		"Authorization": []string{fmt.Sprintf("OAuth %s", storage.Args.Token)},
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://api-v2.soundcloud.com/users/%s", storage.Args.UserId), nil)
	tools.Errors(err, 1)
	req.Header = headers

	resp, err := client.Do(req)
	tools.Errors(err, 1)

	defer resp.Body.Close()

	// parse user data
	var userData models.UserData
	err = tools.JsonDecode(resp.Body, &userData)
	tools.Errors(err, 1)

	return userData
}

func GetLikes() {
	// get likes
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://api-v2.soundcloud.com/users/%s/track_likes?client_id=%s&limit=200&linked_partitioning=1",
		storage.Args.UserId, storage.Args.ClientId), nil)
	tools.Errors(err, 1)

	resp, err := client.Do(req)
	tools.Errors(err, 1)

	defer resp.Body.Close()

	// parse likes
	var likes models.Likes
	err = tools.JsonDecode(resp.Body, &likes)
	tools.Errors(err, 1)

}
