package news

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/apimachinery/pkg/util/uuid"
	"landing-api/models"
	"landing-api/models/dto"
	"log"
	"time"
)

type News struct {
	ID          string    `bson:"id,omitempty"`
	Title       string    `bson:"title,omitempty"`
	Description string    `bson:"description,omitempty"`
	Image       []byte    `bson:"image,omitempty"`
	PostedAt    time.Time `bson:"postedAt"`
}

func Create(newsCreate *dto.NewsCreate) (*News, error) {
	file, err := newsCreate.Image.Open()
	if err != nil {
		return nil, err
	}

	var imageBytes []byte
	_, err = file.Read(imageBytes)
	if err != nil {
		return nil, err
	}

	news := News{
		ID:          string(uuid.NewUUID()),
		Title:       newsCreate.Title,
		Description: newsCreate.Description,
		Image:       imageBytes,
		PostedAt:    time.Now().UTC(),
	}

	_, err = models.NewsCollection.InsertOne(context.TODO(), news)
	if err != nil {
		err = fmt.Errorf("failed to create news (title: %s). details: %s", newsCreate.Title, err)
		return nil, err
	}

	return &news, nil
}

func Delete(id string) error {
	_, err := models.NewsCollection.DeleteOne(context.TODO(), bson.D{{"id", id}})
	if err != nil {
		err = fmt.Errorf("failed to delete news %s. details: %s", id, err)
		log.Println(err)
		return err
	}
	return nil
}

func Get(id string) (*News, error) {
	var news News
	err := models.NewsCollection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&news)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		err = fmt.Errorf("failed to fetch news with id %s. details: %s", id, err)
		return nil, err
	}

	return &news, err
}

func GetAll() ([]News, error) {
	cursor, err := models.NewsCollection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		err = fmt.Errorf("failed to get all news. details: %s", err)
		log.Println(err)
		return nil, err
	}

	var allNews []News
	for cursor.Next(context.TODO()) {
		var news News

		err = cursor.Decode(&news)
		if err != nil {
			err = fmt.Errorf("failed to decode news getting all news. details: %s", err)
			log.Println(err)
			return nil, err
		}
		allNews = append(allNews, news)
	}

	return allNews, nil
}
