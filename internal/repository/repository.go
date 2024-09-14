package repository

import (
	"event-tracking/config"
	kafkaRepo "event-tracking/internal/repository/kafka"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type Repository interface {
	GetMessageRepository()
}

type postgresRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func InitPostgresRepository(
	cfg *config.Config,
	db *gorm.DB,
) Repository {
	return &postgresRepository{
		cfg: cfg,
		db:  db,
	}
}

func (r *postgresRepository) GetMessageRepository() {
	// return consume.NewMessageService(r.db)
}

// kafka

type KafkaRepository interface {
	GetKafkaReader() kafkaRepo.ReaderRepository
	GetKafkaWriter() kafkaRepo.WriteMsgRepository
}

type kafkaRepository struct {
	cfg    *config.Config
	reader *kafka.Reader
	writer *kafka.Writer
}

func InitKafkaRepository(
	c *config.Config,
	r *kafka.Reader,
	w *kafka.Writer,
) KafkaRepository {
	return &kafkaRepository{
		reader: r,
		cfg:    c,
		writer: w,
	}
}

// other database

type MongoRepository interface{}

type mongoRepository struct {
	cfg *config.Config
	db  *gorm.DB
}

func InitMongoRepository(
	cfg *config.Config,
	db *gorm.DB,
) MongoRepository {
	return &mongoRepository{
		cfg: cfg,
		db:  db,
	}
}

func (r *kafkaRepository) GetKafkaReader() kafkaRepo.ReaderRepository {
	return kafkaRepo.NewReaderRepository(r.reader)
}

func (w *kafkaRepository) GetKafkaWriter() kafkaRepo.WriteMsgRepository {
	return kafkaRepo.NewWriterRepository(w.cfg, w.writer)
}
