package process

import (
	"fmt"
)

var (
	globalUserMgr *UserMgr
)

type UserMgr struct {
	onlieUsers map[int]*UserProcess
}

func InitGlobalUserMgr() {
	globalUserMgr = &UserMgr{
		onlieUsers: make(map[int]*UserProcess, 1024),
	}
}

func (userMgr *UserMgr) AddOnlieUser(userProcess *UserProcess) {
	userMgr.onlieUsers[userProcess.UserId] = userProcess
}

func (userMgr *UserMgr) DeleteOnlieUser(userId int) {
	delete(userMgr.onlieUsers, userId)
}

func (userMgr *UserMgr) GetAllOnlieUsers() map[int]*UserProcess {
	return userMgr.onlieUsers
}

func (userMgr *UserMgr) GetOnlieUserById(userId int) (up *UserProcess, err error) {
	up, ok := userMgr.onlieUsers[userId]
	if !ok {
		err = fmt.Errorf("user %d not exist", userId)
	}
	return
}
