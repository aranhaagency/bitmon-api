package types

type ReqParams struct {
	ID            string `json:"id"`
	AdventureType string `json:"adventure_type"`
	TicketID      string `json:"ticket_id"`
	TicketProof   string `json:"ticket_proof"`
}
