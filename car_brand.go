package carsBrandRandomGenerator

type Brand struct {
	Id         int    `json: "id"`
	Name       string `json: "name"`
	ImageBrand []byte `json: "imageBrand"`
}

type Score struct {
	Id           int `json: "id"`
	Player1Score int `json: "player1Score"`
	Player2Score int `json: "player2Score"`
}

type Battle struct {
	Id             int `json: "id"`
	Player1Id      int `json: "player1Id"`
	Player2Id      int `json: "player2Id"`
	ScoreId        int `json: "scoreId"`
	CurrentBrandId int `json: "currentBrandId"`
}
