package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
)

// Logging struct construction.
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Struct for JWT.
type JwtClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

var jwtKey = []byte("SecretStaticJWTkey") // Secret Static Key for JWT, only for testing :)

func main() {
	r := gin.Default()

	// MySql_MariaDB connection setting.
	db, err := sql.Open("mysql", "usertest1:7777777xX@tcp(localhost:3306)/test_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Endpoint for loging.
	r.POST("/login", func(c *gin.Context) {
		var loginReq Login
		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Checking loging details (hardcoded for example)
		if loginReq.Username != "jankowalsky" || loginReq.Password != "987656789aA!" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect login details."})
			return
		}

		// JWT Token generation code (5 min. token validation)
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &JwtClaims{
			Username: loginReq.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token cannot be generated!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	})

	r.POST("/addtestuser1", func(c *gin.Context) {
		// Using actual date and time as unique surname.
		currentTimestamp := time.Now().Format("20060102150405") // Format YYYYMMDDHHMMSS

		// Test user data.
		name := "TestUser"
		surname := currentTimestamp // user surname as a actual date and time.
		age := "30"
		sex := "M"

		// INSERT query with UPDATE options if key already exists.
		query := "INSERT INTO tab1 (NAME, SURNAME, AGE, SEX) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE NAME = VALUES(NAME), AGE = VALUES(AGE), SEX = VALUES(SEX)"
		_, err := db.Exec(query, name, surname, age, sex)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User has been added or updated."})
	})

	r.DELETE("/deleteuser/:surname", func(c *gin.Context) {
		nazwisko := c.Param("surname")

		// DELETE query execution.
		_, err := db.Exec("DELETE FROM tab1 WHERE SURNAME = ?", nazwisko)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User has been deleted. "})
	})

	r.GET("/checkuser/:surname", func(c *gin.Context) {
		nazwisko := c.Param("surname")

		// Check whether the user exists in the database.
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM tab1 WHERE SURNAME = ?", nazwisko).Scan(&count)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count > 0 {
			c.JSON(http.StatusOK, gin.H{"exists": true})
		} else {
			c.JSON(http.StatusOK, gin.H{"exists": false})
		}
	})

	// Endpoint SQL Query.
	r.POST("/query", func(c *gin.Context) {
		var request struct {
			Token string `json:"token"`
			Query string `json:"query"` // Accepts an SQL query.
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// JWT token verification.
		token, err := jwt.ParseWithClaims(request.Token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// SQL query execution.
		rows, err := db.Query(request.Query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error until query execution."})
			return
		}
		defer rows.Close()

		// collecting result.
		var results []map[string]interface{}
		cols, _ := rows.Columns()
		for rows.Next() {
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i := range columns {
				columnPointers[i] = &columns[i]
			}

			if err := rows.Scan(columnPointers...); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			m := make(map[string]interface{})
			for i, colName := range cols {
				val := columnPointers[i].(*interface{})
				m[colName] = *val
			}
			results = append(results, m)
		}

		c.JSON(http.StatusOK, results)
	})

	// Uruchomienie serwera
	r.Run(":8080")
}
