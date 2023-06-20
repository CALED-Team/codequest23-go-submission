package main

func main() {
	game := NewGame()
	for game.ReadNextTurnData() {
		game.RespondToTurn()
	}
}
