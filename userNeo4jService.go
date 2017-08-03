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

var NoUserFound = fmt.Errorf("no user found")
var MultipleUserFound = fmt.Errorf("multiple user found")

func AddUser(user User) error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE (user:User { Id:{userId}, FirstName:{userFirstName}, LastName:{userLastName}, Email:{userEmail}, Password:{userPassword}, Active:{userActive}, Role:{userRole} })", map[string]interface{}{
		"userId":        user.Id,
		"userFirstName": user.FirstName,
		"userLastName":  user.LastName,
		"userEmail":     user.Email,
		"userPassword":  user.Password,
		"userActive":    user.Active,
		"userRole":      user.Role,
	})

	return err
}

func GetAllUser() ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User) RETURN user", nil)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			user.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			user.LastName = LastName
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Password, ok := node.Properties["Password"].(string); ok {
			user.Password = Password
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Role, ok := node.Properties["Role"].(string); ok {
			user.Role = Role
		}

		users = append(users, user)
	}

	return users, nil
}

func IndexUserById() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :User(Id)", nil)

	return err
}

func GetUserById(id string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Id:{ Id } }) RETURN user", map[string]interface{}{
		"Id": id,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetOnlyOneUserById(id string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Id:{ Id } }) RETURN user", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetAllUserById(id string) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Id:{ Id } }) RETURN user", map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			user.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			user.LastName = LastName
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Password, ok := node.Properties["Password"].(string); ok {
			user.Password = Password
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Role, ok := node.Properties["Role"].(string); ok {
			user.Role = Role
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserById(id string, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Id:{ Id } }) SET user += { Id:{userId}, FirstName:{userFirstName}, LastName:{userLastName}, Email:{userEmail}, Password:{userPassword}, Active:{userActive}, Role:{userRole} }", map[string]interface{}{
		"Id":            id,
		"userId":        user.Id,
		"userFirstName": user.FirstName,
		"userLastName":  user.LastName,
		"userEmail":     user.Email,
		"userPassword":  user.Password,
		"userActive":    user.Active,
		"userRole":      user.Role,
	})
	return err
}

func DeleteAllUserById(id string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Id:{ Id }) DETACH DELETE user", map[string]interface{}{
		"Id": id,
	})
	return err
}

func GetUserByFirstName(firstName string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ FirstName:{ FirstName } }) RETURN user", map[string]interface{}{
		"FirstName": firstName,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetOnlyOneUserByFirstName(firstName string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ FirstName:{ FirstName } }) RETURN user", map[string]interface{}{
		"FirstName": firstName,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetAllUserByFirstName(firstName string) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ FirstName:{ FirstName } }) RETURN user", map[string]interface{}{
		"FirstName": firstName,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			user.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			user.LastName = LastName
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Password, ok := node.Properties["Password"].(string); ok {
			user.Password = Password
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Role, ok := node.Properties["Role"].(string); ok {
			user.Role = Role
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByFirstName(firstName string, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ FirstName:{ FirstName } }) SET user += { Id:{userId}, FirstName:{userFirstName}, LastName:{userLastName}, Email:{userEmail}, Password:{userPassword}, Active:{userActive}, Role:{userRole} }", map[string]interface{}{
		"FirstName":     firstName,
		"userId":        user.Id,
		"userFirstName": user.FirstName,
		"userLastName":  user.LastName,
		"userEmail":     user.Email,
		"userPassword":  user.Password,
		"userActive":    user.Active,
		"userRole":      user.Role,
	})
	return err
}

func DeleteAllUserByFirstName(firstName string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ FirstName:{ FirstName }) DETACH DELETE user", map[string]interface{}{
		"FirstName": firstName,
	})
	return err
}

func GetUserByLastName(lastName string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ LastName:{ LastName } }) RETURN user", map[string]interface{}{
		"LastName": lastName,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetOnlyOneUserByLastName(lastName string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ LastName:{ LastName } }) RETURN user", map[string]interface{}{
		"LastName": lastName,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetAllUserByLastName(lastName string) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ LastName:{ LastName } }) RETURN user", map[string]interface{}{
		"LastName": lastName,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			user.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			user.LastName = LastName
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Password, ok := node.Properties["Password"].(string); ok {
			user.Password = Password
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Role, ok := node.Properties["Role"].(string); ok {
			user.Role = Role
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByLastName(lastName string, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ LastName:{ LastName } }) SET user += { Id:{userId}, FirstName:{userFirstName}, LastName:{userLastName}, Email:{userEmail}, Password:{userPassword}, Active:{userActive}, Role:{userRole} }", map[string]interface{}{
		"LastName":      lastName,
		"userId":        user.Id,
		"userFirstName": user.FirstName,
		"userLastName":  user.LastName,
		"userEmail":     user.Email,
		"userPassword":  user.Password,
		"userActive":    user.Active,
		"userRole":      user.Role,
	})
	return err
}

