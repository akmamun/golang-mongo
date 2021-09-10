package controllers

import (
	"net/http"
	"time"

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
	var todos models.Todo
	if err := ctx.BindJSON(&todos); err != nil {
		ctx.JSON(http.StatusBadRequest,
			gin.H{"success": false, "data": nil, "message": err.Error()})
		return
	}

	validationErr := validate.Struct(todos)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	todos.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	todos.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// todo := models.Todo{Title: todos.Title, Body: todos.Body}
	
	todoCollection.InsertOne(ctx, &todos)
	ctx.JSON(http.StatusOK,
		gin.H{"success": true, "data": nil, "message": "successfully created"})

}
