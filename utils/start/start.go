package start

import (
	"database/sql"
	"fmt"
	"log"

	models "github.com/GineHyte/sc_to_np/models"
	requests "github.com/GineHyte/sc_to_np/utils/requests"
	storage "github.com/GineHyte/sc_to_np/utils/storage"
	tools "github.com/GineHyte/sc_to_np/utils/tools"

	_ "github.com/mattn/go-sqlite3"
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
	var likes models.Likes
	link := fmt.Sprintf("http://api-v2.soundcloud.com/users/%s/track_likes?client_id=%s&limit=200&linked_partitioning=1",
		storage.Args.UserId, storage.Args.ClientId)
	currentIndex := int32(0)
	for currentIndex < userData.LikesCount {
		tempLikes := requests.GetLikes(link, currentIndex)

		if (tempLikes.Collection == nil) || (tempLikes.NextHref == "") {
			break
		}
		likes.Collection = append(likes.Collection, tempLikes.Collection...)
		link = tempLikes.NextHref
		currentIndex += int32(len(tempLikes.Collection))
	}
	storage.Likes = tools.LikesToStreams(likes, 0)

	println("Likes:")
	fmt.Printf("  LikesCount found: %v\n", len(storage.Likes))

	// pack streams into playlists
	// create "likes" playlist
	storage.Playlists = append(storage.Playlists, models.Playlist{
		Uid:                   1,
		Name:                  "Likes",
		IsThrumbnailPermanent: false,
		ThumbnailStreamId:     "1",
	})

	// create "likes" playlist join
	for i, _ := range storage.Likes {
		storage.PlaylistStreamJoins = append(storage.PlaylistStreamJoins, models.PlaylistStreamJoin{
			PlaylistId: 1,
			StreamId:   int64(i) + 1,
			JoinIndex:  int64(i),
		})
	}

	// create sql table in .db file
	CreateSQL()
}

