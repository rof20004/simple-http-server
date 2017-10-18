package main

import "github.com/labstack/echo"
import "os"

func main() {
	e := echo.New()
	e.Static("/", ".")

	var port = os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
}
