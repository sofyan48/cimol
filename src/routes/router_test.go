package routes

import (
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type UserRouterLoader struct{}

func LoadRouterTestMock() (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	resp := httptest.NewRecorder()
	context, routers := gin.CreateTestContext(resp)

	// routerLoader := &users.UserRouterLoader{}
	// routerLoader.UserRouter(routers)

	return context, routers, resp
}

// func (rLoader *UserRouterLoader) UserRouterTestMock(router *gin.Engine) {
// 	handler := &users.UserController{
// 		UserService: &UserServiceMock{},
// 	}
// 	rLoader.routerDefinition(router, handler)
// }