// create sql table in .db file
func CreateSQL() {
	// open db
	db, err := sql.Open("sqlite3", storage.Args.Output)
	tools.Errors(err, 1)
	defer db.Close()

	log.Printf("Creating database file: %s", storage.Args.Output)
	// create tables
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS streams (
			uid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			service_id INTEGER NOT NULL,
			url TEXT NOT NULL,
			title TEXT NOT NULL,
			stream_type TEXT NOT NULL,
			duration INTEGER NOT NULL,
			uploader TEXT NOT NULL,
			uploader_url TEXT,
			thumbnail_url TEXT,
			view_count INTEGER,
			textual_upload_date TEXT,
			upload_date INTEGER,
			is_upload_date_approximation INTEGER
		)
	`)
	tools.Errors(err, 1)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS playlists (
			uid INTEGER PRIMARY KEY NOT NULL,
			name TEXT,
			is_thumbnail_permanent INTEGER NOT NULL,
			thumbnail_stream_id INTEGER NOT NULL
		)
	`)
	tools.Errors(err, 1)
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS playlist_stream_join (
			playlist_id INTEGER NOT NULL,
			stream_id INTEGER NOT NULL,
			join_index INTEGER NOT NULL,
			PRIMARY KEY(playlist_id, join_index),
			FOREIGN KEY(playlist_id) REFERENCES playlists(uid) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED,
			FOREIGN KEY(stream_id) REFERENCES streams(uid) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED
		)
	`)
	tools.Errors(err, 1)

	// create other tables
	log.Printf("Creating other tables in database file: %s", storage.Args.Output)
	_, _ = db.Exec("CREATE TABLE android_metadata (locale TEXT)")
	_, _ = db.Exec("INSERT INTO android_metadata VALUES ('en_US')")
	_, _ = db.Exec("CREATE TABLE feed (stream_id INTEGER NOT NULL, subscription_id INTEGER NOT NULL, PRIMARY KEY(stream_id, subscription_id), FOREIGN KEY(stream_id) REFERENCES streams(uid) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED, FOREIGN KEY(subscription_id) REFERENCES subscriptions(uid) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED)")
	_, _ = db.Exec("CREATE INDEX index_feed_subscription_id ON feed (subscription_id)")
	_, _ = db.Exec("CREATE TABLE feed_group (uid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, name TEXT NOT NULL, icon_id INTEGER NOT NULL, sort_order INTEGER NOT NULL)")
	_, _ = db.Exec("CREATE INDEX index_feed_group_sort_order ON feed_group (sort_order)")
	_, _ = db.Exec("CREATE TABLE feed_group_subscription_join (group_id INTEGER NOT NULL, subscription_id INTEGER NOT NULL, PRIMARY KEY(group_id, subscription_id), FOREIGN KEY(group_id) REFERENCES feed_group(uid) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED, FOREIGN KEY(subscription_id) REFERENCES subscriptions(uid) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED)")
	_, _ = db.Exec("CREATE INDEX index_feed_group_subscription_join_subscription_id ON feed_group_subscription_join (subscription_id)")
	_, _ = db.Exec("CREATE TABLE feed_last_updated (subscription_id INTEGER NOT NULL, last_updated INTEGER, PRIMARY KEY(subscription_id), FOREIGN KEY(subscription_id) REFERENCES subscriptions(uid) ON UPDATE CASCADE ON DELETE CASCADE DEFERRABLE INITIALLY DEFERRED)")
	_, _ = db.Exec("CREATE TABLE remote_playlists (uid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, service_id INTEGER NOT NULL, name TEXT, url TEXT, thumbnail_url TEXT, uploader TEXT, stream_count INTEGER)")
	_, _ = db.Exec("CREATE INDEX index_remote_playlists_name ON remote_playlists (name)")
	_, _ = db.Exec("CREATE UNIQUE INDEX index_remote_playlists_service_id_url ON remote_playlists (service_id, url)")
	_, _ = db.Exec("CREATE TABLE room_master_table (id INTEGER PRIMARY KEY,identity_hash TEXT)")
	_, _ = db.Exec("CREATE TABLE search_history (creation_date INTEGER, service_id INTEGER NOT NULL, search TEXT, id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL)")
	_, _ = db.Exec("CREATE INDEX index_search_history_search ON search_history (search)")
	_, _ = db.Exec("CREATE TABLE stream_history (stream_id INTEGER NOT NULL, access_date INTEGER NOT NULL, repeat_count INTEGER NOT NULL, PRIMARY KEY(stream_id, access_date), FOREIGN KEY(stream_id) REFERENCES streams(uid) ON UPDATE CASCADE ON DELETE CASCADE )")
	_, _ = db.Exec("CREATE INDEX index_stream_history_stream_id ON stream_history (stream_id)")
	_, _ = db.Exec("CREATE TABLE stream_state (stream_id INTEGER NOT NULL, progress_time INTEGER NOT NULL, PRIMARY KEY(stream_id), FOREIGN KEY(stream_id) REFERENCES streams(uid) ON UPDATE CASCADE ON DELETE CASCADE )")
	_, _ = db.Exec("CREATE TABLE subscriptions (uid INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, service_id INTEGER NOT NULL, url TEXT, name TEXT, avatar_url TEXT, subscriber_count INTEGER, description TEXT, notification_mode INTEGER NOT NULL)")
	_, _ = db.Exec("CREATE UNIQUE INDEX index_subscriptions_service_id_url ON subscriptions (service_id, url)")

	// insert streams
	log.Printf("Inserting %v streams into database file: %s", len(storage.Likes), storage.Args.Output)
	stmt, err := db.Prepare("INSERT INTO streams(uid, service_id, url, title, stream_type, duration, uploader, uploader_url, thumbnail_url, view_count, textual_upload_date, upload_date, is_upload_date_approximation) values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	tools.Errors(err, 1)
	defer stmt.Close()

	for _, stream := range storage.Likes {
		_, err = stmt.Exec(stream.Uid, stream.ServiceId, stream.Url, stream.Title, stream.StreamType, stream.Duration, stream.Uploader, stream.UploaderUrl, stream.ThumbnailUrl, stream.ViewCount, stream.TextualUploadDate, stream.UploadDate, stream.IsUploadDateApproximation)
		tools.Errors(err, 1)
	}

	// insert playlists
	log.Printf("Inserting %v playlists into database file: %s", len(storage.Playlists), storage.Args.Output)
	stmt, err = db.Prepare("INSERT INTO playlists(uid, name, is_thumbnail_permanent, thumbnail_stream_id) values(?,?,?,?)")
	tools.Errors(err, 1)
	defer stmt.Close()

	for _, playlist := range storage.Playlists {
		_, err = stmt.Exec(playlist.Uid, playlist.Name, playlist.IsThrumbnailPermanent, playlist.ThumbnailStreamId)
		tools.Errors(err, 1)
	}

	// insert playlist_stream_joins
	log.Printf("Inserting %v playlist_stream_joins into database file: %s", len(storage.PlaylistStreamJoins), storage.Args.Output)
	stmt, err = db.Prepare("INSERT INTO playlist_stream_join(playlist_id, stream_id, join_index) values(?,?,?)")
	tools.Errors(err, 1)
	defer stmt.Close()

	for _, playlistStreamJoin := range storage.PlaylistStreamJoins {
		_, err = stmt.Exec(playlistStreamJoin.PlaylistId, playlistStreamJoin.StreamId, playlistStreamJoin.JoinIndex)
		tools.Errors(err, 1)
	}

}
