package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string
	Email string
}

var users = []User{}

func startPAge(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Echo!")
}

func helloGuest(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Guest!")
}

func helloUser(c echo.Context) error {
	return c.String(http.StatusOK, "Hello "+c.Param("name"))
}

func addUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, users)
}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func deleteUser(c echo.Context) error {
	for i, user := range users {
		if user.Name == c.Param("name") {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	return c.JSON(http.StatusOK, users)
}

func main() {
	e := echo.New()
	e.GET("/", startPAge)
	e.GET("/hello", helloGuest)
	e.GET("/hello/:name", helloUser)
	e.GET("/users", getUsers)
	e.POST("/register", addUser)
	e.DELETE("/delete/:name", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
