package main

import (
	"context"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"go-enm/routes"
	"github.com/jackc/pgx/v4"
)

func main(){
	conn, err := connectDB()

	if(err != nil){
		return 
	}
	router := gin.Default()

	router.Use(DbMiddleware(*conn));
	usersGroup := router.Group("users")
	{
		usersGroup.POST("register",routes.UserRegister)
	}

	router.Run(); 
}


func connectDB() (c *pgx.Conn, err error){
	conn, err := pgx.Connect(context.Background(),"postgresql://postgres:Password@2020@localhost:5432/enmgo")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	
	pingErr := conn.Ping(context.Background())

	if(pingErr != nil){
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	
	return conn,err
}

func DbMiddleware(conn pgx.Conn) gin.HandlerFunc{
	return func(c *gin.Context){
		//
		c.Set("db", conn)
		c.Next();
	}
}