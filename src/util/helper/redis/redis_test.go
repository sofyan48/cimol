package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectRedis(t *testing.T) {
	assert := assert.New(t)
	storeTest := GetConnection()
	assert.Equal(Store, storeTest, "OK")
}

func TestRowsCached(t *testing.T) {
	assert := assert.New(t)
	data := []byte("{\"data\":\"ok\"}")
	dataCache, _ := RowsCached("test", data, 10)
	assert.Equal(data, dataCache, "OK")
}

type JSONer interface {
	JSON(code int)
}

type UserController struct{}

func (ctrl UserController) GetAll(c JSONer) {
	c.JSON(200)
}

type ContextMock struct {
	JSONCalled  bool
	Authoration string
}

func (c *ContextMock) JSON(code int) {
	c.JSONCalled = true
}

// func TestGetToken(t *testing.T) {
// 	assert := assert.New(t)
// 	resp := httptest.NewRecorder()
// 	c, r := gin.CreateTestContext(resp)
// 	gin.SetMode(gin.TestMode)
// 	c := &ContextMock{false}
// 	GetTokenData(c)
// 	// assert.Equal(data,, "OK")
// }
