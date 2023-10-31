package controller

import (
	"context"
	"fmt"
	"golang-res/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")


func GetOrders() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(context.TODO().bson.M{})
		defer cancel()
		if err!=nil{
			c.JSON(http.StatusInternalServerError, gin.H("error": "error occured while listing order items"))
		}

		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err !=nil{
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders[0])

	}
}

func GetOrder() gin.HandlerFunc{
	
		return func(c *gin.Context) {
			var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
			orderId := c.Params("order_id")
			var order models.Order
	
			err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
			defer cancel()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching orders"})
			}
			c.JSON(http.StatusOK, order)
	
		
	
}
}

func CreateOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		var table models.Table
		var order model.Order



		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


		validationErr := validate.Struct{order}

		if validationErr!=nil{

			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}


		if order.Table_id! = nil {

			err:= tableCollection.FindOne(ctx, bson.M{"table_id" : order.Table_id}).Decode(&table)

			defer cancel()
			if err!=nil{
				msg := fmt.Sprintf("message: Table was not found")

				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})

				return
			}

		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))


		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()
		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil{
			msg := fmt.Sprintf("order item was not created")

			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})

			return
		}


		defer cancel()
		c.JSON(http.StatusOK, result)


		
	}
}



func UpdateOrder() gin.HandlerFunc{
	return func(c *gin.Context){
		var table models.Menu
		var order models.Order


		var updateObj primitive.D

		orderId : c.Params("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Table_id~=nil{

			err := menuCollection.FindOne(ctx, bson.M{"table_id": food.Menu_id}).Decode(&table)
			defer cancel()
			if err!=nil{
				msg := fmt.Sprintf("message: Menu was not found")

				c.JSON(http.StatusInternalServerError, gin.H{"error" : msg})
			}

			updateObj = append(updateObj, bson.E{"menu", order.Table_id})

		}


		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		updateObj = append(updateObj, bson.E{"updated_at", food.Updated_at})


		upsert := true
		filter := bson.M{"order_id": orderId}
		opt := options.UpdateOptions{
			upsert : &upsert,


		}

		result, err := orderCollection.UpdateOne{
			ctx,
			filter,
			bson.D{
				{"$st", updatedObj},
			},

			&opt,

		}

		if err != nil {
			msg := fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})

			return
		}


		defer cancel()
		c.JSON(http.StatusOk, result)

		
	}
}



func OrderItemOrderCreator(order models.Order) String{

	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))


	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(ctx,order )
	defer cancel()

	return order.Order_id



}