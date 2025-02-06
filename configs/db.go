package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MinIO Config for MinIO client and bucket details
type MinioConfig struct {
	Client   *minio.Client
	Bucket   string
	Endpoint string
}

// Initialization for Postgres database connection
func InitDatabase(config *Config) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME,
	)

	db, pgErr := sqlx.ConnectContext(ctx, "postgres", dsn)
	if pgErr != nil {
		return nil, fmt.Errorf("failed to connect to Postgres: %w", pgErr)
	}
	log.Println("✅ Postgres Connected!")

	return db, nil
}

// Initialization for MinIO S3 connection
func InitMinio(config *Config) (*MinioConfig, error) {
	minioClient, minIOErr := minio.New(config.MINIO_ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MINIO_ACCESS_KEY, config.MINIO_SECRET_KEY, ""),
		Secure: false,
	})
	if minIOErr != nil {
		return nil, fmt.Errorf("failed to connect to MinIO: %w", minIOErr)
	}
	log.Println("✅ MinIO S3 Connected!")

	// MinIO Bucket exists
	ctx := context.Background()
	existsBucket, bucketErr := minioClient.BucketExists(ctx, config.MINIO_BUCKET)
	if bucketErr != nil {
		return nil, fmt.Errorf("failed checking bucket existance: %w", bucketErr)
	}
	if !existsBucket {
		err := minioClient.MakeBucket(ctx, config.MINIO_BUCKET, minio.MakeBucketOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed creating new bucket: %w", bucketErr)
		}
	}
	log.Printf("✅ MinIO Bucket '%s' Created!", config.MINIO_BUCKET)

	return &MinioConfig{
		Client:   minioClient,
		Bucket:   config.MINIO_BUCKET,
		Endpoint: config.MINIO_ENDPOINT,
	}, nil
}

func MongoInit(config *Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, mongoErr := mongo.Connect(ctx, options.Client().ApplyURI(config.MONGO_URL))
	if mongoErr != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", mongoErr)
	}
	log.Println("✅ MongoDB Connected!")
	return mongoClient, nil
}
