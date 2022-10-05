package utils

import (
	"content/src/config"
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, cookieName, email string) {
	cookie := &http.Cookie{
		Name:  cookieName,
		Value: email,
	}
	http.SetCookie(w, cookie)
}

func ClearTokenHandler(w http.ResponseWriter, cookieName string) {
	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   "",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(w, cookie)
}

func HasCookieAdminWriteHeader(w http.ResponseWriter, r *http.Request) bool {
	permission, err := r.Cookie(config.MANAGER_PERSMISSION)
	if permission == nil || err != nil {
		http.Error(w, "Доступ запрещен", http.StatusForbidden)
		return false
	}
	return true
}

func HasCookieUserWriteHeader(w http.ResponseWriter, r *http.Request) bool {
	permission, err := r.Cookie(config.USER_PERSMISSION)
	if permission == nil || err != nil {
		http.Error(w, "Доступ запрещен", http.StatusForbidden)
		return false
	}
	return true
}

func HasCookieAdmin(r *http.Request) bool {
	permission, err := r.Cookie(config.MANAGER_PERSMISSION)
	if permission == nil || err != nil {
		return false
	}
	return true
}
