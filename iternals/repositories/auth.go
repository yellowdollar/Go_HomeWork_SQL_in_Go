package repositories

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"home_work_sql_gin/iternals/models"
	"log"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

func HashPassword(passwordToHash string) string {
	hash := sha256.New()
	hash.Write([]byte(passwordToHash))
	hashedBytes := hash.Sum(nil)

	return hex.EncodeToString(hashedBytes)
}

func SignUp(db *sqlx.DB, data models.AuthUser) error {

	hashedPassword := HashPassword(data.Password)

	stmt := `
		INSERT INTO tb_users(login, password)
		VALUES($1, $2)
	`

	_, err := db.Exec(stmt, data.Login, hashedPassword)

	return err
}

const (
	Ttl        = 60
	ServerName = "TestProject"
	SecretKey  = "some_secret_key_hello_world"
)

func GenerateJWT(userId int, userLogin string) (string, error) {
	claims := models.JWTcols{
		UserId:    userId,
		UserLogin: userLogin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Duration(Ttl) * time.Minute),
			Issuer:    ServerName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}

func SignIn(login string, password string, db *sqlx.DB) string {
	stmt := `
		SELECT * FROM tb_users WHERE login = $1
	`

	var userData []models.AuthUser

	err := db.Select(&userData, stmt, login)

	fmt.Println(userData)
	if err != nil {
		log.Fatal("User not found")
	}

	if userData[0].Password != HashPassword(password) {
		log.Fatal("Wrong password")
	}

	token, err1 := GenerateJWT(userData[0].Id, userData[0].Login)

	if err1 != nil {
		log.Fatal(err1.Error())
	}

	return token

}
