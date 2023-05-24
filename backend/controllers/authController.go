package controllers

import (
	"backend/database"
	"backend/models"
	"backend/utility"
	"math/rand"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	token, err := utility.VerifyAuth(c)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("client", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := database.GetUserFromToken(token)

	if user.AdminRole == 0 {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid admin role",
		})
	}

	var data map[string]string
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	adminRole, _ := strconv.Atoi(data["adminRole"])
	businessRole, _ := strconv.Atoi(data["businessRole"])
	residentialRole, _ := strconv.Atoi(data["residentialRole"])

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if err = utility.ValidateNewPassword(data["password"], nil); err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	userToAdd := models.User{
		Name:              data["name"],
		Email:             data["email"],
		Password:          password,
		AdminRole:         uint(adminRole),
		BusinessRole:      uint(businessRole),
		ResidentialRole:   uint(residentialRole),
		Blocked:           2,
		LastModified:      time.Now(),
		LastLoginAttempt:  time.Time{},
		LoginAttemptCount: 0,
	}

	err = database.DB.Create(&userToAdd).Error

	if err != nil {
		utility.LogInfo("system", "errCreateUser")
		return c.JSON(fiber.Map{
			"message": "error creating the user. Try with another email. ",
			"error":   err.Error(),
		})
	}

	passwordHistory := models.PasswordHistory{
		UserEmail:   userToAdd.Email,
		CreatedTime: time.Now(),
		Password:    password,
	}

	errPwdHistory := database.DB.Create(&passwordHistory).Error

	if errPwdHistory != nil {
		utility.LogInfo("system", "errCreatePwdHistory")
		return c.JSON(fiber.Map{
			"message": "error creating the password history.",
			"error":   errPwdHistory.Error(),
		})
	}

	utility.LogInfo(data["name"], "createUser")

	return c.JSON(userToAdd)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := database.GetUserFromEmail(data["email"])

	if err != nil {
		c.Status(fiber.StatusNotFound)
		utility.LogInfo("system", "userNotFound")
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	loginPolicy, err := database.GetLoginPolicy(user)

	if err != nil {
		c.Status(fiber.StatusNotFound)
		utility.LogInfo("system", "loginpolicynotfound")
		return c.JSON(fiber.Map{
			"message": "loginpolicy not found",
		})
	}

	if err := utility.Login(user, loginPolicy, data["password"]); err != nil {
		c.Status(fiber.StatusBadRequest)
		utility.LogInfo("system", err.Error())

		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if !loginPolicy.TwoFA {
		utility.GenerateJWT(*user, c)
	}

	database.DB.Save(&user)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func ModifyPassword(c *fiber.Ctx) error {

	token, err := utility.VerifyAuth(c)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo("client", "unauthenticated")
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	user := database.GetUserFromToken(token)

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var userToEdit models.User
	database.DB.Where("email = ?", data["email"]).First(&userToEdit)

	if user.AdminRole == 0 && user.Email != userToEdit.Email {

		c.Status(fiber.StatusUnauthorized)
		utility.LogInfo(user.Name, "unauthorized")
		return c.JSON(fiber.Map{
			"message": "the user does not have a valid flicka the wrist",
		})
	}

	err = utility.ValidateNewPassword(data["password"], nil)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	passwordHistoryList := utility.GetLastXPassword(userToEdit)

	newPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "error creating new password.",
		})
	}

	for _, password := range *passwordHistoryList {
		err := bcrypt.CompareHashAndPassword(password.Password, []byte(data["password"]))
		if err == nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "Password already used",
			})
		}
	}

	userToEdit.LastModified = time.Now()
	userToEdit.Password = newPassword
	userToEdit.LoginAttemptCount = 0

	// The user must change his password if was blocked by admin
	if user.Email != userToEdit.Email {
		userToEdit.Blocked = 2
	} else {
		userToEdit.Blocked = 0
	}

	errModifyPassword := database.DB.Save(&userToEdit).Error

	if errModifyPassword != nil {
		utility.LogInfo("system", "errModifyPassword")
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "error modifying the password.",
		})
	}

	passwordHistory := models.PasswordHistory{
		UserEmail:   user.Email,
		CreatedTime: time.Now(),
		Password:    newPassword,
	}

	errPwdList := database.DB.Create(&passwordHistory).Error

	if errPwdList != nil {
		utility.LogInfo("system", "errCreatePwdList")
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "error creating the password list.",
		})
	}

	utility.LogInfo(data["name"], "modifyPassword")

	return c.JSON(fiber.Map{
		"user":    user,
		"message": "success",
	})

}

func Logout(c *fiber.Ctx) error {
	// creer un cookie dans le passé pour "retirer" le cookie du browser
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	utility.LogInfo("system", "logout")

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Login2FA(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	posx, _ := strconv.Atoi(data["positionx"])
	posy, _ := strconv.Atoi(data["positiony"])

	if utility.VerifyTwoFA(posx, posy, data["code"]) {
		var user models.User
		database.DB.Where("email = ?", data["email"]).First(&user)

		if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
			c.Status(fiber.StatusBadRequest)
			user.LoginAttemptCount = user.LoginAttemptCount + 1
			user.LastLoginAttempt = time.Now()
			database.DB.Save(&user)
			utility.LogInfo("system", "wrongPassword")
			return c.JSON(fiber.Map{
				"message": "incorrect password",
			})
		}
		utility.GenerateJWT(user, c)

		return c.JSON(fiber.Map{
			"user":    user,
			"message": "success",
		})

	} else {
		utility.LogInfo("system", "errCreatePwdList")
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"error": "Incorrect code",
		})
	}
}

func GetCode(c *fiber.Ctx) error {
	answer := c.JSON(fiber.Map{
		"positionx": rand.Intn(2),
		"positiony": rand.Intn(2),
	})

	return answer
}
