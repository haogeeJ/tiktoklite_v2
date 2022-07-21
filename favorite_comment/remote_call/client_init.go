package remote_call

import (
	"TikTokLite_v2/favorite_comment/remote_call/call_user_follow"
	"TikTokLite_v2/favorite_comment/remote_call/call_video"
)

func Init() {
	call_video.Init()
	call_user_follow.Init()
}
