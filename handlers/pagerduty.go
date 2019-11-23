package handlers

import "encoding/json"

type PDReq struct {
	Incident PDIncident `json:"incident"`
}

type PDIncident struct {
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	Service PDService `json:"service"`
}

type PDService struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func SetPagerdutyMessage(req *Request) ([]byte, error) {
	PDService := PDService{ID: req.ServiceId, Type: "service_reference"}
	incident := PDIncident{Type: "incident", Title: req.Title, Service: PDService}
	return json.Marshal(PDReq{Incident: incident})
}
