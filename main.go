package main

import (
	"main/db/moke"
	"main/service"

	"github.com/gin-gonic/gin"
)

func main() {
	//env
	r := gin.Default()
	db := moke.New()
	s := service.New(db)

	//routes
	r.GET("/snkrs", s.GetAllSneakers)
	r.GET("/snkrs/${id}", s.GetSneakers)
	//r.GET("/snkrs/${id}?name=&brand=&price=", s.GetUser)
	/* r.GET("/collections", s.GetAllCollection)
	r.GET("/collections/${id_collection}", s.GetCollection) */
	r.POST("/snkrs", s.CreateSneakers)
	r.DELETE("/snkrs/${id}", s.DeleteSneakers)
	/* r.POST("/collections", s.CreateCollection)
	r.POST("/collections/${id_collection}/snkrs/${id_sneakers}", s.AddSneakers)
	r.DELETE("/collections/${id_collection}/snkrs/${id_sneakers}", s.DeleteSneakers)*/

	//r.POST("/login", s.Login)

	//port
	r.Run(":8081")
}
