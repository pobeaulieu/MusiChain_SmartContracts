package utility

import (
	"backend/database"
	"backend/models"
	"time"
)

func LogInfo(username string, infoLog string) {
	log := models.Log{
		UserName: username,
		LogTime:  time.Now(),
		Message:  SetLogMessage(infoLog),
	}

	database.DB.Create(&log)
}

func SetLogMessage(info string) string {
	var msg = ""
	switch info {
	case "errCreateUser":
		msg = "Error creating user"
	case "createUser":
		msg = "User was created"
	case "login":
		msg = "Connection succeed"
	case "errDBSave":
		msg = "Could not save user data"
	case "wrongPassword":
		msg = "Incorrect password"
	case "accountBlocked":
		msg = "Account blocked"
	case "shortTimeInterval":
		msg = "Too short time interval between attempts"
	case "failedLogin":
		msg = "Login failed"
	case "userNotFound":
		msg = "User not found"
	case "logout":
		msg = "User logged out"
	case "admin":
		msg = "User unauthenticated"
	case "unauthorized":
		msg = "User does not have a valid admin role"
	case "unblocked":
		msg = "User unblocked"
	case "maxAttemptsLimit":
		msg = "Max attempts must be greater than 1"
	case "errMaxAttemptsSave":
		msg = "Could not save the new max attempts"
	case "maxAttempts":
		msg = "Max attempts limit changed"
	case "minLoginTimeInterval":
		msg = "Login time interval must be positive"
	case "errMinLoginTimeIntervalSave":
		msg = "Could not save the new min login time interval"
	case "loginTimeInterval":
		msg = "Login time interval changed"
	case "passwordPolicy":
		msg = "Password policy saved"
	case "getPasswordPolicy":
		msg = "Password policy found"
	case "getLoginPolicy":
		msg = "Login policy found"
	case "userRole":
		msg = "User role edited"
	case "noBusinessClients":
		msg = "No business clients found"
	case "businessClients":
		msg = "Business clients found"
	case "noResidentialClients":
		msg = "No residential clients found"
	case "residentialClients":
		msg = "Residential clients found"
	case "user":
		msg = "User found"
	case "noUsers":
		msg = "No users found"
	case "users":
		msg = "Users found"
	default:
		msg = "Welcome GTI619"
	}

	return msg
}
