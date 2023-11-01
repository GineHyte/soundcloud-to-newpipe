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
		CreatedAt string `json:"created_at"`
		Track     struct {
			ArtworkUrl    string `json:"artwork_url"`
			PermalinkUrl  string `json:"permalink_url"`
			PlaybackCount int32  `json:"playback_count"`
			Title         string `json:"title"`
			Duration      int32  `json:"duration"`
			User          struct {
				Username     string `json:"username"`
				PermalinkUrl string `json:"permalink_url"`
			} `json:"user"`
		} `json:"track"`
	} `json:"collection"`
	NextHref string `json:"next_href"`
}

type PlaylistsSoundcloud struct {
	Collection []struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
		Playlist struct {
			Title        string `json:"title"`
			PermalinkUrl string `json:"permalink_url"`
			Public       bool   `json:"public"`
			ArtworkUrl   string `json:"artwork_url"`
			TrackCount   int32  `json:"track_count"`
		} `json:"playlist"`
	} `json:"collection"`
	NextHref string `json:"next_href"`
}

type RemotePlaylist struct {
	Uid          int64
	ServiceId    int16
	Name         string
	Url          string
	ThumbnailUrl string
	Uploder      string
	StreamCount  int32
}
