package domain

import (
	"context"

	"github.com/bancodobrasil/stop-analyzing-api/internal/db"
	"github.com/sirupsen/logrus"
)

type Service struct {
	client *db.PrismaClient
}

//Init .
func NewService() (*Service, error) {
	service := Service{}
	service.client = db.NewClient()
	err := service.client.Connect()
	if err != nil {
		logrus.Errorf("Error at connect to Client: %s", err)
		return &service, err
	}
	logrus.Infof("Database sucessfully connected!")
	return &service, nil
}

//Disconnect .
func (d *Service) Disconnect() {
	logrus.Info("Disconnecting DB Client")
	err := d.client.Disconnect()
	if err != nil {
		panic(err)
	}
}

//DropAllItems removes all existing database items
func (d *Service) DropAllItems() (int, error) {
	ctx := context.Background()
	br, err := d.client.Item.FindMany().Delete().Exec(ctx)
	if err != nil {
		return 0, err
	}
	return br.Count, nil
}
