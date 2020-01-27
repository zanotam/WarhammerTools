/*
package KTRoster holds all the info for saving and manipulating the actual rosters in Golang.
It also contains the functions for turning user entry into a proper roster unit, for communication between golang
and JS via JSON for the purpose of modifying rosters in the back-end and loading them in the front-end.
*/
package KTRoster

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//These maps are used to translate between JSON and Golang.
var (
	UnitMap    map[string]Unit           = make(map[string]Unit)
	WeaponMap  map[string]Weapon         = make(map[string]Weapon)
	PsychicMap map[string]PsychicPower   = make(map[string]PsychicPower)
	SkillMap   map[string]Skill          = make(map[string]Skill)
	TraitMap   map[string]CommanderTrait = make(map[string]CommanderTrait)
)

//RosterPrep prepares roster's maps for JSON to Golang data from the database for usage.
func init() {
	//set up database
	_, _ = fmt.Printf("opening database \n")
	db, dbErr := sql.Open("mysql", "root:@BlueSyke01@/KillTeamRoster")
	if dbErr != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal("unable to use data source name", dbErr)
	}
	//test database
	fmt.Printf("testing database connection \n")
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("unable to connect to database", err)
	}
	defer db.Close()
	//loading units into map
	unitCommand := `SELECT * FROM units`
	unitRows, _ := commitTransaction(unitCommand, db)
	for unitRows.Next() {
		unit := copyUnitStatsFromRow(unitRows)
		UnitMap[unit.Name] = unit
	}
	//loading weapons into maps
	weaponCommand := `SELECT * FROM weapons`
	weaponRows, _ := commitTransaction(weaponCommand, db)
	for weaponRows.Next() {
		weapon := copyWeaponStatsFromRow(weaponRows)
		WeaponMap[weapon.Name] = weapon
	}
	//loading psychic powers into maps
	psychicCommand := `SELECT * FROM psychicpowers`
	psychicRows, _ := commitTransaction(psychicCommand, db)
	for psychicRows.Next() {
		psychic := copyPsychicStatsFromRow(psychicRows)
		PsychicMap[psychic.Name] = psychic
	}
	//loading skills into maps
	skillCommand := `SELECT * FROM skills`
	skillRows, _ := commitTransaction(skillCommand, db)
	for skillRows.Next() {
		skill := copySkillStatsFromRow(skillRows)
		SkillMap[skill.Name] = skill
	}
	//loading traits into maps
	traitCommand := `SELECT * FROM commandertraits`
	traitRows, _ := commitTransaction(traitCommand, db)
	for traitRows.Next() {
		trait := copyTraitStatsFromRow(traitRows)
		TraitMap[trait.Name] = trait
	}
}

func commitTransaction(command string, db *sql.DB) (row *sql.Rows, err error) {
	tx, txErr := db.Begin()
	if txErr != nil {
		return row, txErr
	}
	row, queryErr := tx.Query(command)
	if queryErr != nil {
		return row, queryErr
	}
	commitErr := tx.Commit()
	if commitErr != nil {
		return row, commitErr
	}
	return
}

func copyUnitStatsFromRow(rows *sql.Rows) (unit Unit) {
	var (
		name        string
		points      int
		move        int
		ws          int
		bs          int
		strength    int
		attacks     int
		wounds      int
		leadership  int
		commander   bool
		description string
	)
	rows.Scan(name, points, move, ws, bs, strength, attacks, wounds, leadership, commander, description)
	unit = Unit{name, points, move, ws, bs, strength, attacks, wounds, leadership, commander, description}
	return
}

func copyWeaponStatsFromRow(rows *sql.Rows) (weapon Weapon) {
	var (
		name        string
		distance    string
		style       string
		strength    string
		ap          int
		damage      string
		description string
	)
	rows.Scan(name, distance, style, strength, ap, damage, description)
	weapon = Weapon{name, distance, style, strength, ap, damage, description}
	return
}

func copyPsychicStatsFromRow(rows *sql.Rows) (psychic PsychicPower) {
	var (
		category    string
		name        string
		warpCharge  int
		description string
	)
	rows.Scan(category, name, warpCharge, description)
	psychic = PsychicPower{category, name, warpCharge, description}
	return
}

func copySkillStatsFromRow(rows *sql.Rows) (skill Skill) {
	var (
		category string
		tier     int
		name     string
	)
	rows.Scan(category, tier, name)
	skill = Skill{category, tier, name}
	return
}

func copyTraitStatsFromRow(rows *sql.Rows) (trait CommanderTrait) {
	var (
		name        string
		points      int
		level4      bool
		description string
	)
	rows.Scan(name, points, level4, description)
	trait = CommanderTrait{name, points, level4, description}
	return
}
