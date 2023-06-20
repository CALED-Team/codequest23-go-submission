package main

func main() {
	game := NewGame()

	// Continue reading and responding to turns until the game ends
	for game.ReadNextTurnData() {
		game.RespondToTurn()
	}
}
