package controllers

import (
	"fmt"
	"myapp/database"
	"myapp/models"
	"myapp/utils"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myapp/workers"
)

func CreateOrder(c *gin.Context) {

	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	email := order.User


	//verifying user detail who is trying to create orders
	if utils.UserAuth(email) {
		collection := database.GetCollection("orders")
		order.ID = primitive.NewObjectID()
		_, err := collection.InsertOne(ctx, order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
			return
		}

		c.JSON(http.StatusCreated, order)
		fmt.Println(order)
	}

}

func BulkOrder(c *gin.Context) {
	
	workers.TaskQueue<-func(){
		var orders models.Order

		if err := c.ShouldBindJSON(&orders); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx := c.Request.Context()

		collection := database.GetCollection("orders")
		orders.ID = primitive.NewObjectID()
		_, err := collection.InsertOne(ctx, orders)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
			return
		}

		c.JSON(http.StatusCreated, orders)
		fmt.Println(orders)
	}
}

func GetOrder(c *gin.Context) {
	ctx := c.Request.Context()

	collection := database.GetCollection("orders")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch orders"})
		return
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	for cursor.Next(ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not decode orders"})
			return
		}
		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, orders)

}

func UpdateOrder(c *gin.Context) {
	var updatedOrder models.Order
	if err := c.ShouldBindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

}

func DeleteOrder(c *gin.Context) {
	idParam := c.Query("id")
	_, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

}
