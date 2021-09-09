package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/neuron/src/database"

	 models "github.com/neuron/src/models"

)

//get todoCollection
var todoCollection *mongo.Collection = database.OpenCollection(database.Client, "todos")

func CreateTodo(ctx *gin.Context) {
	var todo models.Todo

	println(ctx.BindJSON(&todo))
}
