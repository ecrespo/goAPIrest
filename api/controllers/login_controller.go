package controllers

import (
	"encoding/json"
	"github.com/ecrespo/goAPIrest/api/auth"
	"github.com/ecrespo/goAPIrest/api/models"
	"github.com/ecrespo/goAPIrest/api/responses"
	"github.com/ecrespo/goAPIrest/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func (server *Server) readBody(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

func (server *Server) processUser(body []byte) (models.User, error) {
	user := models.User{}
	err := json.Unmarshal(body, &user)

	if err != nil {
		return user, err
	}

	user.Prepare()
	err = user.Validate("login")

	return user, err
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := server.readBody(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user, err := server.processUser(body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {
	user := models.User{}
	err := server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(user.ID)
}
