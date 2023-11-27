package carsBrandsBattleGame

import "errors"

type Brand struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name" binding:"required"`
	ImageBrand string `json:"imageBrand" db:"imagebrand" binding:"required"`
}

type Battle struct {
	Id             *int    `json:"id" db:"id"`
	Player1Id      *int    `json:"player1Id" db:"player1id" binding:"required"`
	Player2Id      *int    `json:"player2Id" db:"player2id" binding:"required"`
	Punishment     *string `json:"punishment" db:"punishment" binding:"required"`
	IsFinished     *bool   `json:"isFinished" db:"isfinished"`
	CurrentBrandId *int    `json:"currentBrandId" db:"currentbrandid" binding:"required"`
}

type Score struct {
	Id          int `json:"id" db:"id"`
	UserId      int `json:"userId" db:"userid" binding:"required"`
	BattleId    int `json:"battleId" db:"battleid" binding:"required"`
	PlayerScore int `json:"playerScore" db:"playerscore" binding:"required"`
}

type Token struct {
	Id         int    `json:"id" db:"id"`
	TokenValue string `json:"tokenValue" db:"tokenvalue" binding:"required"`
	Revoked    bool   `json:"revoked" db:"revoked"`
	UserId     int    `json:"userId" db:"userid"`
}

type UpdateBrandInput struct {
	Name       *string `json:"name"`
	ImageBrand *string `json:"imageBrand"`
}

type UpdateBattleInput struct {
	Player1Id      *int    `json:"player1Id"`
	Player2Id      *int    `json:"player2Id"`
	Punishment     *string `json:"punishment"`
	IsFinished     *bool   `json:"isFinished"`
	CurrentBrandId *int    `json:"currentBrandId"`
}

type UpdateScoreInput struct {
	UserId   *int `json:"userId"`
	BattleId *int `json:"battleId"`
}

type UpdateTokenInput struct {
	TokenValue *string `json:"tokenValue"`
	Revoked    *bool   `json:"revoked"`
	UserId     *int    `json:"userId"`
}

type UserInterfaceData struct {
	BattleId        int    `json:"battleId" db:"battle_id" binding:"required"`
	Player1Username string `json:"player1Username" db:"player1_username" binding:"required"`
	Player2Username string `json:"player2Username" db:"player2_username" binding:"required"`
	PlayerScore1    int    `json:"playerScore1" db:"player1_score" binding:"required"`
	PlayerScore2    int    `json:"playerScore2" db:"player2_score" binding:"required"`
	Score1Id        int    `json:"score1Id" db:"score1_id" binding:"required"`
	Score2Id        int    `json:"score2Id" db:"score2_id" binding:"required"`
	Punishment      string `json:"punishment" db:"brandpunishment" binding:"required"`
	BrandName       string `json:"brandName" db:"brandname" binding:"required"`
}

type UserStatistics struct {
	UserId     int    `json:"userId" db:"id" binding:"required"`
	Username   string `json:"username" db:"username" binding:"required"`
	TotalScore int    `json:"totalScore" db:"sum" binding:"required"`
}

type RefreshTokenInput struct {
	UserId int `json:"userId" binding:"required"`
}

func (i UpdateBrandInput) Validate() error {
	if i.Name == nil && i.ImageBrand == nil {
		return errors.New("update structure has no value")
	}
	return nil
}

func (i UpdateBattleInput) Validate() error {
	if i.Player1Id == nil && i.Player2Id == nil && i.Punishment == nil && i.IsFinished == nil && i.CurrentBrandId == nil {
		return errors.New("update structure has no value")
	}
	return nil
}

func (i UpdateScoreInput) Validate() error {
	if i.UserId == nil && i.BattleId == nil {
		return errors.New("update structure has no value")
	}
	return nil
}

func (i UpdateTokenInput) Validate() error {
	if i.UserId == nil && i.TokenValue == nil && i.Revoked == nil {
		return errors.New("update structure has no value")
	}
	return nil
}
