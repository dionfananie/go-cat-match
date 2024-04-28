package main


import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type user struct {
    Id        int     `json:"id"`
    Username  string  `json:"username"`
}

var users = []user{
    {Id: 546, Username: "John"},
    {Id: 894, Username: "Mary"},
    {Id: 326, Username: "Jane"},
}

func getUsers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, users)
}
func getHome(c *gin.Context){
	c.String(http.StatusOK, "hello world")
}
func main() {
    router := gin.Default()
    router.GET("/users", getUsers);
	
	router.GET("/", getHome);

    router.Run("localhost:8080")
}