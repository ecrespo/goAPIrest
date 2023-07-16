package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ecrespo/goAPIrest/api/auth"
	"github.com/ecrespo/goAPIrest/api/models"
	"github.com/ecrespo/goAPIrest/api/responses"
	"github.com/ecrespo/goAPIrest/api/utils/formaterror"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) validateAndUnmarshalPost(r *http.Request, post *models.Post) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &post)
	if err != nil {
		return err
	}
	post.Prepare()
	return post.Validate()
}

func (server *Server) parseAndCheckPostIDAndUserID(r *http.Request) (uid uint32, pid uint64, err error) {
	vars := mux.Vars(r)
	pid, err = strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	uid, err = auth.ExtractTokenID(r)
	if err != nil {
		return 0, 0, err
	}
	return uid, pid, err
}

func (server *Server) findPostByID(uid uint32, pid uint64) (models.Post, error) {
	post, err := server.PostRepository.FindByID(pid)
	if err != nil {
		return nil, err
	}
	if uid != post.AuthorID {
		return nil, errors.New("Unauthorized")
	}
	return post, nil
}

func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	err := server.validateAndUnmarshalPost(r, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil || uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	postCreated, err := server.PostRepository.Save(&post)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.ID))
	responses.JSON(w, http.StatusCreated, postCreated)
}

func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {

	posts, err := server.PostRepository.FindAll()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	_, pid, err := server.parseAndCheckPostIDAndUserID(r)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	postReceived, err := server.findPostByID(pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, postReceived)
}

func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	uid, pid, err := server.parseAndCheckPostIDAndUserID(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	post, err := server.findPostByID(pid)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Post not found"))
		return
	}

	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	postUpdate, err := server.validateAndUnmarshalPost(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	if uid != postUpdate.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	postUpdate.ID = pid
	postUpdated, err := server.PostRepository.Update(*postUpdate)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, postUpdated)
}

func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	uid, pid, err := server.parseAndCheckPostIDAndUserID(r)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	post, err := server.findPostByID(pid)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	_, err = server.PostRepository.Delete(pid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
