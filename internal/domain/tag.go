package domain

import (
	"context"
	"github.com/bancodobrasil/stop-analyzing-api/internal/db"
)

//CreateTag .
func (s *Service) CreateTag(tagName string) (db.TagModel, error) {
	ctx := context.Background()
	return s.client.Tag.CreateOne(
		db.Tag.Text.Set(tagName),
	).Exec(ctx)
}

//GetAllTags .
func (s *Service) GetAllTags() ([]db.TagModel, error) {
	ctx := context.Background()
	tags, err := s.client.Tag.FindMany().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

//FetchTags searchs for tags using theirs unique name
func (s *Service) FetchTags(names []string) ([]db.TagModel, error) {
	ctx := context.Background()
	return s.client.Tag.FindMany(
		db.Tag.Text.In(names),
	).Exec(ctx)
}

//FetchOrCreateTags searchs for tags (and creates them, if nedeed) using theirs unique name
func (s *Service) FetchOrCreateTags(names []string) (map[string]db.TagModel, error) {
	tags, err := s.FetchTags(names)
	if err != nil {
		return nil, err
	}

	//Inserting into a map to reduce sort and search overhead
	nTags := make(map[string]db.TagModel, len(tags))
	for _, tag := range tags {
		nTags[tag.Text] = tag
	}

	for _, name := range names {
		_, ok := nTags[name]
		if !ok {
			newTag, err := s.CreateTag(name)
			if err != nil {
				return nil, err
			}
			nTags[name] = newTag
		}
	}

	return nTags, nil
}

//DeleteTag remove an existing database tag
func (s *Service) DeleteTag(name string) error {
	_, err := s.client.Tag.FindOne(
		db.Tag.Text.Equals(name),
	).Delete().Exec(context.Background())

	return err
}

//DropAllTags removes all existing database tags
func (s *Service) DropAllTags() (int, error) {
	ctx := context.Background()
	return s.client.Tag.FindMany().Delete().Exec(ctx)
}
