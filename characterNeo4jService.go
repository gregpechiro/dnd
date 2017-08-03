/*
* CODE GENERATED AUTOMATICALLY WITH github.com/gregpechiro/neo4jGenerator
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

var NoCharacterFound = fmt.Errorf("no character found")
var MultipleCharacterFound = fmt.Errorf("multiple character found")

func AddCharacter(character Character) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE (character:Character { Id:{characterId}, FirstName:{characterFirstName}, LastName:{characterLastName}, Race:{characterRace}, Class:{characterClass}, NPC:{characterNPC} })", map[string]interface{}{
		"characterId":        character.Id,
		"characterFirstName": character.FirstName,
		"characterLastName":  character.LastName,
		"characterRace":      character.Race,
		"characterClass":     character.Class,
		"characterNPC":       character.NPC,
	})

	return err
}

func GetAllCharacter() ([]Character, error) {
	var characters []Character
	conn, err := driver.OpenPool()
	if err != nil {
		return characters, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character) RETURN character", nil)
	if err != nil {
		return characters, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return characters, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return characters, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		character := Character{}
		if Id, ok := node.Properties["Id"].(string); ok {
			character.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			character.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			character.LastName = LastName
		}
		if Race, ok := node.Properties["Race"].(string); ok {
			character.Race = Race
		}
		if Class, ok := node.Properties["Class"].(string); ok {
			character.Class = Class
		}
		if NPC, ok := node.Properties["NPC"].(bool); ok {
			character.NPC = NPC
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func IndexCharacterById() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :Character(Id)", nil)

	return err
}

func GetCharacterById(id string) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ Id:{ Id } }) RETURN character", map[string]interface{}{
		"Id": id,
	})
	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetOnlyOneCharacterById(id string) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ Id:{ Id } }) RETURN character", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return character, MultipleCharacterFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetAllCharacterById(id string) ([]Character, error) {
	var characters []Character
	conn, err := driver.OpenPool()
	if err != nil {
		return characters, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ Id:{ Id } }) RETURN character", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return characters, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return characters, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return characters, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		character := Character{}
		if Id, ok := node.Properties["Id"].(string); ok {
			character.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			character.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			character.LastName = LastName
		}
		if Race, ok := node.Properties["Race"].(string); ok {
			character.Race = Race
		}
		if Class, ok := node.Properties["Class"].(string); ok {
			character.Class = Class
		}
		if NPC, ok := node.Properties["NPC"].(bool); ok {
			character.NPC = NPC
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func UpdateAllCharacterById(id string, character Character) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ Id:{ Id } }) SET character += { Id:{characterId}, FirstName:{characterFirstName}, LastName:{characterLastName}, Race:{characterRace}, Class:{characterClass}, NPC:{characterNPC} }", map[string]interface{}{
		"Id":                 id,
		"characterId":        character.Id,
		"characterFirstName": character.FirstName,
		"characterLastName":  character.LastName,
		"characterRace":      character.Race,
		"characterClass":     character.Class,
		"characterNPC":       character.NPC,
	})
	return err
}

func DeleteAllCharacterById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ Id:{ Id }) DETACH DELETE character", map[string]interface{}{
		"Id": id,
	})
	return err
}

func GetCharacterByFirstName(firstName string) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ FirstName:{ FirstName } }) RETURN character", map[string]interface{}{
		"FirstName": firstName,
	})
	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetOnlyOneCharacterByFirstName(firstName string) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ FirstName:{ FirstName } }) RETURN character", map[string]interface{}{
		"FirstName": firstName,
	})

	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return character, MultipleCharacterFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetAllCharacterByFirstName(firstName string) ([]Character, error) {
	var characters []Character
	conn, err := driver.OpenPool()
	if err != nil {
		return characters, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ FirstName:{ FirstName } }) RETURN character", map[string]interface{}{
		"FirstName": firstName,
	})

	if err != nil {
		return characters, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return characters, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return characters, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		character := Character{}
		if Id, ok := node.Properties["Id"].(string); ok {
			character.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			character.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			character.LastName = LastName
		}
		if Race, ok := node.Properties["Race"].(string); ok {
			character.Race = Race
		}
		if Class, ok := node.Properties["Class"].(string); ok {
			character.Class = Class
		}
		if NPC, ok := node.Properties["NPC"].(bool); ok {
			character.NPC = NPC
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func UpdateAllCharacterByFirstName(firstName string, character Character) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ FirstName:{ FirstName } }) SET character += { Id:{characterId}, FirstName:{characterFirstName}, LastName:{characterLastName}, Race:{characterRace}, Class:{characterClass}, NPC:{characterNPC} }", map[string]interface{}{
		"FirstName":          firstName,
		"characterId":        character.Id,
		"characterFirstName": character.FirstName,
		"characterLastName":  character.LastName,
		"characterRace":      character.Race,
		"characterClass":     character.Class,
		"characterNPC":       character.NPC,
	})
	return err
}

func DeleteAllCharacterByFirstName(firstName string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ FirstName:{ FirstName }) DETACH DELETE character", map[string]interface{}{
		"FirstName": firstName,
	})
	return err
}

func GetCharacterByLastName(lastName string) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ LastName:{ LastName } }) RETURN character", map[string]interface{}{
		"LastName": lastName,
	})
	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetOnlyOneCharacterByLastName(lastName string) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ LastName:{ LastName } }) RETURN character", map[string]interface{}{
		"LastName": lastName,
	})

	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return character, MultipleCharacterFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetAllCharacterByLastName(lastName string) ([]Character, error) {
	var characters []Character
	conn, err := driver.OpenPool()
	if err != nil {
		return characters, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ LastName:{ LastName } }) RETURN character", map[string]interface{}{
		"LastName": lastName,
	})

	if err != nil {
		return characters, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return characters, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return characters, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		character := Character{}
		if Id, ok := node.Properties["Id"].(string); ok {
			character.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			character.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			character.LastName = LastName
		}
		if Race, ok := node.Properties["Race"].(string); ok {
			character.Race = Race
		}
		if Class, ok := node.Properties["Class"].(string); ok {
			character.Class = Class
		}
		if NPC, ok := node.Properties["NPC"].(bool); ok {
			character.NPC = NPC
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func UpdateAllCharacterByLastName(lastName string, character Character) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ LastName:{ LastName } }) SET character += { Id:{characterId}, FirstName:{characterFirstName}, LastName:{characterLastName}, Race:{characterRace}, Class:{characterClass}, NPC:{characterNPC} }", map[string]interface{}{
		"LastName":           lastName,
		"characterId":        character.Id,
		"characterFirstName": character.FirstName,
		"characterLastName":  character.LastName,
		"characterRace":      character.Race,
		"characterClass":     character.Class,
		"characterNPC":       character.NPC,
	})
	return err
}

func DeleteAllCharacterByLastName(lastName string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ LastName:{ LastName }) DETACH DELETE character", map[string]interface{}{
		"LastName": lastName,
	})
	return err
}

func GetCharacterByRace(race string) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ Race:{ Race } }) RETURN character", map[string]interface{}{
		"Race": race,
	})
	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetOnlyOneCharacterByRace(race string) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ Race:{ Race } }) RETURN character", map[string]interface{}{
		"Race": race,
	})

	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return character, MultipleCharacterFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetAllCharacterByRace(race string) ([]Character, error) {
	var characters []Character
	conn, err := driver.OpenPool()
	if err != nil {
		return characters, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ Race:{ Race } }) RETURN character", map[string]interface{}{
		"Race": race,
	})

	if err != nil {
		return characters, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return characters, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return characters, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		character := Character{}
		if Id, ok := node.Properties["Id"].(string); ok {
			character.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			character.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			character.LastName = LastName
		}
		if Race, ok := node.Properties["Race"].(string); ok {
			character.Race = Race
		}
		if Class, ok := node.Properties["Class"].(string); ok {
			character.Class = Class
		}
		if NPC, ok := node.Properties["NPC"].(bool); ok {
			character.NPC = NPC
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func UpdateAllCharacterByRace(race string, character Character) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ Race:{ Race } }) SET character += { Id:{characterId}, FirstName:{characterFirstName}, LastName:{characterLastName}, Race:{characterRace}, Class:{characterClass}, NPC:{characterNPC} }", map[string]interface{}{
		"Race":               race,
		"characterId":        character.Id,
		"characterFirstName": character.FirstName,
		"characterLastName":  character.LastName,
		"characterRace":      character.Race,
		"characterClass":     character.Class,
		"characterNPC":       character.NPC,
	})
	return err
}

func DeleteAllCharacterByRace(race string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ Race:{ Race }) DETACH DELETE character", map[string]interface{}{
		"Race": race,
	})
	return err
}

func GetCharacterByClass(class string) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ Class:{ Class } }) RETURN character", map[string]interface{}{
		"Class": class,
	})
	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetOnlyOneCharacterByClass(class string) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ Class:{ Class } }) RETURN character", map[string]interface{}{
		"Class": class,
	})

	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return character, MultipleCharacterFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetAllCharacterByClass(class string) ([]Character, error) {
	var characters []Character
	conn, err := driver.OpenPool()
	if err != nil {
		return characters, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ Class:{ Class } }) RETURN character", map[string]interface{}{
		"Class": class,
	})

	if err != nil {
		return characters, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return characters, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return characters, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		character := Character{}
		if Id, ok := node.Properties["Id"].(string); ok {
			character.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			character.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			character.LastName = LastName
		}
		if Race, ok := node.Properties["Race"].(string); ok {
			character.Race = Race
		}
		if Class, ok := node.Properties["Class"].(string); ok {
			character.Class = Class
		}
		if NPC, ok := node.Properties["NPC"].(bool); ok {
			character.NPC = NPC
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func UpdateAllCharacterByClass(class string, character Character) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ Class:{ Class } }) SET character += { Id:{characterId}, FirstName:{characterFirstName}, LastName:{characterLastName}, Race:{characterRace}, Class:{characterClass}, NPC:{characterNPC} }", map[string]interface{}{
		"Class":              class,
		"characterId":        character.Id,
		"characterFirstName": character.FirstName,
		"characterLastName":  character.LastName,
		"characterRace":      character.Race,
		"characterClass":     character.Class,
		"characterNPC":       character.NPC,
	})
	return err
}

func DeleteAllCharacterByClass(class string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ Class:{ Class }) DETACH DELETE character", map[string]interface{}{
		"Class": class,
	})
	return err
}

func GetCharacterByNPC(nPC bool) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ NPC:{ NPC } }) RETURN character", map[string]interface{}{
		"NPC": nPC,
	})
	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetOnlyOneCharacterByNPC(nPC bool) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ NPC:{ NPC } }) RETURN character", map[string]interface{}{
		"NPC": nPC,
	})

	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return character, MultipleCharacterFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetAllCharacterByNPC(nPC bool) ([]Character, error) {
	var characters []Character
	conn, err := driver.OpenPool()
	if err != nil {
		return characters, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (character:Character{ NPC:{ NPC } }) RETURN character", map[string]interface{}{
		"NPC": nPC,
	})

	if err != nil {
		return characters, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return characters, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return characters, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		character := Character{}
		if Id, ok := node.Properties["Id"].(string); ok {
			character.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			character.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			character.LastName = LastName
		}
		if Race, ok := node.Properties["Race"].(string); ok {
			character.Race = Race
		}
		if Class, ok := node.Properties["Class"].(string); ok {
			character.Class = Class
		}
		if NPC, ok := node.Properties["NPC"].(bool); ok {
			character.NPC = NPC
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func UpdateAllCharacterByNPC(nPC bool, character Character) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ NPC:{ NPC } }) SET character += { Id:{characterId}, FirstName:{characterFirstName}, LastName:{characterLastName}, Race:{characterRace}, Class:{characterClass}, NPC:{characterNPC} }", map[string]interface{}{
		"NPC":                nPC,
		"characterId":        character.Id,
		"characterFirstName": character.FirstName,
		"characterLastName":  character.LastName,
		"characterRace":      character.Race,
		"characterClass":     character.Class,
		"characterNPC":       character.NPC,
	})
	return err
}

func DeleteAllCharacterByNPC(nPC bool) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (character:Character{ NPC:{ NPC }) DETACH DELETE character", map[string]interface{}{
		"NPC": nPC,
	})
	return err
}

func GetCharacterByCustom(query map[string]interface{}) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	queryStr := "MATCH (character:Character{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN character"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetOnlyOneCharacterByCustom(query map[string]interface{}) (Character, error) {
	character := Character{}

	conn, err := driver.OpenPool()
	if err != nil {
		return character, err
	}
	defer conn.Close()

	queryStr := "MATCH (character:Character{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN character"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return character, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return character, NoCharacterFound
	}
	if err != nil {
		return character, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return character, MultipleCharacterFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return character, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		character.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		character.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		character.LastName = LastName
	}
	if Race, ok := node.Properties["Race"].(string); ok {
		character.Race = Race
	}
	if Class, ok := node.Properties["Class"].(string); ok {
		character.Class = Class
	}
	if NPC, ok := node.Properties["NPC"].(bool); ok {
		character.NPC = NPC
	}

	return character, nil
}

func GetAllCharacterByCustom(query map[string]interface{}) ([]Character, error) {
	var characters []Character

	conn, err := driver.OpenPool()
	if err != nil {
		return characters, err
	}
	defer conn.Close()

	queryStr := "MATCH (character:Character{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN character"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return characters, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return characters, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return characters, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		character := Character{}
		if Id, ok := node.Properties["Id"].(string); ok {
			character.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			character.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			character.LastName = LastName
		}
		if Race, ok := node.Properties["Race"].(string); ok {
			character.Race = Race
		}
		if Class, ok := node.Properties["Class"].(string); ok {
			character.Class = Class
		}
		if NPC, ok := node.Properties["NPC"].(bool); ok {
			character.NPC = NPC
		}

		characters = append(characters, character)
	}

	return characters, nil
}

func UpdateAllCharacterByCustom(params map[string]interface{}, character Character) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (character:Character{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET character += { Id:{characterId}, FirstName:{characterFirstName}, LastName:{characterLastName}, Race:{characterRace}, Class:{characterClass}, NPC:{characterNPC} }"

	params["characterId"] = character.Id
	params["characterFirstName"] = character.FirstName
	params["characterLastName"] = character.LastName
	params["characterRace"] = character.Race
	params["characterClass"] = character.Class
	params["characterNPC"] = character.NPC

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAllCharacterByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (character:Character{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE character"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
