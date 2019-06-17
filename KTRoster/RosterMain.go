/*
package KTRoster holds all the info for saving and manipulating the actual rosters in Golang.
It also contains the functions for turning user entry into a proper roster unit, for communication between golang
and JS via JSON for the purpose of modifying rosters in the back-end and loading them in the front-end.
*/
package KTRoster

//These maps are used to translate between JSON and Golang.
var (
	UnitMap    map[string]Unit
	WeaponMap  map[string]Weapon
	PsychicMap map[string]PsychicPower
	SkillMap   map[string]Skill
	TraitMap   map[string]CommanderTraits
)

//RosterPrep prepares roster's maps for JSON to Golang data from the database for usage.
func RosterPrep() {

}
