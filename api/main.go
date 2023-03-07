package main

// import (
// 	"api/router"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// func idMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Set("id", uuid.New().String())
// 	}
// }
// func main() {

// 	r := gin.Default()
// 	r.Use(idMiddleware())
// 	router.GetRoutes(r)
// 	r.Run()
// }
