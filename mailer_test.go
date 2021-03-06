package auth

import (
	"html/template"
	"testing"
)

const verifyEmailTmpl string = "code:{{ .VerificationCode }}, email:{{ .Email }}"

func TestRenderHtmlBody(t *testing.T) {
	m := Emailer{TemplateCache: template.Must(template.New("verifyEmail").Parse(verifyEmailTmpl))}
	if txt, err := m.renderHTMLBody("verifyEmail", sendVerifyParams{VerificationCode: "1234", Email: "email@example.com"}); txt != "code:1234, email:email@example.com" || err != nil {
		t.Error("expected correct txt and no error", txt, err)
	}

	m = Emailer{TemplateCache: template.Must(template.New("verifyEmail").Parse(verifyEmailTmpl))}
	if _, err := m.renderHTMLBody("verifyEmail..", nil); err == nil {
		t.Error("expected error", err)
	}
}

func TestSends(t *testing.T) {
	sender := &NilSender{}
	m := Emailer{
		Sender:                  sender,
		VerifyEmailTemplate:     "testTemplates/verifyEmail.html",
		VerifyEmailSubject:      "verifyEmailSubject",
		WelcomeTemplate:         "testTemplates/welcomeEmail.html",
		WelcomeSubject:          "welcomeSubject",
		NewLoginTemplate:        "testTemplates/newLogin.html",
		NewLoginSubject:         "newLoginSubject",
		LockedOutTemplate:       "testTemplates/lockedOut.html",
		LockedOutSubject:        "lockedOutSubject",
		EmailChangedTemplate:    "testTemplates/emailChanged.html",
		EmailChangedSubject:     "emailChangedSubject",
		PasswordChangedTemplate: "testTemplates/passwordChanged.html",
		PasswordChangedSubject:  "passwordChangedSubject",
	}
	m.TemplateCache = template.Must(template.ParseFiles(m.VerifyEmailTemplate, m.WelcomeTemplate,
		m.NewLoginTemplate, m.LockedOutTemplate, m.EmailChangedTemplate, m.PasswordChangedTemplate))
	data := &emailSession{Email: "myemail@here.com"}
	err := m.SendVerify("to", data)
	if err != nil || sender.LastBody != "verifyEmail:myemail@here.com" || sender.LastTo != "to" || sender.LastSubject != "verifyEmailSubject" {
		t.Error("expected valid values", sender, err)
	}

	err = m.SendWelcome("to1", data)
	if err != nil || sender.LastBody != "welcomeEmail:myemail@here.com" || sender.LastTo != "to1" || sender.LastSubject != "welcomeSubject" {
		t.Error("expected valid values", sender, err)
	}

	err = m.SendNewLogin("to2", data)
	if err != nil || sender.LastBody != "newLogin:myemail@here.com" || sender.LastTo != "to2" || sender.LastSubject != "newLoginSubject" {
		t.Error("expected valid values", sender, err)
	}

	err = m.SendLockedOut("to3", data)
	if err != nil || sender.LastBody != "lockedOut:myemail@here.com" || sender.LastTo != "to3" || sender.LastSubject != "lockedOutSubject" {
		t.Error("expected valid values", sender, err)
	}

	err = m.SendEmailChanged("to4", data)
	if err != nil || sender.LastBody != "emailChanged:myemail@here.com" || sender.LastTo != "to4" || sender.LastSubject != "emailChangedSubject" {
		t.Error("expected valid values", sender, err)
	}

	err = m.SendPasswordChanged("to5", data)
	if err != nil || sender.LastBody != "passwordChanged:myemail@here.com" || sender.LastTo != "to5" || sender.LastSubject != "passwordChangedSubject" {
		t.Error("expected valid values", sender, err)
	}
}

/***************************************************************************************/

type NilSender struct {
	LastTo      string
	LastSubject string
	LastBody    string
}

func (s *NilSender) Send(to, subject, body string) error {
	s.LastTo = to
	s.LastSubject = subject
	s.LastBody = body
	return nil
}

type TextMailer struct {
	Err error
	Mailer
	MessageTo   string
	MessageData interface{}
}

func (t *TextMailer) SendVerify(to string, data interface{}) error {
	return t.send(to, data)
}

func (t *TextMailer) SendWelcome(to string, data interface{}) error {
	return t.send(to, data)
}

func (t *TextMailer) SendNewLogin(to string, data interface{}) error {
	return t.send(to, data)
}

func (t *TextMailer) SendLockedOut(to string, data interface{}) error {
	return t.send(to, data)
}

func (t *TextMailer) SendEmailChanged(to string, data interface{}) error {
	return t.send(to, data)
}

func (t *TextMailer) SendPasswordChanged(to string, data interface{}) error {
	return t.send(to, data)
}

func (t *TextMailer) send(to string, data interface{}) error {
	t.MessageTo = to
	t.MessageData = data
	return t.Err
}
