/*
 * User management REST API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zigamedved/user-management-api/controllers/groups"
	"github.com/zigamedved/user-management-api/controllers/users"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	router := gin.Default()
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the GroupsAPI part of the API
	GroupsAPI groups.GroupsAPI
	// Routes for the UsersAPI part of the API
	UsersAPI users.UsersAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{
		{
			"GroupsGet",
			http.MethodGet,
			"/groups",
			handleFunctions.GroupsAPI.GroupsGet,
		},
		{
			"GroupsIdDelete",
			http.MethodDelete,
			"/groups/:id",
			handleFunctions.GroupsAPI.GroupsIdDelete,
		},
		{
			"GroupsIdGet",
			http.MethodGet,
			"/groups/:id",
			handleFunctions.GroupsAPI.GroupsIdGet,
		},
		{
			"GroupsIdPut",
			http.MethodPut,
			"/groups/:id",
			handleFunctions.GroupsAPI.GroupsIdPut,
		},
		{
			"GroupsPost",
			http.MethodPost,
			"/groups",
			handleFunctions.GroupsAPI.GroupsPost,
		},
		{
			"UsersGet",
			http.MethodGet,
			"/users",
			handleFunctions.UsersAPI.UsersGet,
		},
		{
			"UsersIdDelete",
			http.MethodDelete,
			"/users/:id",
			handleFunctions.UsersAPI.UsersIdDelete,
		},
		{
			"UsersIdGet",
			http.MethodGet,
			"/users/:id",
			handleFunctions.UsersAPI.UsersIdGet,
		},
		{
			"UsersIdPut",
			http.MethodPut,
			"/users/:id",
			handleFunctions.UsersAPI.UsersIdPut,
		},
		{
			"UsersPost",
			http.MethodPost,
			"/users",
			handleFunctions.UsersAPI.UsersPost,
		},
	}
}