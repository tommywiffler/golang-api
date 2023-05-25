package cmd

import (
	"encoding/json"
	"fmt"
	"golang-api/infra"
	"golang-api/pkg/db"
	"golang-api/pkg/utils"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to read request body, see error: %v", err))
		return
	}

	var createUser infra.User
	err = json.Unmarshal(body, &createUser)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to unmarshal request body, see error: %v", err))
		return
	}

	_, err = createUser.CreateUser(db.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to create user, see error: %v", err))
		return
	}

	var userModel infra.UserModel
	err = json.Unmarshal(body, &userModel)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to unmarshal request body, see error: %v", err))
		return
	}

	var friends []infra.Friend
	for _, i := range userModel.Friends {
		var friend infra.Friend
		friend.FriendID = i
		friend.UserID = userModel.ID
		friends = append(friends, friend)

	}

	_, err = infra.CreateFriends(db.DB, friends)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to create friends of user, see error: %v", err))
		return
	}

	utils.JSON(w, http.StatusOK, userModel)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users infra.User

	results, err := users.GetUsers(db.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to get users, see error: %v", err))
		return
	}

	var userResults []infra.UserModel
	for _, user := range *results {
		var userResult infra.UserModel
		userResult.ID = user.ID
		userResult.Name = user.Name
		userResult.Email = user.Email
		userResult.Age = user.Age

		var friends *[]infra.Friend
		friends, err := infra.GetFriends(db.DB, user.ID)
		if err != nil {
			utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to get friends of users, see error: %v", err))
			return
		}

		for _, friend := range *friends {
			userResult.Friends = append(userResult.Friends, friend.FriendID)
		}

		userResults = append(userResults, userResult)

	}

	utils.JSON(w, http.StatusOK, userResults)

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid := vars["id"]
	var user infra.User

	res, err := user.GetUser(db.DB, uid)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to get user, see error: %v", err))
		return
	}

	var userResult infra.UserModel
	userResult.ID = res.ID
	userResult.Name = res.Name
	userResult.Email = res.Email
	userResult.Age = res.Age

	var friends *[]infra.Friend
	friends, err = infra.GetFriends(db.DB, user.ID)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to get friends of user, see error: %v", err))
		return
	}

	for _, friend := range *friends {
		userResult.Friends = append(userResult.Friends, friend.FriendID)
	}

	utils.JSON(w, http.StatusOK, userResult)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to read request body, see error: %v", err))
		return
	}

	var userModel infra.UserModel
	err = json.Unmarshal(body, &userModel)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to unmarshal request body, see error: %v", err))
		return
	}

	var user infra.User
	user.Name = userModel.Name
	user.Email = userModel.Email
	user.Age = userModel.Age

	var userResult infra.UserModel

	res, err := user.UpdateUser(db.DB, uid)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to update user, see error: %v", err))
		return
	}

	userResult.ID = res.ID
	userResult.Name = res.Name
	userResult.Email = res.Email
	userResult.Age = res.Age

	var friends *[]infra.Friend
	friends, err = infra.GetFriends(db.DB, userResult.ID)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to get friends of updated user, see error: %v", err))
		return
	}

	for _, friend := range *friends {
		userResult.Friends = append(userResult.Friends, friend.FriendID)
	}

	utils.JSON(w, http.StatusOK, userResult)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]
	var user *infra.User

	_, err := user.DeleteUser(db.DB, uid)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to delete user, see error: %v", err))
		return
	}

	err = infra.DeleteFriends(db.DB, uid)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to delete friends of user, see error: %v", err))
		return
	}

	res := "User has been deleted"

	utils.JSON(w, http.StatusOK, res)

}
