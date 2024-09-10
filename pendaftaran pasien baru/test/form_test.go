package test

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)
func TestForm(t *testing.T) {
	app := fiber.New()
	app.Post("/form", func(ctx *fiber.Ctx) error {

		name := ctx.FormValue("name")
		return ctx.SendString("Hello " + name)
	})
	body := strings.NewReader("name=Iqbal")
	request := httptest.NewRequest("POST", "/form", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Iqbal", string(bytes))
}
func TestQueryError(t *testing.T)  {
	app := fiber.New()
	app.Get("/hello", func (ctx *fiber.Ctx)error  {
		name := ctx.Query("name", "Guest")
		return ctx.SendString("Hello " + name)
	})
	request := httptest.NewRequest("GET", "/hello?name=Iqbal", nil)
	response, err  := app.Test(request)
	assert.Nil(t, err)
	bytes, err  := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello Iqbal", string(bytes))
}