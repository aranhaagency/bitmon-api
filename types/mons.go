package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/olympus-protocol/ogen/utils/serializer"
	"io"
)

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
	Nature      Nature     `bson:"nature" json:"nature"`
	Variant     Variant    `bson:"variant" json:"variant"`
	Generation  Generation `bson:"generation" json:"generation"`
	Purity      int        `bson:"purity" json:"purity"`
	Specimen    Specimen   `bson:"specimen" json:"specimen"`
	Nonce       int        `bson:"nonce" json:"nonce"`
	Other       []int      `bson:"other" json:"other"`
	IV          Stats      `bson:"iv" json:"iv"`
}

func (adn *ADN) GetID() (string, error) {
	buf := bytes.NewBuffer([]byte{})
	err := adn.Serialize(buf)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(buf.Bytes())
	lastBytes := hash[len(hash)-4:]
	return hex.EncodeToString(lastBytes), nil
}

func (adn *ADN) Serialize(w io.Writer) error {
	err := serializer.WriteElements(w, adn.Gender, adn.BirthHeight, adn.Nature, adn.Variant, adn.Generation, adn.Purity, adn.Specimen, adn.Nonce)
	if err != nil {
		return err
	}
	err = serializer.WriteVarInt(w, uint64(len(adn.Other)))
	if err != nil {
		return err
	}
	for _, other := range adn.Other {
		err = serializer.WriteElement(w, other)
		if err != nil {
			return err
		}
	}
	err = adn.IV.Serialize(w)
	if err != nil {
		return err
	}
	return nil
}

func (adn *ADN) Deserialize(r io.Reader) error {
	err := serializer.ReadElements(r, &adn.Gender, &adn.BirthHeight, &adn.Nature, &adn.Variant, &adn.Generation, &adn.Purity, &adn.Specimen, &adn.Nonce)
	if err != nil {
		return err
	}
	count, err := serializer.ReadVarInt(r)
	if err != nil {
		return err
	}
	adn.Other = make([]int, count)
	for i := uint64(0); i < count; i++ {
		var element int
		err = serializer.ReadElement(r, &element)
		if err != nil {
			return err
		}
		adn.Other = append(adn.Other, element)
	}
	err = adn.IV.Derialize(r)
	if err != nil {
		return err
	}
	return nil
}

type StatsTable map[int]Stats

type Stats struct {
	H  int `bson:"h" json:"h"`
	A  int `bson:"a" json:"a"`
	SA int `bson:"sa" json:"sa"`
	D  int `bson:"d" json:"d"`
	SD int `bson:"sd" json:"sd"`
}

func (s *Stats) Serialize(w io.Writer) error {
	err := serializer.WriteElements(w, s.H, s.A, s.SA, s.D, s.SD)
	if err != nil {
		return err
	}
	return nil
}

func (s *Stats) Derialize(r io.Reader) error {
	err := serializer.ReadElements(r, &s.H, &s.A, &s.SA, &s.D, &s.SD)
	if err != nil {
		return err
	}
	return nil
}

type AttacksTable map[int][]Attack

type Attack struct {
	Name int `bson:"name" json:"name"`
	Type int `bson:"type" json:"type"`
	F    int `bson:"f" json:"f"`
}
