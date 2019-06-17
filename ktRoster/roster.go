package ktRoster

type RosterUnit struct {
	Base *BasicUnit
	Weapons []*Weapon
	Powers []*PsychicPower
	Skills []*Skill
	Name string
	Points int
}

type Roster []RosterUnit