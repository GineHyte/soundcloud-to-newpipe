package models

var CLEAR = "\033[2J"
var RESET = "\033[0m"
var BOLD = "\033[1m"
var RED = "\033[31m"
var GREEN = "\033[32m"
var YELLOW = "\033[33m"
var BLUE = "\033[34m"
var MAGENTA = "\033[35m"

type Args struct {
	UserId   string
	ClientId string
	Token    string
	Output   string
}

type Steam struct {
	Uid                       int64
	ServiceId                 int16
	Url                       string
	Title                     string
	StreamType                string
	Duration                  int32
	Uploader                  string
	UploaderUrl               string
	ThumbnailUrl              string
	ViewCount                 int32
	TextualUploadDate         string
	UploadDate                int64
	IsUploadDateApproximation bool
}

type Playlist struct {
	Uid                   int64
	Name                  string
	IsThrumbnailPermanent bool
	ThumbnailStreamId     string
}

type PlaylistStreamJoin struct {
	PlaylistId int64
	StreamId   int64
	JoinIndex  int64
}

type Subscription struct {
	Uid              int64
	ServiceId        int16
	Url              string
	Name             string
	AvatarUrl        string
	SubscriberCount  int32
	Description      string
	NotificationMode int32
}

type UserData struct {
	UserId              int64  `json:"id"`
	FullName            string `json:"full_name"`
	LikesCount          int32  `json:"likes_count"`
	PlaylistsLikesCount int32  `json:"playlist_likes_count"`
}

type Likes struct {
	Collection []struct {
		Track struct {
			PermalinkUrl string `json:"permalink_url"`
		}
	}
}
