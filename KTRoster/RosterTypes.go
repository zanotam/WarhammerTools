package KTRoster

//Model Holds all the info for a single unit within a roster
type Model struct {
	Name    string
	Base    *Unit
	Weapons []*Weapon
	Powers  []*PsychicPower
	Exp     int
	Level   *UnitLevel
	Skills  []*Skill
	Traits  []*CommanderTraits
	Points  int
}

//Roster holds all the data for a full roster from the 20 units to things like the player's name.
type Roster struct {
	TeamName   string
	PlayerName string
	Faction    string
	Mission    string
	Background string
	SquadQuirk string
	Resources  *Resource
	Units      [20]Model
}

//Resource is an alias of [4]int that has get and set methods for the 4 types of resources.
type Resource [4]int

func (resources *Resource) Intelligence() int {
	return [4]int(*resources)[0]
}
func (resources *Resource) SetIntelligence(value int) {
	[4]int(*resources)[0] = value
}
func (resources *Resource) Material() int {
	return [4]int(*resources)[1]
}
func (resources *Resource) SetMaterial(value int) {
	[4]int(*resources)[1] = value
}
func (resources *Resource) Morale() int {
	return [4]int(*resources)[2]
}
func (resources *Resource) SetMorale(value int) {
	[4]int(*resources)[2] = value
}
func (resources *Resource) Territory() int {
	return [4]int(*resources)[3]
}
func (resources *Resource) SetTerriotyr(value int) {
	[4]int(*resources)[3] = value
}
