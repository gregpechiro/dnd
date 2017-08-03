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

var NoLocationFound = fmt.Errorf("no location found")
var MultipleLocationFound = fmt.Errorf("multiple location found")

func AddLocation(location Location) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE (location:Location { Id:{locationId}, Name:{locationName} })", map[string]interface{}{
		"locationId":   location.Id,
		"locationName": location.Name,
	})

	return err
}

func GetAllLocation() ([]Location, error) {
	var locations []Location
	conn, err := driver.OpenPool()
	if err != nil {
		return locations, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (location:Location) RETURN location", nil)
	if err != nil {
		return locations, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return locations, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return locations, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		location := Location{}
		if Id, ok := node.Properties["Id"].(string); ok {
			location.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			location.Name = Name
		}

		locations = append(locations, location)
	}

	return locations, nil
}

func IndexLocationById() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :Location(Id)", nil)

	return err
}

func GetLocationById(id string) (Location, error) {
	location := Location{}

	conn, err := driver.OpenPool()
	if err != nil {
		return location, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (location:Location{ Id:{ Id } }) RETURN location", map[string]interface{}{
		"Id": id,
	})
	if err != nil {
		return location, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return location, NoLocationFound
	}
	if err != nil {
		return location, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return location, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		location.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		location.Name = Name
	}

	return location, nil
}

func GetOnlyOneLocationById(id string) (Location, error) {
	location := Location{}

	conn, err := driver.OpenPool()
	if err != nil {
		return location, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (location:Location{ Id:{ Id } }) RETURN location", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return location, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return location, NoLocationFound
	}
	if err != nil {
		return location, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return location, MultipleLocationFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return location, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		location.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		location.Name = Name
	}

	return location, nil
}

func GetAllLocationById(id string) ([]Location, error) {
	var locations []Location
	conn, err := driver.OpenPool()
	if err != nil {
		return locations, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (location:Location{ Id:{ Id } }) RETURN location", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return locations, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return locations, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return locations, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		location := Location{}
		if Id, ok := node.Properties["Id"].(string); ok {
			location.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			location.Name = Name
		}

		locations = append(locations, location)
	}

	return locations, nil
}

func UpdateAllLocationById(id string, location Location) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (location:Location{ Id:{ Id } }) SET location += { Id:{locationId}, Name:{locationName} }", map[string]interface{}{
		"Id":           id,
		"locationId":   location.Id,
		"locationName": location.Name,
	})
	return err
}

func DeleteAllLocationById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (location:Location{ Id:{ Id }) DETACH DELETE location", map[string]interface{}{
		"Id": id,
	})
	return err
}

func GetLocationByName(name string) (Location, error) {
	location := Location{}

	conn, err := driver.OpenPool()
	if err != nil {
		return location, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (location:Location{ Name:{ Name } }) RETURN location", map[string]interface{}{
		"Name": name,
	})
	if err != nil {
		return location, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return location, NoLocationFound
	}
	if err != nil {
		return location, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return location, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		location.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		location.Name = Name
	}

	return location, nil
}

func GetOnlyOneLocationByName(name string) (Location, error) {
	location := Location{}

	conn, err := driver.OpenPool()
	if err != nil {
		return location, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (location:Location{ Name:{ Name } }) RETURN location", map[string]interface{}{
		"Name": name,
	})

	if err != nil {
		return location, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return location, NoLocationFound
	}
	if err != nil {
		return location, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return location, MultipleLocationFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return location, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		location.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		location.Name = Name
	}

	return location, nil
}

func GetAllLocationByName(name string) ([]Location, error) {
	var locations []Location
	conn, err := driver.OpenPool()
	if err != nil {
		return locations, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (location:Location{ Name:{ Name } }) RETURN location", map[string]interface{}{
		"Name": name,
	})

	if err != nil {
		return locations, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return locations, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return locations, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		location := Location{}
		if Id, ok := node.Properties["Id"].(string); ok {
			location.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			location.Name = Name
		}

		locations = append(locations, location)
	}

	return locations, nil
}

func UpdateAllLocationByName(name string, location Location) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (location:Location{ Name:{ Name } }) SET location += { Id:{locationId}, Name:{locationName} }", map[string]interface{}{
		"Name":         name,
		"locationId":   location.Id,
		"locationName": location.Name,
	})
	return err
}

func DeleteAllLocationByName(name string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (location:Location{ Name:{ Name }) DETACH DELETE location", map[string]interface{}{
		"Name": name,
	})
	return err
}

func GetLocationByCustom(query map[string]interface{}) (Location, error) {
	location := Location{}

	conn, err := driver.OpenPool()
	if err != nil {
		return location, err
	}
	defer conn.Close()

	queryStr := "MATCH (location:Location{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN location"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return location, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return location, NoLocationFound
	}
	if err != nil {
		return location, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return location, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		location.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		location.Name = Name
	}

	return location, nil
}

func GetOnlyOneLocationByCustom(query map[string]interface{}) (Location, error) {
	location := Location{}

	conn, err := driver.OpenPool()
	if err != nil {
		return location, err
	}
	defer conn.Close()

	queryStr := "MATCH (location:Location{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN location"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return location, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return location, NoLocationFound
	}
	if err != nil {
		return location, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return location, MultipleLocationFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return location, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		location.Id = Id
	}
	if Name, ok := node.Properties["Name"].(string); ok {
		location.Name = Name
	}

	return location, nil
}

func GetAllLocationByCustom(query map[string]interface{}) ([]Location, error) {
	var locations []Location

	conn, err := driver.OpenPool()
	if err != nil {
		return locations, err
	}
	defer conn.Close()

	queryStr := "MATCH (location:Location{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN location"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return locations, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return locations, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return locations, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		location := Location{}
		if Id, ok := node.Properties["Id"].(string); ok {
			location.Id = Id
		}
		if Name, ok := node.Properties["Name"].(string); ok {
			location.Name = Name
		}

		locations = append(locations, location)
	}

	return locations, nil
}

func UpdateAllLocationByCustom(params map[string]interface{}, location Location) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (location:Location{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET location += { Id:{locationId}, Name:{locationName} }"

	params["locationId"] = location.Id
	params["locationName"] = location.Name

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAllLocationByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (location:Location{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE location"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
