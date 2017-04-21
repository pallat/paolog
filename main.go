package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type message struct {
	User    string `json:"user"`
	Success string `json:"msg_success"`
	Fail    string `json:"msg_fail"`
}

// Handler
// ddmmyy.hhmmss|TransactionID|SUCCESS/Fail|Message
func logger(c echo.Context) error {
	var msg message
	var s string
	c.Bind(&msg)

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile|log.Ltime|log.Ldate|log.Lmicroseconds)

	if msg.Fail != "" {
		s = fmt.Sprintf("Fail|%s", msg.Fail)
	} else {
		s = fmt.Sprintf("Fail|%s", msg.Success)
	}

	logger.Print(s)
	fmt.Print(&buf)

	return c.String(http.StatusOK, s)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", logger)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
