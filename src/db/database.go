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

//CreateTag .
func (d *DatabasePrisma) CreateTag(tagName string) (TagModel, error) {
	ctx := context.Background()
	return d.client.Tag.CreateOne(
		Tag.Text.Set(tagName),
	).Exec(ctx)
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

//FetchTags searchs for tags using theirs unique name
func (d *DatabasePrisma) FetchTags(names []string) ([]TagModel, error) {
	ctx := context.Background()
	return d.client.Tag.FindMany(
		Tag.Text.In(names),
	).Exec(ctx)
}

//FetchOrCreateTags searchs for tags (and creates them, if nedeed) using theirs unique name
func (d *DatabasePrisma) FetchOrCreateTags(names []string) (map[string]TagModel, error) {
	tags, err := d.FetchTags(names)
	if err != nil {
		return nil, err
	}

	//Inserting into a map to reduce sort and search overhead
	nTags := make(map[string]TagModel, len(tags))
	for _, tag := range tags {
		nTags[tag.Text] = tag
	}

	for _, name := range names {
		_, ok := nTags[name]
		if !ok {
			newTag, err := d.CreateTag(name)
			if err != nil {
				return nil, err
			}
			nTags[name] = newTag
		}
	}

	return nTags, nil
}

//DeleteTag remove an existing database tag
func (d *DatabasePrisma) DeleteTag(name string) error {
	_, err := d.client.Tag.FindOne(
		Tag.Text.Equals(name),
	).Delete().Exec(context.Background())

	return err
}

//DropAllTags removes all existing database tags
func (d *DatabasePrisma) DropAllTags() (int, error) {
	ctx := context.Background()
	return d.client.Tag.FindMany().Delete().Exec(ctx)
}

//DropAllItems removes all existing database items
func (d *DatabasePrisma) DropAllItems() (int, error) {
	ctx := context.Background()
	return d.client.Item.FindMany().Delete().Exec(ctx)
}

//CreateItem creates a new item using provided information and new tags if needed.
//It tries to reuse existing tag with the same name
func (d *DatabasePrisma) CreateItem(title, subtitle, content string, tags []string) (ItemModel, error) {

	mTags, err := d.FetchOrCreateTags(tags)
	if err != nil {
		return ItemModel{}, err
	}

	item, err := d.client.Item.CreateOne(
		Item.Title.Set(title),
		Item.Subtitle.Set(subtitle),
		Item.ContentURL.Set(content),
		Item.Active.Set(true),
	).Exec(context.Background())

	if err != nil {
		return item, err
	}

	//Link tags - TODO: Check if its possible to use/create one method to link them all
	findResult := d.client.Item.FindOne(
		Item.ID.Equals(item.ID),
	)
	for _, mTag := range mTags {
		item, err = findResult.Update(
			Item.Tags.Link(Tag.ID.Equals(mTag.ID)),
		).Exec(context.Background())
	}

	return item, err
}

//FetchItem searchs for an item using its unique id
func (d *DatabasePrisma) FetchItem(id int) (ItemModel, error) {
	ctx := context.Background()
	return d.client.Item.FindOne(
		Item.ID.Equals(id),
	).With(
		Item.Tags.Fetch(),
	).Exec(ctx)
}

//DeleteItem remove an existing database item
func (d *DatabasePrisma) DeleteItem(id int) (ItemModel, error) {
	return d.client.Item.FindOne(
		Item.ID.Equals(id),
	).Delete().Exec(context.Background())
}
