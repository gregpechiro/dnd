package main

import (
	"log"
	"net/http"

	"github.com/cagnosolutions/web"
)

var adminHome = web.Route{"GET", "/admin", func(w http.ResponseWriter, r *http.Request) {

	characters, err := GetAllCharacter()
	if err != nil {
		log.Printf("\nadminRoutes.go >> GetAllCharacter() >> %v\n", err)
	}

	locations, err := GetAllLocation()
	if err != nil {
		log.Printf("\nadminRoutes.go >> GetAllLocation() >> %v\n", err)
	}

	tmpl.Render(w, r, "admin.tmpl", web.Model{
		"characters": characters,
		"locations":  locations,
	})
}}

var adminCharacterEdit = web.Route{"GET", "/admin/character/edit/:id", func(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue(":id")
	character, err := GetCharacterById(id)
	if err != nil && id != "new" {
		log.Printf("\nadminRoutes.go >> adminCharacterEdit >> GetCtaracterById() >> %v\n", err)
		web.SetErrorRedirect(w, r, "/admin", "Error finding character")
		return
	}

	locations, err := GetAllLocation()
	if err != nil {
		log.Printf("\nadminRoutes.go >> GetAllLocation() >> %v\n", err)
	}

	tmpl.Render(w, r, "character-edit.tmpl", web.Model{
		"character": character,
		"locations": locations,
	})
}}

var adminCharacterSave = web.Route{"POST", "/admin/character/save", func(w http.ResponseWriter, r *http.Request) {
	var character Character
	r.ParseForm()

	web.FormToStruct(&character, r.Form, "")
	var err error
	if character.Id == "" {
		character.Id = NewId()
		err = AddCharacter(character)
	} else {
		err = UpdateAllCharacterById(character.Id, character)
	}

	if err != nil {
		log.Printf("\nadminRoutes.go >> adminCharactersave >> %v\n", err)
		web.SetErrorRedirect(w, r, "/admin", "Error saving character")
		return
	}

	if lId := r.FormValue("location"); lId != "" {
		conn, err := driver.OpenPool()
		if err != nil {
			log.Printf("\nadminRoutes.go >> adminCharactersave >> %v\n", err)
			web.SetErrorRedirect(w, r, "/admin", "Error saving character Location")
			return
		}
		defer conn.Close()
		if _, err = conn.ExecNeo("MATCH (c:Character{Id: {cId} }), (l:Location{Id: {lId} }) MERGE (c)-[:LOCATED_IN]->(l)", map[string]interface{}{
			"cId": character.Id,
			"lId": lId,
		}); err != nil {
			log.Printf("\nadminRoutes.go >> adminCharactersave >> %v\n", err)
			web.SetErrorRedirect(w, r, "/admin", "Error saving character Location")
			return
		}
	}

	web.SetSuccessRedirect(w, r, "/admin", "Successfully saved character")
	return

}}

var adminLocationEdit = web.Route{"GET", "/admin/location/edit/:id", func(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue(":id")
	location, err := GetLocationById(id)
	if err != nil && id != "new" {
		log.Printf("\nadminRoutes.go >> adminLocationEdit >> GetCtaracterById() >> %v\n", err)
		web.SetErrorRedirect(w, r, "/admin", "Error finding location")
		return
	}

	tmpl.Render(w, r, "location-edit.tmpl", web.Model{
		"location": location,
	})
}}

var adminLocationSave = web.Route{"POST", "/admin/location/save", func(w http.ResponseWriter, r *http.Request) {
	var location Location
	r.ParseForm()

	web.FormToStruct(&location, r.Form, "")
	var err error
	if location.Id == "" {
		location.Id = NewId()
		err = AddLocation(location)
	} else {
		err = UpdateAllLocationById(location.Id, location)
	}

	if err != nil {
		log.Printf("\nadminRoutes.go >> adminLocationsave >> %v\n", err)
		web.SetErrorRedirect(w, r, "/admin", "Error saving location")
		return
	}

	web.SetSuccessRedirect(w, r, "/admin", "Successfully saved location")
	return

}}

var adminCampaignEdit = web.Route{"GET", "/admin/campaign/edit/:id", func(w http.ResponseWriter, r *http.Request) {
	var campaign Campaign
	id := r.FormValue(":id")
	if !db.Get("campaign", id, &campaign) && id != "new" {
		web.SetErrorRedirect(w, r, "/admin", "Error finding campaign")
		return
	}
	tmpl.Render(w, r, "campaign-edit.tmpl", web.Model{
		"campaign": campaign,
	})

}}

var adminCampaignSave = web.Route{"POST", "/admin/campaign/save", func(w http.ResponseWriter, r *http.Request) {
	var campaign Campaign
	r.ParseForm()

	web.FormToStruct(&campaign, r.Form, "")
	if campaign.Id == "" {
		campaign.Id = NewId()
	}

	db.Set("campaign", campaign.Id, campaign)
	web.SetSuccessRedirect(w, r, "/admin", "Successfully saved campaign")
	return

}}

var adminQuestEdit = web.Route{"GET", "/admin/quest/edit/:id", func(w http.ResponseWriter, r *http.Request) {
	var quest Quest
	id := r.FormValue(":id")
	if !db.Get("quest", id, &quest) && id != "new" {
		web.SetErrorRedirect(w, r, "/admin", "Error finding quest")
		return
	}
	tmpl.Render(w, r, "quest-edit.tmpl", web.Model{
		"quest": quest,
	})

}}

var adminQuestSave = web.Route{"POST", "/admin/quest/save", func(w http.ResponseWriter, r *http.Request) {
	var quest Quest
	r.ParseForm()

	web.FormToStruct(&quest, r.Form, "")
	if quest.Id == "" {
		quest.Id = NewId()
	}

	db.Set("quest", quest.Id, quest)
	web.SetSuccessRedirect(w, r, "/admin", "Successfully saved quest")
	return

}}
