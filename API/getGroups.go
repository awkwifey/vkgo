package API

import (
	"fmt"
)

func(config *Options) GetGroup(groupID int) *ResponseGroupsByID {
	var (
		body		= config.HeaderExecute("groups.getById", fmt.Sprintf("group_id=%d", groupID))
		groups		= ResponseGroupsByID{}
	)
	groups.Bot		= config
	groups.Body		= body

	return &groups
}

func(config *Options) GetGroupID(groupID int) int {
	groups := config.GetGroup(groupID).Parse()
	for _, group := range groups.Response.Groups {
		return group.ID
	}
	return 0
}