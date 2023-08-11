package carsBrandsBattle

import "errors"

type Brand struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name" binding:"required"`
	ImageBrand []byte `json:"imageBrand" db:"imagebrand" binding:"required"`
}

type Battle struct {
	Id             int `json:"id" db:"id"`
	Player1Id      int `json:"player1Id" db:"player1id" binding:"required"`
	Player2Id      int `json:"player2Id" db:"player2id" binding:"required"`
	CurrentBrandId int `json:"currentBrandId" db:"currentbrandid" binding:"required"`
}

type UserBattle struct {
	Id             int `json:"id" db:"id"`
	Player1Id      int `json:"player1Id" db:"player1id" binding:"required"`
	Player2Id      int `json:"player2Id" db:"player2id" binding:"required"`
	CurrentBrandId int `json:"currentBrandId" db:"currentbrandid" binding:"required"`
}

type Score struct {
	Id          int `json:"id" db:"id"`
	UserId      int `json:"userId" db:"userid" binding:"required"`
	BattleId    int `json:"battleId" db:"battleid" binding:"required"`
	PlayerScore int `json:"playerScore" db:"playerscore" binding:"required"`
}

type UpdateBrandInput struct {
	Name       *string `json:"name"`
	ImageBrand *[]byte `json:"imageBrand"`
}

type UpdateBattleInput struct {
	Player1Id      *int `json:"player1Id"`
	Player2Id      *int `json:"player2Id"`
	CurrentBrandId *int `json:"currentBrandId""`
}

type UpdateScoreInput struct {
	UserId      *int `json:"userId"`
	BattleId    *int `json:"battleId"`
	PlayerScore *int `json:"playerScore"`
}

type UserInterfaceData struct {
	BattleId        int    `json:"battleId" db:"battle_id"`
	Player1Username string `json:"player1Username" db:"player1_username"`
	Player2Username string `json:"player2Username" db:"player2_username"`
	PlayerScore1    int    `json:"playerScore1" db:"player1_score"`
	PlayerScore2    int    `json:"playerScore2" db:"player2_score"`
	Score1Id        int    `json:"score1Id" db:"score1_id"`
	Score2Id        int    `json:"score2Id" db:"score2_id"`
}

type UserStatistics struct {
	UserId     int    `json:"userId" db:"id"`
	UserName   string `json:"userName" db:"username"`
	TotalScore int    `json:"totalScore" db:"sum"`
}

func (i UpdateBrandInput) Validate() error {
	if i.Name == nil && i.ImageBrand == nil {
		return errors.New("update structure has no value")
	}
	return nil
}

func (i UpdateBattleInput) Validate() error {
	if i.Player1Id == nil && i.Player2Id == nil && i.CurrentBrandId == nil {
		return errors.New("update structure has no value")
	}
	return nil
}

func (i UpdateScoreInput) Validate() error {
	if i.UserId == nil && i.BattleId == nil && i.PlayerScore == nil {
		return errors.New("update structure has no value")
	}
	return nil
}
