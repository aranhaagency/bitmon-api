package types

type Gender uint32
type Nature uint32
type Variant uint32
type Generation uint32
type Specimen uint32

type GeneralMon struct {
	Name             string     `bson:"name" json:"name"`
	SpritesUrl       string     `bson:"sprites" json:"sprites"`
	StatsBaseTable   StatsMap   `bson:"stats_base_table" json:"stats_base_table"`
	AttacksBaseTable AttacksMap `bson:"attacks_base_table" json:"attacks_base_table"`
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

type NaturesMan map[int]NatureSpecifics

type NatureSpecifics struct {
	Name               string `bson:"name" json:"name"`
	StatsModifications Stats  `bson:"stats_modifications" json:"stats_modifications"`
}

type SpeciesMap map[int]SpeciesInfo

type SpeciesInfo struct {
	Name string `bson:"name" json:"name"`
}
