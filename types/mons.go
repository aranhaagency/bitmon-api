package types

type GeneralMon struct {
	Name             string     `bson:"name" json:"name"`
	SpritesUrl       string     `bson:"sprites" json:"sprites"`
	StatsBaseTable   StatsMap   `bson:"stats_base_table" json:"stats_base_table"`
	AttacksBaseTable AttacksMap `bson:"attacks_base_table" json:"attacks_base_table"`
	Nature           Nature     `bson:"nature" json:"nature"`
	Specie           Specie     `bson:"specie" json:"specie"`
	Type             int        `bson:"type" json:"type"`
}

type StatsMap map[int]Stats

type Stats struct {
	H  uint8 `bson:"h" json:"h"`
	A  uint8 `bson:"a" json:"a"`
	SA uint8 `bson:"sa" json:"sa"`
	D  uint8 `bson:"d" json:"d"`
	SD uint8 `bson:"sd" json:"sd"`
}

type AttacksMap map[int][]Attack

type Attack struct {
	Name  string `bson:"name" json:"name"`
	Type  int    `bson:"type" json:"type"`
	Force int    `bson:"force" json:"force"`
}

type NaturesMan map[int]Nature

type Nature struct {
	Name               string `bson:"name" json:"name"`
	StatsModifications Stats  `bson:"stats_modifications" json:"stats_modifications"`
}

type SpeciesMap map[int]Specie

type Specie struct {
	Name string `bson:"name" json:"name"`
}
