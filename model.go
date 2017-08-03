package main

//go:generate neo4jGenerator User
type User struct {
	Id        string `neo4j:"index"`
	FirstName string
	LastName  string
	Email     string `neo4j:"index"`
	Password  string
	Active    bool
	Role      string
}

//go:generate neo4jGenerator Character
type Character struct {
	Id        string `neo4j:"index"`
	FirstName string
	LastName  string
	Race      string
	Class     string
	NPC       bool
}

//go:generate neo4jGenerator Location
type Location struct {
	Id   string `neo4j:"index"`
	Name string
}

type Campaign struct {
	Id          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Setting     string `json:"setting,omitempty"`
}

type Quest struct {
	Id          string `json:"id"`
	CampaignId  string `json:"campaignId,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Entry       string `json:"entry,omitempty"`
}

type Stage struct {
	Id          string `json:"id"`
	QuestId     string `json:"questId,omitempty"`
	Description string `json:"description,omitempty"`
	Entry       string `json:"entry,omitempty"`
}

type JournalEntry struct {
	Id    string `json:"id"`
	Entry string `json:"entry,omitempty"`
	Date  int64  `json:"date"`
}
