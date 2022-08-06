package main

import (
	"fmt"
)

func (p *Plugin) BackgroundJob() {
	// configuration := p.getConfiguration()
	// if configuration.disabled {
	// 	return
	// }
	fmt.Println("テスト")
	// for _, channelID := range configuration.demoChannelIDs {
	// 	_, err := p.API.CreatePost(&model.Post{
	// 		UserId:    p.botID,
	// 		ChannelId: channelID,
	// 		Message:   "Background job executed",
	// 	})
	// 	if err != nil {
	// 		p.API.LogError(
	// 			"Failed to post BackgroundJob message",
	// 			"channel_id", channelID,
	// 			"error", err.Error(),
	// 		)
	// 	}
	// }
}
