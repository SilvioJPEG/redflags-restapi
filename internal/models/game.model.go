package apiserver

type Game struct {
	ID        string `json:"uid"`
	CreatorId string `json:"creator_id"`
	Judge     string `json:"title"`
	Status    string `json:"status"`
	Turn      string `json:"current_player_id"`
}

func (g *Game) Validate() {
	
}