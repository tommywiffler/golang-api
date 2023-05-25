package cmd

import (
	"encoding/json"
	"fmt"
	"golang-api/infra"
	"golang-api/pkg/db"
	"golang-api/pkg/utils"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateFriend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]

	var userID int64
	userID, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to parse user id, see error: %v", err))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to read request body, see error: %v", err))
		return
	}

	var friendID infra.FriendID
	err = json.Unmarshal(body, &friendID)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to unmarshal request body, see error: %v", err))
		return
	}

	var createFriend infra.Friend
	createFriend.FriendID = friendID.ID
	createFriend.UserID = userID

	res, err := infra.CreateFriend(db.DB, createFriend)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to add friend, see error: %v", err))
		return
	}

	utils.JSON(w, http.StatusOK, res)

}

func GetFriends(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["id"]

	var userID int64
	userID, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, fmt.Errorf("unable to parse user id, see error: %v", err))
		return
	}

	var friends *[]infra.Friend
	friends, err = infra.GetFriends(db.DB, userID)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to get friends, see error: %v", err))
		return
	}

	var res []int64

	for _, friend := range *friends {
		res = append(res, friend.FriendID)
	}

	utils.JSON(w, http.StatusOK, res)
}

func DeleteFriend(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	fid := vars["fid"]

	err := infra.DeleteFriend(db.DB, uid, fid)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, fmt.Errorf("unable to delete friend, see error: %v", err))
		return
	}

	res := "Friend has been deleted"

	utils.JSON(w, http.StatusOK, res)

}
