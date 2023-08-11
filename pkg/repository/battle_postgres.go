package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	carsBrandsBattle "github.com/viktorkaramba/cars-brand-random-generator-app"
)

type BattlePostgres struct {
	db *sqlx.DB
}

func NewBattlePostgres(db *sqlx.DB) *BattlePostgres {
	return &BattlePostgres{db: db}
}

func (r *BattlePostgres) Create(battle carsBrandsBattle.Battle) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	var battleId int
	createBattle := fmt.Sprintf("INSERT INTO %s (player1Id, player2Id, currentBrandId) values ($1, $2, $3) RETURNING id", battlesTable)
	row := r.db.QueryRow(createBattle, battle.Player1Id, battle.Player2Id, battle.CurrentBrandId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	battleId = id
	createScore1 := fmt.Sprintf("INSERT INTO %s (userId, battleId, playerScore) values ($1, $2, $3) RETURNING id", scoresTable)
	row = r.db.QueryRow(createScore1, battle.Player1Id, battleId, 0)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createScore2 := fmt.Sprintf("INSERT INTO %s (userId, battleId, playerScore) values ($1, $2, $3) RETURNING id", scoresTable)
	row = r.db.QueryRow(createScore2, battle.Player2Id, battleId, 0)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return battleId, tx.Commit()
}

func (r *BattlePostgres) GetAll() ([]carsBrandsBattle.Battle, error) {
	var battles []carsBrandsBattle.Battle
	query := fmt.Sprintf("SELECT * FROM %s", battlesTable)
	err := r.db.Select(&battles, query)
	return battles, err
}

func (r *BattlePostgres) GetById(id int) (carsBrandsBattle.Battle, error) {
	var battle carsBrandsBattle.Battle
	query := fmt.Sprintf("SELECT * FROM %s WHERE id= $1", battlesTable)
	err := r.db.Get(&battle, query, id)
	return battle, err
}

func (r *BattlePostgres) Update(id int, battle carsBrandsBattle.UpdateBattleInput) error {
	query := fmt.Sprintf("UPDATE %s SET player1Id=$1, player2Id=$2, CurrentBrandId=$3 WHERE id=$4", battlesTable)
	_, err := r.db.Exec(query, battle.Player1Id, battle.Player2Id, battle.CurrentBrandId, id)
	return err
}

func (r *BattlePostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id= $1", battlesTable)
	_, err := r.db.Exec(query, id)
	return err
}
