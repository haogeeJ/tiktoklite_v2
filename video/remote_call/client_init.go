package remote_call

import (
	"TikTokLite_v2/video/remote_call/call_fav_com"
	"TikTokLite_v2/video/remote_call/call_user_follow"
)

func Init() {
	call_fav_com.Init()
	call_user_follow.Init()
	//call_video.Init()
}
