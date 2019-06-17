package ktRoster

//basic unit
//skill
//psychic power(s)
// weapon(s)

type BasicUnit struct {
	Name string
	Points int
	Move int
	WS int
	BS int
	Strength int
	Attacks int
	Wounds int
	Leadership int
}

type Weapon struct {
	Name string
	Distance string
	Style string
	Strength string
	AP int
	Description string
}

type PsychicPower struct {

}

type Skill struct {

}