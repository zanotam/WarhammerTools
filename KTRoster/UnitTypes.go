package KTRoster

//Unit is the base type for models with only the base stats
type Unit struct {
	Name        string
	Points      int
	Move        int
	WS          int
	BS          int
	Strength    int
	Attacks     int
	Wounds      int
	Leadership  int
	Commander   bool
	description string
}

//Weapon is the type for Wargear for Models
type Weapon struct {
	Name        string
	Distance    string
	Style       string
	Strength    string
	AP          int
	Description string
}

//UnitLevel is used to handle commander levels primarily.
type UnitLevel struct {
	Name   string
	Level  int
	Points int
}

//PsychicPower is used to hold psychic powers like psybolt and what not.
type PsychicPower struct {
	Category    string
	WarpCharge  int
	Description string
}

//Skill is used for model's skills
type Skill struct {
	Category string
	Tier     int
	Name     string
}

//CommanderTraits is used to represent the limited number of universal commander traits available
type CommanderTraits struct {
	Name   string
	Points int
	Level4 bool
}
