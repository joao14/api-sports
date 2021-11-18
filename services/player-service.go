package services

import (
	"github.com/joao14/api-sports.git/model"
)

type PlayerService interface {
	Save(model.Player) model.Player
	FindAll() []model.Player
}

type playerService struct {
	players []model.Player
}

func New() PlayerService {
	return &playerService{}
}

func (service *playerService) Save(player model.Player) model.Player {
	//var playerCollection *mongo.Collection = config.OpenCollection(config.Client, "player")
	//service.players = append(service.players, player)
	//result, insertErr := playerCollection.InsertOne(player)
	return player
}

func (service *playerService) FindAll() []model.Player {
	return service.players
}
