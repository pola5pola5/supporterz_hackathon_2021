// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func upload(c echo.Context) error {
// 	// Read form fields
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")

// 	//-----------
// 	// Read file
// 	//-----------

// 	// Source
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return err
// 	}
// 	src, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close()

// 	// Destination
// 	dst, err := os.Create(file.Filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer dst.Close()

// 	// Copy
// 	if _, err = io.Copy(dst, src); err != nil {
// 		return err
// 	}

// 	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
// }

// func main() {
// 	// Echo instance
// 	e := echo.New()

// 	// Middleware
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	// Routes
// 	e.Static("/", "static")
// 	e.POST("/upload", upload)

// 	// Start server
// 	e.Logger.Fatal(e.Start(":1323"))
// }
