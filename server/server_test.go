package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func testStatus200(t *testing.T, app *fiber.App, url string, method string) {
	req := httptest.NewRequest(method, url, nil)

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 200, resp.StatusCode, "Status code")
}

func testStatus404(t *testing.T, app *fiber.App, url string, method string) {
	req := httptest.NewRequest(method, url, nil)

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err, "app.Test(req)")
	utils.AssertEqual(t, 404, resp.StatusCode, "Status code")
}

func TestHelloWorld(t *testing.T) {
	app := initServer()
	testStatus200(t, app, "/text", fiber.MethodGet)
	testStatus404(t, app, "/notExist", fiber.MethodGet)
}
