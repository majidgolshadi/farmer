package api

import (
	"time"
	"math/rand"
	"encoding/json"

	"github.com/ravaj-group/farmer/mysql/api/daemon/db"
	"github.com/ravaj-group/farmer/mysql/api/daemon/api/request"
	"github.com/ravaj-group/farmer/mysql/api/daemon/api/response"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func Create(req request.DbRequest) (int, string) {
	if (db.DB == nil) {
		return 500, db.DB_ERROR.Error()
	}

	username := newUsername();
    password := randStringBytesMaskImprSrc(12);

	if err := db.CreateDatabase(req.Database); err != nil {
		return 500, err.Error()
	}

	if err := db.CreateUser(req.Database, username, password); err != nil {
		return 500, err.Error()
	}

	json, _ := json.Marshal(&response.DbResponse{
		Database: req.Database,
		Username: username,
		Password: password,
	})

	return 201, string(json)
}

func randStringBytesMaskImprSrc(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func newUsername() string {
	founded := 1;
	username := "default_user";

	for founded > 0 {
		username = randStringBytesMaskImprSrc(16);
		if rows, _ := db.DB.Query("SELECT User FROM mysql.user WHERE user = '" + username + "'"); !rows.Next() {
			founded = 0
		}
	}

	return username;
}
