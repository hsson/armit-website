package models

type Partner struct {
  Id        int           `json:"id"`
  Name      string        `json:"name"`
  Website   string        `json:"website"`
  LogoUrl   string        `json:"logoUrl"`
}

type Partners []Partner