func DeleteAllUserByLastName(lastName string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ LastName:{ LastName }) DETACH DELETE user", map[string]interface{}{
		"LastName": lastName,
	})
	return err
}

func IndexUserByEmail() error {
	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("CREATE INDEX ON :User(Email)", nil)

	return err
}

func GetUserByEmail(email string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Email:{ Email } }) RETURN user", map[string]interface{}{
		"Email": email,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetOnlyOneUserByEmail(email string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Email:{ Email } }) RETURN user", map[string]interface{}{
		"Email": email,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetAllUserByEmail(email string) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Email:{ Email } }) RETURN user", map[string]interface{}{
		"Email": email,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			user.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			user.LastName = LastName
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Password, ok := node.Properties["Password"].(string); ok {
			user.Password = Password
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Role, ok := node.Properties["Role"].(string); ok {
			user.Role = Role
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByEmail(email string, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Email:{ Email } }) SET user += { Id:{userId}, FirstName:{userFirstName}, LastName:{userLastName}, Email:{userEmail}, Password:{userPassword}, Active:{userActive}, Role:{userRole} }", map[string]interface{}{
		"Email":         email,
		"userId":        user.Id,
		"userFirstName": user.FirstName,
		"userLastName":  user.LastName,
		"userEmail":     user.Email,
		"userPassword":  user.Password,
		"userActive":    user.Active,
		"userRole":      user.Role,
	})
	return err
}

func DeleteAllUserByEmail(email string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Email:{ Email }) DETACH DELETE user", map[string]interface{}{
		"Email": email,
	})
	return err
}

func GetUserByPassword(password string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Password:{ Password } }) RETURN user", map[string]interface{}{
		"Password": password,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetOnlyOneUserByPassword(password string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Password:{ Password } }) RETURN user", map[string]interface{}{
		"Password": password,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetAllUserByPassword(password string) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Password:{ Password } }) RETURN user", map[string]interface{}{
		"Password": password,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			user.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			user.LastName = LastName
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Password, ok := node.Properties["Password"].(string); ok {
			user.Password = Password
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Role, ok := node.Properties["Role"].(string); ok {
			user.Role = Role
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByPassword(password string, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Password:{ Password } }) SET user += { Id:{userId}, FirstName:{userFirstName}, LastName:{userLastName}, Email:{userEmail}, Password:{userPassword}, Active:{userActive}, Role:{userRole} }", map[string]interface{}{
		"Password":      password,
		"userId":        user.Id,
		"userFirstName": user.FirstName,
		"userLastName":  user.LastName,
		"userEmail":     user.Email,
		"userPassword":  user.Password,
		"userActive":    user.Active,
		"userRole":      user.Role,
	})
	return err
}

func DeleteAllUserByPassword(password string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Password:{ Password }) DETACH DELETE user", map[string]interface{}{
		"Password": password,
	})
	return err
}

func GetUserByActive(active bool) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Active:{ Active } }) RETURN user", map[string]interface{}{
		"Active": active,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetOnlyOneUserByActive(active bool) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Active:{ Active } }) RETURN user", map[string]interface{}{
		"Active": active,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetAllUserByActive(active bool) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Active:{ Active } }) RETURN user", map[string]interface{}{
		"Active": active,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			user.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			user.LastName = LastName
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Password, ok := node.Properties["Password"].(string); ok {
			user.Password = Password
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Role, ok := node.Properties["Role"].(string); ok {
			user.Role = Role
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByActive(active bool, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Active:{ Active } }) SET user += { Id:{userId}, FirstName:{userFirstName}, LastName:{userLastName}, Email:{userEmail}, Password:{userPassword}, Active:{userActive}, Role:{userRole} }", map[string]interface{}{
		"Active":        active,
		"userId":        user.Id,
		"userFirstName": user.FirstName,
		"userLastName":  user.LastName,
		"userEmail":     user.Email,
		"userPassword":  user.Password,
		"userActive":    user.Active,
		"userRole":      user.Role,
	})
	return err
}

func DeleteAllUserByActive(active bool) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Active:{ Active }) DETACH DELETE user", map[string]interface{}{
		"Active": active,
	})
	return err
}

