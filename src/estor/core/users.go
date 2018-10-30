package core

import (
	"bufio"
	"github.com/ssrs100/logs"
	"os/exec"
	"strconv"
	"strings"
)

var (
	log = logs.GetLogger()
)

const (
	USER_HOME_PREFIX = "/home"
)

type User struct {
	Name string `json:"name"`
	Uid int `json:"uid"`
	Gid int `json:"gid"`
	Description string `description`
}

func GetUsers() []User {
	users := make([]User, 0)

	cmd := exec.Command("/usr/bin/getent", "passwd")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Error("getusers can not obtain pipe, msg:%s", err.Error())
		return nil
	}

	if err := cmd.Start(); err != nil {
		log.Error("getusers can not start cmd, msg:%s", err.Error())
		return nil
	}

	output := bufio.NewReader(stdout)

	for {
		outline, _, err := output.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				log.Error("getusers read error, msg:%s", err.Error())
			}
			break
		}
		// test:x:1000:100:for test:/home/test:/bin/bash
		line := strings.Split(string(outline), ":")
		if len(line) < 7 {
			log.Error("invalid line:%s", line)
			continue
		}
		userDir := line[5]
		if !strings.HasPrefix(userDir, USER_HOME_PREFIX) {
			continue
		}
		uid,_ := strconv.Atoi(line[2])
		gid,_ := strconv.Atoi(line[3])
		user := User{
			Name:line[0],
			Uid:uid,
			Gid:gid,
			Description:line[4],
		}
		users = append(users, user)
	}

	if err := cmd.Wait(); err != nil {
		log.Error("getusers wait error, msg:%s", err.Error())
	}
	return users
}

func AddUser(user User) error {
	params := []string{"-s", "/bin/bash", "-m", user.Name}
	if user.Uid > 0 {
		params = append(params, "-u")
		params = append(params, strconv.Itoa(user.Uid))
	}
	if user.Gid > 0 {
		params = append(params, "-g")
		params = append(params, strconv.Itoa(user.Gid))
	}
	cmd := exec.Command("/usr/sbin/useradd", params...)
	return cmd.Run()
}

func DelUser(username string) error {
	params := []string{"-r", username}
	cmd := exec.Command("/usr/sbin/userdel", params...)
	return cmd.Run()
}
