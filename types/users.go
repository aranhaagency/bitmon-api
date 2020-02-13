package types

type UserMons struct {
	Acc  string     `json:"acc"`
	Mons []MonsData `json:"mons"`
}

type MonsData struct {
	General    GeneralMon    `json:"general"`
	Particular ParticularMon `json:"particular"`
	OnChain    OnChainData   `json:"on_chain"`
}