func GetUserByRole(role string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Role:{ Role } }) RETURN user", map[string]interface{}{
		"Role": role,
	})
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetOnlyOneUserByRole(role string) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Role:{ Role } }) RETURN user", map[string]interface{}{
		"Role": role,
	})

	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}
	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetAllUserByRole(role string) ([]User, error) {
	var users []User
	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (user:User{ Role:{ Role } }) RETURN user", map[string]interface{}{
		"Role": role,
	})

	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			user.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			user.LastName = LastName
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Password, ok := node.Properties["Password"].(string); ok {
			user.Password = Password
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Role, ok := node.Properties["Role"].(string); ok {
			user.Role = Role
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByRole(role string, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Role:{ Role } }) SET user += { Id:{userId}, FirstName:{userFirstName}, LastName:{userLastName}, Email:{userEmail}, Password:{userPassword}, Active:{userActive}, Role:{userRole} }", map[string]interface{}{
		"Role":          role,
		"userId":        user.Id,
		"userFirstName": user.FirstName,
		"userLastName":  user.LastName,
		"userEmail":     user.Email,
		"userPassword":  user.Password,
		"userActive":    user.Active,
		"userRole":      user.Role,
	})
	return err
}

func DeleteAllUserByRole(role string) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo("MATCH (user:User{ Role:{ Role }) DETACH DELETE user", map[string]interface{}{
		"Role": role,
	})
	return err
}

func GetUserByCustom(query map[string]interface{}) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	queryStr := "MATCH (user:User{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN user"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetOnlyOneUserByCustom(query map[string]interface{}) (User, error) {
	user := User{}

	conn, err := driver.OpenPool()
	if err != nil {
		return user, err
	}
	defer conn.Close()

	queryStr := "MATCH (user:User{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN user"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	data, _, err := rows.NextNeo()
	if err == io.EOF {
		return user, NoUserFound
	}
	if err != nil {
		return user, err
	}

	if _, _, err := rows.NextNeo(); err != io.EOF {
		return user, MultipleUserFound
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		return user, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
	}

	if Id, ok := node.Properties["Id"].(string); ok {
		user.Id = Id
	}
	if FirstName, ok := node.Properties["FirstName"].(string); ok {
		user.FirstName = FirstName
	}
	if LastName, ok := node.Properties["LastName"].(string); ok {
		user.LastName = LastName
	}
	if Email, ok := node.Properties["Email"].(string); ok {
		user.Email = Email
	}
	if Password, ok := node.Properties["Password"].(string); ok {
		user.Password = Password
	}
	if Active, ok := node.Properties["Active"].(bool); ok {
		user.Active = Active
	}
	if Role, ok := node.Properties["Role"].(string); ok {
		user.Role = Role
	}

	return user, nil
}

func GetAllUserByCustom(query map[string]interface{}) ([]User, error) {
	var users []User

	conn, err := driver.OpenPool()
	if err != nil {
		return users, err
	}
	defer conn.Close()

	queryStr := "MATCH (user:User{"
	var qKeys []string
	for k, _ := range query {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) RETURN user"

	rows, err := conn.QueryNeo(queryStr, query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for data, _, err := rows.NextNeo(); err != io.EOF; data, _, err = rows.NextNeo() {
		if err != nil {
			return users, err
		}
		node, ok := data[0].(graph.Node)
		if !ok {
			return users, fmt.Errorf("data[0] is not type graph.Node it is %T\n", data[0])
		}
		user := User{}
		if Id, ok := node.Properties["Id"].(string); ok {
			user.Id = Id
		}
		if FirstName, ok := node.Properties["FirstName"].(string); ok {
			user.FirstName = FirstName
		}
		if LastName, ok := node.Properties["LastName"].(string); ok {
			user.LastName = LastName
		}
		if Email, ok := node.Properties["Email"].(string); ok {
			user.Email = Email
		}
		if Password, ok := node.Properties["Password"].(string); ok {
			user.Password = Password
		}
		if Active, ok := node.Properties["Active"].(bool); ok {
			user.Active = Active
		}
		if Role, ok := node.Properties["Role"].(string); ok {
			user.Role = Role
		}

		users = append(users, user)
	}

	return users, nil
}

func UpdateAllUserByCustom(params map[string]interface{}, user User) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (user:User{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) SET user += { Id:{userId}, FirstName:{userFirstName}, LastName:{userLastName}, Email:{userEmail}, Password:{userPassword}, Active:{userActive}, Role:{userRole} }"

	params["userId"] = user.Id
	params["userFirstName"] = user.FirstName
	params["userLastName"] = user.LastName
	params["userEmail"] = user.Email
	params["userPassword"] = user.Password
	params["userActive"] = user.Active
	params["userRole"] = user.Role

	_, err = conn.ExecNeo(queryStr, params)
	return err
}

func DeleteAllUserByCustom(params map[string]interface{}) error {

	conn, err := driver.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	queryStr := "MATCH (user:User{"
	var qKeys []string
	for k, _ := range params {
		qKeys = append(qKeys, k+": {"+k+"}")
	}
	queryStr += strings.Join(qKeys, ", ")
	queryStr += "}) DETACH DELETE user"

	_, err = conn.ExecNeo(queryStr, params)
	return err
}
