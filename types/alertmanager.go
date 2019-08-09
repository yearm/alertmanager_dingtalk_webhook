package types

import "time"

type AlertMessage struct {
	Version     string  `json:"version"`
	GroupKey    string  `json:"groupKey"`
	Receiver    string  `json:"receiver"`
	Status      string  `json:"status"`
	Alerts      []Alert `json:"alerts"`
	ExternalURL string  `json:"externalUrl"`
}

type Alert struct {
	Labels       *Labels      `json:"labels"`
	Annotations  *Annotations `json:"annotations"`
	StartsAt     time.Time    `json:"startsAt"`
	EndsAt       time.Time    `json:"endsAt"`
	GeneratorURL string       `json:"generatorUrl"`
}

type Labels struct {
	Status    string `json:"status"`
	Alertname string `json:"alertname"`
	Group     string `json:"group"`
	Instance  string `json:"instance"`
	Job       string `json:"job"`
}

type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}
