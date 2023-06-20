package main

import (
	"math/rand"
)

// Game represents the game state and data.
type Game struct {
	TankID             string
	Objects            map[string]map[string]interface{}
	Width, Height      float64
	CurrentTurnMessage map[string]interface{}
}

// NewGame creates a new instance of Game and initializes its properties.
func NewGame() *Game {
	game := &Game{
		Objects: make(map[string]map[string]interface{}),
	}

	// Read the tank ID message
	tankIDMessage := ReadMessage()
	game.TankID = tankIDMessage.(map[string]interface{})["message"].(map[string]interface{})["your-tank-id"].(string)

	// Read messages until END_INIT_SIGNAL is received
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

	// Find the boundaries and determine the map size
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

// ReadNextTurnData reads the next turn data from the game server.
// It updates the game state based on the received message.
// Returns true if the game continues, false if the end game signal is received.
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

// RespondToTurn is the logic for responding to the current turn.
// Modify this method with your own logic for responding to the game.
// This example shoots randomly every turn.
func (g *Game) RespondToTurn() {
	shootAngle := rand.Float64() * 360
	message := map[string]interface{}{
		"shoot": shootAngle,
	}

	PostMessage(message)
}
