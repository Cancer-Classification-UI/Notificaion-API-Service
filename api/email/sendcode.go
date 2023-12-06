package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"ccu/api"
	mAPI "ccu/model/api"

	"regexp"

	log "github.com/sirupsen/logrus"
)

// PostSendCode godoc
// @Summary      Sends code to a user for verification
// @Description  Connects to Google SMTP Server and sends an email with a 6 digit code
// @Tags         Code
// @Accept       json
// @Produce      json
// @Param        email             query string    true "email to send the code to"
// @Param        code              query int       true "code to send in the email"
// @Success      200  {array}   mAPI.SendCodeMessage
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /send-code [post]
func PostSendCode(w http.ResponseWriter, r *http.Request) {
	log.Info("In code sending handler -------------------------")

	r.ParseForm()
	if r.Method != http.MethodPost {
		api.Respond(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	email := r.Form.Get("email")
	code := r.Form.Get("code")

	regex := regexp.MustCompile(`^..*@.*.\\.(com|net|org|edu)$`)

	if !regex.MatchString(email) {
		api.Respond(w, "Invalid Email Parameter", http.StatusBadRequest)
		return
	}

	if code == "" { // What is the default int value?
		api.Respond(w, "Invalid Code Parameter", http.StatusBadRequest)
		return
	}

	response := mAPI.SendCodeMessage{
		DateCreated: time.Now(),
		Success:     SendCode(email, code),
	}

	api.RespondOK(w, response)
}

func SendCode(email string, code string) bool {
	// Define the SMTP server and authentication information
	smtpServer := "smtp.gmail.com"
	smtpPort := 587 // Use 587 with STARTTLS or 465 with SSL/TLS
	senderEmail := "cancerclassificationproject@gmail.com"
	senderPassword := "pdkz trbp jpwh bwjh"

	// Specify the file path
	filePath := "email.html"

	// Read the file
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("Error reading file: ", err)
		return false
	}

	// Compose the email message
	subject := "Your Security Code"
	codeInt, err := strconv.Atoi(code)
	if err != nil {
		log.Error("Error converting code to int: ", err)
		return false
	}
	codeStr := fmt.Sprintf("%06d", codeInt)
	body := string(content)
	body = strings.Replace(body, "XXXXXX", codeStr, -1)

	// Set up the authentication for the SMTP server
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	// Compose the email message
	message := "Subject: " + subject + "\r\n" +
		"To: " + email + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n\r\n" +
		body

	// Connect to the SMTP server
	smtpAddress := smtpServer + ":" + strconv.Itoa(smtpPort)
	err = smtp.SendMail(smtpAddress, auth, senderEmail, []string{email}, []byte(message))
	log.Info(codeStr)
	if err != nil {
		log.Error("Error sending email: ", err)
		return false
	}

	return true
}
