package types

type Gender uint32
type Nature uint32
type Variant uint32
type Generation uint32
type Specimen uint32

type GeneralMon struct {
	Name             string       `bson:"name" json:"name"`
	SpritesUrl       string       `bson:"sprites" json:"sprites"`
	StatsBaseTable   StatsTable   `bson:"stats_base_table" json:"stats_base_table"`
	AttacksBaseTable AttacksTable `bson:"attacks_base_table" json:"attacks_base_table"`
	Type             int          `bson:"type" json:"type"`
}

type StatsTable map[int]Stats

type Stats struct {
	H  uint32 `bson:"h" json:"h"`
	A  uint32 `bson:"a" json:"a"`
	SA uint32 `bson:"sa" json:"sa"`
	D  uint32 `bson:"d" json:"d"`
	SD uint32 `bson:"sd" json:"sd"`
}

type AttacksTable map[int][]Attack

type Attack struct {
	Name int `bson:"name" json:"name"`
	Type int `bson:"type" json:"type"`
	F    int `bson:"f" json:"f"`
}
