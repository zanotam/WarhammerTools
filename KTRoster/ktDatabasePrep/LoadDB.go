package ktDatabasePrep

import (
	"context"

	//	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"strings"
)

//LoadKTDB Loads the killteam data
func LoadKTDB() error {
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
	//	ctx, stop := context.WithCancel(context.Background())
	//	defer stop()
	//	if err := db.PingContext(ctx); err != nil {
	//		log.Fatal("unable to connect to database", err)
	//	}
	defer db.Close()
	//call functions to upload with appropriate filenames
	unitsFile := "Units.tsv"
	fmt.Print("loading units \n")
	unitErr := loadUnits(unitsFile, db)
	if unitErr != nil {
		return unitErr
	}
	fmt.Print("loading Weapons \n")
	weaponsFile := "Weapons.tsv"
	weaponErr := loadWeapons(weaponsFile, db)
	if weaponErr != nil {
		return weaponErr
	}
	fmt.Print("loading skills \n")
	skillsFile := "Skills.tsv"
	skillErr := loadSkills(skillsFile, db)
	if skillErr != nil {
		return skillErr
	}
	fmt.Print("loading psychic \n")
	psychicFile := "Psychic.tsv"
	psychicErr := loadPsychicPowers(psychicFile, db)
	if psychicErr != nil {
		return psychicErr
	}
	return nil
}

func convertToRows(fileName string) ([]string, error) {
	rawText, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	text := string(rawText)
	rows := strings.Split(text, "\r\n")
	return rows, nil
}
func commitTransaction(command string, db *sql.DB) error {
	tx, txErr := db.Begin()
	if txErr != nil {
		return txErr
	}
	_, _ = tx.Exec(command)
	commitErr := tx.Commit()
	if commitErr != nil {
		return commitErr
	}
	return nil
}

