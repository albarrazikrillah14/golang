package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var app = fiber.New()

func TestRoutingHelloWorld(t *testing.T) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	request := httptest.NewRequest(fiber.MethodGet, "/", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, response.StatusCode, 200)

	responseBody, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	assert.Equal(t, string(responseBody), "Hello World")
}

func TestCtx(t *testing.T) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")

		return c.SendString("Hello " + name)
	})

	request := httptest.NewRequest(fiber.MethodGet, "/hello?name=Zikri", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)

	body, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	bodyString := string(body)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Hello Zikri", bodyString)

	request = httptest.NewRequest(fiber.MethodGet, "/hello", nil)

	response, err = app.Test(request)
	assert.Nil(t, err)

	body, err = io.ReadAll(response.Body)
	assert.Nil(t, err)

	bodyString = string(body)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Hello Guest", bodyString)
}

func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(c *fiber.Ctx) error {
		first := c.Get("firstname")
		last := c.Cookies("lastname")
		return c.SendString("Hello " + first + " " + last)
	})

	request := httptest.NewRequest(fiber.MethodGet, "/request", nil)
	request.Header.Set("firstname", "Albarra")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Zikrillah"})

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	bodyString := string(body)

	assert.Equal(t, "Hello Albarra Zikrillah", bodyString)
}

func TestRouteParameter(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error {
		params := c.AllParams()
		fmt.Println(params["userId"])
		return c.SendString("test")
	})

	request := httptest.NewRequest(fiber.MethodGet, "/users/1/orders/2", nil)

	_, err := app.Test(request)
	assert.Nil(t, err)
}

func TestFormRequest(t *testing.T) {
	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("Hello " + name)
	})

	body := strings.NewReader("name=Zikri")
	request := httptest.NewRequest(fiber.MethodPost, "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	bodyString := string(bytes)

	assert.Equal(t, "Hello Zikri", bodyString)
}

//go:embed source/contoh.txt
var contohFile []byte

func TestFormUpload(t *testing.T) {
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		err = c.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}

		return c.SendString("Upload Success")
	})

	body := new(bytes.Buffer)
	writter := multipart.NewWriter(body)
	file, err := writter.CreateFormFile("file", "contoh.txt")
	assert.Nil(t, err)
	file.Write(contohFile)
	writter.Close()

	request := httptest.NewRequest(fiber.MethodPost, "/upload", body)
	request.Header.Set("Content-Type", writter.FormDataContentType())

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload Success", string(bytes))
}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/authentications", func(c *fiber.Ctx) error {
		request := new(LoginUser)

		err := c.BodyParser(request)
		if err != nil {
			return err
		}

		return c.SendString("Hello " + request.Username)
	})

	requestBody := `{"username":"medomeckz", "password":"P_assword001"}`
	request := httptest.NewRequest(fiber.MethodPost, "/authentications", strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Hello medomeckz", string(bytes))
}
