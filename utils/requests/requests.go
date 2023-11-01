package requests

import (
	"fmt"
	"log"
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
	if err != nil {
		log.Printf(models.RED, resp.StatusCode, models.RESET)
		log.Printf(models.RED, resp.Body, models.RESET)
	}
	tools.Errors(err, 1)

	return userData
}

func GetLikes(link string, currentIndex int32) models.Likes {
	// get likes
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("OAuth %s", storage.Args.Token)},
	}
	tools.Errors(err, 1)

	resp, err := client.Do(req)
	tools.Errors(err, 1)

	defer resp.Body.Close()

	// parse likes
	var likes models.Likes
	err = tools.JsonDecode(resp.Body, &likes)
	if err != nil {
		log.Printf(models.RED, resp.StatusCode, models.RESET)
		log.Printf(models.RED, resp.Body, models.RESET)
	}
	tools.Errors(err, 1)

	return likes
}

func GetPlaylists(link string) models.PlaylistsSoundcloud {
	// get playlist
	client := &http.Client{}
	req, err := http.NewRequest("GET", link, nil)
	req.Header = http.Header{
		"Authorization": []string{fmt.Sprintf("OAuth %s", storage.Args.Token)},
	}
	tools.Errors(err, 1)

	resp, err := client.Do(req)
	tools.Errors(err, 1)

	defer resp.Body.Close()

	// parse playlist
	var playlist models.PlaylistsSoundcloud
	err = tools.JsonDecode(resp.Body, &playlist)
	if err != nil {
		log.Printf(models.RED, resp.StatusCode, models.RESET)
		log.Printf(models.RED, resp.Body, models.RESET)
	}
	tools.Errors(err, 1)

	return playlist
}
