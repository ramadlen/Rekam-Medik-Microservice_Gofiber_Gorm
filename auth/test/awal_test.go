package test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// var app = fiber.New()
func TestMain(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	request:= httptest.NewRequest("GET", "/", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err:= io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello World", string(bytes))
	

	//post Request
	
}
// func TestForm(t *testing.T) {
// 	app:= fiber.New()
// 	app.Post("/post", func(c *fiber.Ctx) error {

// 		name := c.FormValue("name")
// 		return c.SendString("Hello" + name)
// 	})
// 	body := strings.NewReader("name=Iqbal")
// 	request := httptest.NewRequest("POST", "/post", body)
// 	request.Header.Set("Content-Type", "application/x-www-form-url-urlencoded")
// 	response, err := app.Test(request)
// 	assert.Nil(t, err)
// 	bytes, err := io.ReadAll(response.Body)
// 	assert.Nil(t, err)
// 	assert.Equal(t, "Hello Iqbal", string(bytes))
// }

