package API

import (
	"fmt"
)

func(config *Options) GetServerUserLongpoll() *ResponseLongpollServer {
	var (
		body	= config.HeaderExecute("messages.getLongPollServer", "")
		server	= ResponseLongpollServer{}
	)
	server.Bot	= config
	server.Body	= body

	return &server
}

func(config *Options) GetServerGroupLongpoll(groupID int) *ResponseLongpollServer {
	var (
		body	= config.HeaderExecute("groups.getLongPollServer", fmt.Sprintf("group_id=%d", groupID))
		server	= ResponseLongpollServer{}
	)
	server.Bot	= config
	server.Body	= body

	return &server
}