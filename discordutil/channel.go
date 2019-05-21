package discordutil

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/Bios-Marcel/discordgo"
)

// GetChannelNameForTree returns the string representation of a channel in
// the channel tree on the left.
func GetChannelNameForTree(channel *discordgo.Channel) string {
	if channel.NSFW {
		return channel.Name + " NSFW"
	}

	return channel.Name
}

// SortMessagesByTimestamp sorts all messages in the given array according to
// their creation date.
func SortMessagesByTimestamp(messages []*discordgo.Message) {
	sort.Slice(messages, func(a, b int) bool {
		timeA, parseError := messages[a].Timestamp.Parse()
		if parseError != nil {
			return false
		}

		timeB, parseError := messages[b].Timestamp.Parse()
		if parseError != nil {
			return true
		}

		return timeA.Before(timeB)
	})
}

// GetPrivateChannelName generates a name for a private channel.
func GetPrivateChannelName(channel *discordgo.Channel) string {
	var channelName string
	if channel.Type == discordgo.ChannelTypeDM {
		channelName = channel.Recipients[0].Username
	} else if channel.Type == discordgo.ChannelTypeGroupDM {
		if channel.Name != "" {
			channelName = channel.Name
		} else {
			for index, recipient := range channel.Recipients {
				if index == 0 {
					channelName = recipient.Username
				} else {
					channelName = fmt.Sprintf("%s, %s", channelName, recipient.Username)
				}
			}
		}
	}

	if channelName == "" {
		channelName = "Unnamed"
	}

	return channelName
}

// CompareChannels checks which channel is smaller. Smaller meaning it is the
// one with the more recent message.
func CompareChannels(a, b *discordgo.Channel) bool {
	messageA, parseError := strconv.ParseInt(a.LastMessageID, 10, 64)
	if parseError != nil {
		return false
	}

	messageB, parseError := strconv.ParseInt(b.LastMessageID, 10, 64)
	if parseError != nil {
		return true
	}

	return messageA > messageB
}

// SortPrivateChannels sorts private channels depending on their last message.
func SortPrivateChannels(channels []*discordgo.Channel) {
	sort.Slice(channels, func(a, b int) bool {
		return CompareChannels(channels[a], channels[b])
	})
}
