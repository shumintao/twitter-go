package followers

import "twitter-go/services/gateway/internal/core"

// Routes defines the shape of all the routes for the users package
var Routes = core.Routes{
	core.Route{
		Name:         "FollowUser",
		Method:       "POST",
		Pattern:      "/follow",
		AuthRequired: true,
		HandlerFunc:  FollowUserHandler,
	},
}
