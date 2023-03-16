package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func GetAllUsersGorm(c *gin.Context) {
	db := connectGORM()
	var users []User
	db.Find(&users)
	fmt.Println(users)
	SendRespond(c, "berhasil get semua users", users)
}

func GetAllUsers(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		SendErrorResponse(c, "salah query")
		return
	}

	var user User
	var users []User
	var trash string
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &trash, &user.Country, &user.UserType); err != nil {
			log.Println(err)
			SendErrorResponse(c, "salah")
			return
		} else {
			users = append(users, user)
		}
	}

	if len(users) < 25 {
		SendRespond(c, "berhasil get semua user", users)
	} else {
		SendErrorResponse(c, "gagal get semua user")
	}
}

func InsertUserGORM(c *gin.Context) {
	db := connectGORM()

	err := c.Request.ParseForm()
	if err != nil {
		SendErrorResponse(c, "salah")
		return
	}

	name := c.Request.Form.Get("name")
	email := c.Request.Form.Get("email")
	password := c.Request.Form.Get("password")
	country := c.Request.Form.Get("country")

	h := sha1.New()
	h.Write([]byte(password))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	var user User
	user.Name = name
	user.Email = email
	user.Password = sha1_hash
	user.Country = country

	result := db.Create(&user)

	if result != nil {
		SendRespond(c, "berhasil insert user", user)
	} else {
		SendErrorResponse(c, "salah query")
	}
}

func InsertUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	err := c.Request.ParseForm()
	if err != nil {
		SendErrorResponse(c, "salah")
		return
	}

	name := c.Request.Form.Get("name")
	email := c.Request.Form.Get("email")
	password := c.Request.Form.Get("password")
	country := c.Request.Form.Get("country")

	h := sha1.New()
	h.Write([]byte(password))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	_, errQuery := db.Exec("INSERT INTO users(username, useremail, userpassword, usercountry) values (?,?,?,?)",
		name,
		email,
		sha1_hash,
		country,
	)

	var user User
	user.Name = name
	user.Email = email
	user.Password = sha1_hash
	user.Country = country

	if errQuery == nil {
		SendRespond(c, "berhasil insert user", user)
	} else {
		SendErrorResponse(c, "salah query")
	}
}

func DeleteUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	err := c.Request.ParseForm()
	if err != nil {
		return
	}

	_, errQuery := db.Exec("DELETE FROM users WHERE userid=?",
		c.Param("userId"),
	)

	if errQuery == nil {
		SendRespondDoang(c, "berhasil delete user")
	} else {
		SendErrorResponse(c, "salah query")
	}
}

func UpdateUser(c *gin.Context) {
	db := connect()
	defer db.Close()

	err := c.Request.ParseForm()
	if err != nil {
		SendErrorResponse(c, "salah")
		return
	}

	name := c.Request.Form.Get("name")
	email := c.Request.Form.Get("email")
	password := c.Request.Form.Get("password")
	country := c.Request.Form.Get("country")

	h := sha1.New()
	h.Write([]byte(password))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	_, errQuery := db.Exec("UPDATE users SET username = ?, useremail = ?, userpassword = ?, usercountry = ? WHERE userid = ?",
		name,
		email,
		sha1_hash,
		country,
		c.Param("userId"),
	)

	var user User
	user.Name = name
	user.Email = email
	user.Password = sha1_hash
	user.Country = country

	if errQuery == nil {
		SendRespond(c, "berhasil update user", user)
	} else {
		SendErrorResponse(c, "salah query")
	}
}
