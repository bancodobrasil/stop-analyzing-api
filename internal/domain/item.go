package domain

import (
	"context"

	"github.com/bancodobrasil/stop-analyzing-api/internal/db"
)

//CreateItem creates a new item using provided information and new tags if needes.
//It tries to reuse existing tag with the same name
func (s *Service) CreateItem(title, subtitle, content string, tags []string) (*db.ItemModel, error) {

	mTags, err := s.FetchOrCreateTags(tags)
	if err != nil {
		return &db.ItemModel{}, err
	}

	item, err := s.client.Item.CreateOne(
		db.Item.Title.Set(title),
		db.Item.Subtitle.Set(subtitle),
		db.Item.ContentURL.Set(content),
		db.Item.Active.Set(true),
	).Exec(context.Background())

	if err != nil {
		return item, err
	}

	//Link tags - TODO: Check if its possible to use/create one method to link them all
	findResult := s.client.Item.FindUnique(
		db.Item.ID.Equals(item.ID),
	)
	for _, mTag := range mTags {
		item, err = findResult.Update(
			db.Item.Tags.Link(db.Tag.ID.Equals(mTag.ID)),
		).Exec(context.Background())
	}

	return item, err
}

//FetchItem searchs for an item using its unique id
func (s *Service) FetchItem(id int) (*db.ItemModel, error) {
	ctx := context.Background()
	return s.client.Item.FindUnique(
		db.Item.ID.Equals(id),
	).With(
		db.Item.Tags.Fetch(),
	).Exec(ctx)
}

//DeleteItem remove an existing database item
func (s *Service) DeleteItem(id int) (*db.ItemModel, error) {
	return s.client.Item.FindUnique(
		db.Item.ID.Equals(id),
	).Delete().Exec(context.Background())
}
