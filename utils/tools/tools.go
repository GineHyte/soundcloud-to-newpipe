package tools

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

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

func LikesToStreams(likes models.Likes, startId int64) []models.Steam {
	var streams []models.Steam

	sLayout := "2006-01-02T15:04:05Z"

	for i, like := range likes.Collection {
		tTime, _ := time.Parse(sLayout, like.CreatedAt)
		streams = append(streams, models.Steam{
			Uid:                       startId + int64(i),
			ServiceId:                 1,
			Url:                       like.Track.PermalinkUrl,
			Title:                     like.Track.Title,
			StreamType:                "AUDIO_STREAM",
			Duration:                  like.Track.Duration,
			Uploader:                  like.Track.User.Username,
			UploaderUrl:               like.Track.User.PermalinkUrl,
			ThumbnailUrl:              like.Track.ArtworkUrl,
			ViewCount:                 like.Track.PlaybackCount,
			TextualUploadDate:         like.CreatedAt,
			UploadDate:                tTime.Unix(),
			IsUploadDateApproximation: false,
		})
	}
	return streams
}