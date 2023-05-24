package utility

import (
	"backend/database"
	"backend/models"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWT(user models.User, c *fiber.Ctx) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), //JWT valide 1 heure
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRETKEY")))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
}

func GetLastPassword(user *models.User) *models.PasswordHistory {
	var lastPassword models.PasswordHistory
	database.DB.Where("user_email = ?", user.Email).Order("created_time desc").First(&lastPassword)
	return &lastPassword
}

func GetLastXPassword(user models.User) *[]models.PasswordHistory {
	var passwordHistory models.PasswordHistory
	var lastXPasswords *[]models.PasswordHistory
	var passwordPolicy models.PasswordPolicy

	if pwdPolErr := database.DB.Where("id = ?", 1).First(&passwordPolicy).Error; pwdPolErr != nil {
		fmt.Println("error getting password policy")
	}

	if err := database.DB.Model(&passwordHistory).Order("created_time desc").Limit(*passwordPolicy.HistorySize).Where("user_email = ?", user.Email).Find(&lastXPasswords).Error; err != nil {
		fmt.Println("error getting last passwords used")
	}

	return lastXPasswords
}

func VerifyAuth(c *fiber.Ctx) (*jwt.Token, error) {
	cookie := c.Cookies("jwt")

	return jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRETKEY")), nil
	})
}

func VerifyTwoFA(positionx int, positiony int, code string) bool {
	if positionx == 0 && positiony == 0 {
		return "A" == code
	}
	if positionx == 0 && positiony == 1 {
		return "B" == code
	}
	if positionx == 1 && positiony == 0 {
		return "C" == code
	}
	if positionx == 1 && positiony == 1 {
		return "D" == code
	}
	return false
}

func ChangePasswordAfterTime(user *models.User) bool {
	lastPassword := GetLastPassword(user)
	if lastPassword.CreatedTime.Before(time.Now().Add(-72 * time.Hour)) {
		return true
	}
	return false
}

func ValidateNewPassword(password string, pwdHistory [][]byte) error {
	var policy models.PasswordPolicy
	database.DB.Where("id = ?", 1).First(&policy)

	if len(password) < *policy.MinLength {
		return fmt.Errorf("password must be at least %d characters long", *policy.MinLength)
	}
	if *policy.RequireUpper && !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}
	if *policy.RequireLower && !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}
	if *policy.RequireNumber && !strings.ContainsAny(password, "0123456789") {
		return fmt.Errorf("password must contain at least one digit")
	}
	if *policy.RequireSymbol && !strings.ContainsAny(password, "!@#$%^&*()_-+={}[]\\|/?.>,<~`") {
		return fmt.Errorf("password must contain at least one symbol")
	}

	return nil
}

func Login(user *models.User, loginPolicy *models.LoginPolicy, password string) error {
	if time.Since(user.LastLoginAttempt) < time.Duration(loginPolicy.LoginTimeInterval)*time.Second {
		// L'utilisateur a dépassé la limite de tentatives de connexion et il n'a pas encore attendu suffisamment longtemps
		LogInfo(user.Email, "shortTimeInterval")

		return fmt.Errorf("too short time interval between attempts")

	} else if user.Blocked == 1 || user.LoginAttemptCount > loginPolicy.MaxLoginAttemptCount {
		// Incrementer le nombre de tentatves manquees
		user.LoginAttemptCount = user.LoginAttemptCount + 1
		user.Blocked = 1
		// Mettre a jour le last login manque
		user.LastLoginAttempt = time.Now()
		LogInfo("system", "accountBlocked")

		return fmt.Errorf("The account is blocked. Please contact an administrator to unblock the account.")

	} else if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		// Incrementer le nombre de tentatves manquees
		user.LoginAttemptCount = user.LoginAttemptCount + 1
		user.LastLoginAttempt = time.Now()
		LogInfo("system", "wrongPassword")

		return fmt.Errorf("incorrect password")

	} else if ChangePasswordAfterTime(user) == true {
		user.LastLoginAttempt = time.Time{}
		user.LoginAttemptCount = 0
		user.Blocked = 2
	} else {
		user.LoginAttemptCount = 0
		user.LastLoginAttempt = time.Time{}
		user.Blocked = 0
	}

	return nil
}
