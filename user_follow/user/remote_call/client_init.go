package remote_call

import (
	"TikTokLite_v2/user_follow/user/remote_call/call_fav_com"
	"TikTokLite_v2/user_follow/user/remote_call/call_video"
)

func Init() {
	call_fav_com.Init()
	call_video.Init()
}
