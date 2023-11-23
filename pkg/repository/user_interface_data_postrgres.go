package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
)

type UserInterfaceDataPostgres struct {
	db *sqlx.DB
}

func NewUserInterfaceDataPostgres(db *sqlx.DB) *UserInterfaceDataPostgres {
	return &UserInterfaceDataPostgres{db: db}
}

func (r *UserInterfaceDataPostgres) GetAll(isFinished bool) ([]carsBrandsBattle.UserInterfaceData, error) {
	var userInterfaceData []carsBrandsBattle.UserInterfaceData
	query := fmt.Sprintf(
		"SELECT b.id AS battle_id, b.punishment AS brandpunishment, br.name AS brandname, u1.username AS player1_username, u2.username AS player2_username, "+
			"s1.playerScore AS player1_score, s2.playerScore AS player2_score, s1.id AS score1_id, s2.id AS score2_id "+
			"FROM %s b "+
			"JOIN %s br ON br.id = b.currentbrandid "+
			"JOIN %s u1 ON b.player1Id = u1.id JOIN %s u2 ON b.player2Id = u2.id "+
			"JOIN %s s1 ON b.id = s1.battleId AND b.player1Id = s1.userId "+
			"JOIN %s s2 ON b.id = s2.battleId AND b.player2Id = s2.userId WHERE b.isFinished='%v' ORDER BY b.id ASC ",
		battlesTable, brandsTable, usersTable, usersTable, scoresTable, scoresTable, isFinished)
	err := r.db.Select(&userInterfaceData, query)
	return userInterfaceData, err
}

func (r *UserInterfaceDataPostgres) GetById(battleId int, isFinished bool) (*carsBrandsBattle.UserInterfaceData, error) {
	var userInterfaceData carsBrandsBattle.UserInterfaceData
	query := fmt.Sprintf(
		"SELECT b.id AS battle_id, b.punishment AS brandpunishment, br.name AS brandname, u1.username AS player1_username, u2.username AS player2_username, "+
			"s1.playerScore AS player1_score, s2.playerScore AS player2_score, s1.id AS score1_id, s2.id AS score2_id "+
			"FROM %s b "+
			"JOIN %s br ON br.id = b.currentbrandid "+
			"JOIN %s u1 ON b.player1Id = u1.id JOIN %s u2 ON b.player2Id = u2.id "+
			"JOIN %s s1 ON b.id = s1.battleId AND b.player1Id = s1.userId "+
			"JOIN %s s2 ON b.id = s2.battleId AND b.player2Id = s2.userId WHERE b.isFinished='%v' AND b.id=$1 ORDER BY b.id ASC ",
		battlesTable, brandsTable, usersTable, usersTable, scoresTable, scoresTable, isFinished)
	err := r.db.Get(&userInterfaceData, query, battleId)
	if userInterfaceData == (carsBrandsBattle.UserInterfaceData{}) {
		return &carsBrandsBattle.UserInterfaceData{}, nil
	}
	return &userInterfaceData, err
}
