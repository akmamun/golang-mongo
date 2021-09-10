package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/neuron/src/database"
	"github.com/neuron/src/models"
)

//get todoCollection
var todoCollection *mongo.Collection = database.OpenCollection(database.Client, "todos")

var validate = validator.New()

func CreateTodo(ctx *gin.Context) {
	var todo models.Todo
	if err := ctx.BindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(todo)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

}