func loadUnits(fileName string, db *sql.DB) error {
	//read-in file and prepare text
	rows, fileErr := convertToRows(fileName)
	if fileErr != nil {
		fmt.Print("Converting units file error. \n")
		return fileErr
	}
	/// create table if missing
	var creation string
	creation = `CREATE TABLE IF NOT EXISTS units (
						name VARCHAR(100) NOT NULL PRIMARY KEY,
						points INT NOT NULL,
						move INT, ws INT, bs INT, strength INT, 
						toughness INT NOT NULL, wounds INT NOT NULL, attacks INT NOT NULL, leadership INT NOT NULL,
						save INT NOT NULL,
						description TEXT,
						commander BOOL NOT NULL);`
	commitTableErr := commitTransaction(creation, db)
	if commitTableErr != nil {
		fmt.Print("Create units table errored out. \n")
		return commitTableErr
	}
	//parse lines with a for statement then load lines into the database
	for i, row := range rows {
		if i != (len(rows) - 1) {

			fields := strings.Split(row, "\t")
			insert := `INSERT INTO units (name, points, move, ws, bs, strength, toughness, wounds, attacks, leadership, save, commander, description) VALUES (`
			//add the fields to the insert
			for j, value := range fields {
				if j < (len(fields) - 1) {
					if j == 0 { //name field
						insert = insert + `"` + value + `"`
					} else {
						insert += value
					}
					insert += ", "
				} else { //description field
					//I think this is messing up on null values?
					// 				insert = insert + `"` + value + `"`
					insert += value
					insert += `);`
				}
			}
			//the row insertion transaction stuff
			//		_,_ = fmt.Printf("the row insert looks like: %v \n", insert)
			commitRowErr := commitTransaction(insert, db)
			commitRowErr2 := fmt.Errorf("the row insert looks like: %v", insert)
			if commitRowErr != nil {
				fmt.Print("Units row insertion errored out. \n")
				return commitRowErr2
			}
		}
	}
	return nil
}
func loadWeapons(fileName string, db *sql.DB) error {
	//read-in file and prepare text
	rows, err := convertToRows(fileName)
	if err != nil {
		fmt.Print("converting weapons file error. \n")
		return err
	}
	// create table if missing
	var creation string
	creation = `CREATE TABLE IF NOT EXISTS weapons (
				name VARCHAR(100) NOT NULL PRIMARY KEY, distance Varchar(5), style VARCHAR(12),
				strength VARCHAR(6), ap INT, damage VARCHAR(6), description TEXT, points INT NOT NULL);`
	commitTableErr := commitTransaction(creation, db)
	if commitTableErr != nil {
		fmt.Print("Weapons table errored out. \n")
		return commitTableErr
	}
	//parse lines with a for statement then load lines into the database
	for i, row := range rows {
		if i != (len(rows) - 1) {
			fields := strings.Split(row, "\t")
			insert := `INSERT INTO weapons (name, distance, style, strength, ap, damage, description, points) VALUES (`
			//add the fields to the insert
			for index, value := range fields {
				if index < (len(fields) - 1) {
					if (index == 0) || (index == 1) || (index == 2) || (index == 3) || (index == (len(fields) - 2)) || (index == (len(fields) - 3)) { //name field
						insert = insert + `"` + value + `"`
					} else {
						insert += value
					}
					insert += ", "
				} else { //description field
					insert += value
					insert += ");"
				}
			}
			//the row insertion transaction stuff
			commitRowErr := commitTransaction(insert, db)
			commitRowErr2 := fmt.Errorf("the row insert looks like: %v", insert)
			if commitRowErr != nil {
				fmt.Print("Weapons row insertion errored out. \n")
				return commitRowErr2
			}
		}
	}
	return nil
}
func loadSkills(fileName string, db *sql.DB) error {
	rows, err := convertToRows(fileName)
	if err != nil {
		fmt.Print("converting skills file error. \n")
		return err
	}
	// create table if missing
	var creation string
	creation = `CREATE TABLE IF NOT EXISTS skills (
				name VARCHAR(100) NOT NULL PRIMARY KEY,
				commander BOOL NOT NULL,
				type VARCHAR(100) NOT NULL, level INT NOT NULL, description TEXT NOT NULL);`
	commitTableErr := commitTransaction(creation, db)
	if commitTableErr != nil {
		fmt.Print("Skills table errored out. \n")
		return commitTableErr
	}
	//parse lines with a for statement then load lines into the database
	for i, row := range rows {
		if i != (len(rows) - 1) {
			fields := strings.Split(row, "\t")
			insert := `INSERT INTO skills (type, level, name, commander, description) VALUES (`
			//add the fields to the insert
			for index, value := range fields {
				if index < (len(fields) - 1) {
					if index == 0 { //type field
						insert = insert + `"` + value + `"`
					} else {
						if index == 2 { //name field
							insert = insert + `"` + value + `"`
						} else {
							insert += value
						}
					}
					insert += ", "
				} else { //description field
					insert = insert + `"` + value + `"`
					insert += ");"
				}
			}
			//the row insertion transaction stuff
			commitRowErr := commitTransaction(insert, db)
			if commitRowErr != nil {
				fmt.Print("Skills row errored out. \n")
				return commitRowErr
			}
		}
	}
	return nil
}
func loadPsychicPowers(fileName string, db *sql.DB) error {
	rows, err := convertToRows(fileName)
	if err != nil {
		fmt.Print("Converting psychic file error. \n")
		return err
	}
	// create table if missing
	var creation string
	creation = `CREATE TABLE IF NOT EXISTS psychicpowers (
				name VARCHAR(100) NOT NULL PRIMARY KEY,
				list VARCHAR(100) NOT NULL, warpcharge INT NOT NULL, description TEXT NOT NULL);`
	commitTableErr := commitTransaction(creation, db)
	if commitTableErr != nil {
		fmt.Print("Psychic table errored out. \n")
		return commitTableErr
	}
	//parse lines with a for statement then load lines into the database
	for i, row := range rows {
		if i != (len(rows) - 1) {
			fields := strings.Split(row, "\t")
			insert := `INSERT INTO psychicpowers (list, name, warpcharge, description) VALUES (`
			//add the fields to the insert
			for index, value := range fields {
				if index == 2 {
					insert += value
					insert += ", "
				} else {
					insert = insert + `"` + value + `"`
					if index == (len(fields) - 1) {
						insert += ");"
					} else {
						insert += ", "
					}
				}
			}
			//the row insertion transaction stuff
			commitRowErr := commitTransaction(insert, db)
			if commitRowErr != nil {
				fmt.Print("Psychic row errored out. \n")
				return commitRowErr
			}
		}
	}
	return nil
}
