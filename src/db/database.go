package db

import (
	"context"

	"github.com/sirupsen/logrus"
)

//DatabasePrisma .
type DatabasePrisma struct {
	client *PrismaClient
}

//Connect .
func Connect() (DatabasePrisma, error) {
	database := DatabasePrisma{}
	database.client = NewClient()
	err := database.client.Connect()
	if err != nil {
		logrus.Errorf("Error at connect to Client: %s", err)
		return database, err
	}
	logrus.Infof("Database sucessfully connected!")
	return database, nil
}

//Disconnect .
func (d *DatabasePrisma) Disconnect() {
	logrus.Info("Disconnecting DB Client")
	err := d.client.Disconnect()
	if err != nil {
		panic(err)
	}
}

//CreateImage .
func (d *DatabasePrisma) CreateImage(imageName string, filePath string, tags []string) error {
	ctx := context.Background()
	_, err := d.client.Image.CreateOne(
		Image.Name.Set(imageName),
		Image.FilePath.Set(filePath),
		Image.Tags.Link(
			Tag.TagName.In(tags),
		),
	).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

//CreateTag .
func (d *DatabasePrisma) CreateTag(tagName string) error {
	ctx := context.Background()
	_, err := d.client.Tag.CreateOne(
		Tag.TagName.Set(tagName),
	).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

//GetAllTags .
func (d *DatabasePrisma) GetAllTags() ([]TagModel, error) {
	ctx := context.Background()
	tags, err := d.client.Tag.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
