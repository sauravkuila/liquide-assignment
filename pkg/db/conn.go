package db

import (
	"context"
	"fmt"
	"liquide-assignment/pkg/config"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func PsqlConnect() (*gorm.DB, error) {

	c := config.GetConfig()

	// Init DB connection string params from config
	dsn := fmt.Sprintf("host=%s port=%d user=%s password='%s' dbname=%s sslmode=%s connect_timeout=%d",
		c.GetString("databases.postgres.host"), c.GetInt("databases.postgres.port"), c.GetString("databases.postgres.user"), c.GetString("databases.postgres.password"), c.GetString("databases.postgres.db"), c.GetString("databases.postgres.sslmode"), c.GetInt("databases.postgres.connect_timeout"))

	// Connect to database
	dbc, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: true, /// Removed defuault transaction used by GORM to faster the query and execution
		Logger:                 logger.Default,
	})
	if err != nil {
		log.Printf("Failed to connect to database. Error: %s, conn: %s", err.Error(), dsn)
		return nil, err
	}

	log.Println("Postgres Database Connected")
	return dbc, nil
}

func RedisConnect() (*redis.Client, error) {
	c := config.GetConfig()

	// Initialize Redis client.
	redisClient := redis.NewClient(&redis.Options{
		Addr: c.GetString("databases.redis.host") + ":" + c.GetString("databases.redis.port"), // Adjust if your Redis is hosted elsewhere.
	})

	//check connection
	pong, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		return nil, err
	}
	log.Println("redis connected. result: ", pong)

	return redisClient, nil
}
