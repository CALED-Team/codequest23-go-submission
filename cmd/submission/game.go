package main

import (
	"math/rand"
)

type Game struct {
	TankID             string
	Objects            map[string]map[string]interface{}
	Width, Height      float64
	CurrentTurnMessage map[string]interface{}
}

func NewGame() *Game {
	game := &Game{
		Objects: make(map[string]map[string]interface{}),
	}

	tankIDMessage := ReadMessage()
	game.TankID = tankIDMessage.(map[string]interface{})["message"].(map[string]interface{})["your-tank-id"].(string)

	// read messages until you receive END_INIT_SIGNAL
	nextInitMessage := ReadMessage()
	_, ok := nextInitMessage.(string)
	for !ok {
		objectInfo := nextInitMessage.(map[string]interface{})["message"].(map[string]interface{})["updated_objects"].(map[string]interface{})
		for objectID, objectData := range objectInfo {
			game.Objects[objectID] = objectData.(map[string]interface{})
		}

		nextInitMessage = ReadMessage()
		_, ok = nextInitMessage.(string)
	}

	var boundaries []map[string]interface{}
	for _, gameObject := range game.Objects {
		objectType := int(gameObject["type"].(float64))
		if objectType == int(BOUNDARY) {
			boundaries = append(boundaries, gameObject)
		}
	}

	var biggestX, biggestY float64
	for _, boundary := range boundaries {
		position := ParsePosition(boundary["position"].([]interface{}))
		for _, singlePosition := range position {
			if singlePosition[0] > biggestX {
				biggestX = singlePosition[0]
			}
			if singlePosition[1] > biggestY {
				biggestY = singlePosition[1]
			}
		}
	}

	game.Width = biggestX
	game.Height = biggestY

	return game
}

func (g *Game) ReadNextTurnData() bool {
	rawMessage := ReadMessage()

	if message, ok := rawMessage.(string); ok && message == END_SIGNAL {
		return false
	}

	g.CurrentTurnMessage = rawMessage.(map[string]interface{})

	deletedObjects := g.CurrentTurnMessage["message"].(map[string]interface{})["deleted_objects"].([]interface{})
	for _, deletedObjectID := range deletedObjects {
		delete(g.Objects, deletedObjectID.(string))
	}

	updatedObjects := g.CurrentTurnMessage["message"].(map[string]interface{})["updated_objects"].(map[string]interface{})
	for objectID, objectData := range updatedObjects {
		g.Objects[objectID] = objectData.(map[string]interface{})
	}

	return true
}

func (g *Game) RespondToTurn() {
	// Your logic for responding to the turn goes here.
	// This example just shoots randomly every turn.
	shootAngle := rand.Float64() * 360
	message := map[string]interface{}{
		"shoot": shootAngle,
	}

	PostMessage(message)
}
