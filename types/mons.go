package types

type Gender int
type Nature int
type Variant int
type Generation int
type Specimen int

type GeneralMon struct {
	Name             string       `bson:"name" json:"name"`
	Sprites          string       `bson:"sprites" json:"sprites"`
	StatsBaseTable   StatsTable   `bson:"stats_base_table" json:"stats_base_table"`
	AttacksBaseTable AttacksTable `bson:"attacks_base_table" json:"attacks_base_table"`
	Type             int          `bson:"type" json:"type"`
}

type ParticularMon struct {
	Name    string   `bson:"name" json:"name"`
	Exp     int      `bson:"exp" json:"exp"`
	Level   int      `bson:"level" json:"level"`
	Stats   Stats    `bson:"stats" json:"stats"`
	Attacks []Attack `bson:"attacks" json:"attacks"`
}

type OnChainData struct {
	ERC721ID int    `bson:"erc_721_id" json:"erc_721_id"`
	PartID   string `bson:"part_id" json:"part_id"`
	ADN      ADN    `bson:"adn" json:"adn"`
}

type ADN struct {
	Gender      Gender     `bson:"gender" json:"gender"`
	BirthHeight int32      `bson:"birth_height" json:"birth_height"`
	IV          Stats      `bson:"iv" json:"iv"`
	Nature      Nature     `bson:"nature" json:"nature"`
	Variant     Variant    `bson:"variant" json:"variant"`
	Generation  Generation `bson:"generation" json:"generation"`
	Purity      int        `bson:"purity" json:"purity"`
	Specimen    Specimen   `bson:"specimen" json:"specimen"`
	Other       []int      `bson:"other" json:"other"`
	Nonce       int        `bson:"nonce" json:"nonce"`
}

type StatsTable map[int]Stats

type Stats struct {
	H  int `bson:"h" json:"h"`
	A  int `bson:"a" json:"a"`
	SA int `bson:"sa" json:"sa"`
	D  int `bson:"d" json:"d"`
	SD int `bson:"sd" json:"sd"`
}

type AttacksTable map[int][]Attack

type Attack struct {
	Name int `bson:"name" json:"name"`
	Type int `bson:"type" json:"type"`
	F    int `bson:"f" json:"f"`
}
