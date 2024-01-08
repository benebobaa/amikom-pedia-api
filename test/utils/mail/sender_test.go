package mail

import (
	"amikom-pedia-api/helper"
	"amikom-pedia-api/utils/mail"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSendEmailWithGmail(t *testing.T) {

	if testing.Short() {
		t.Skip("Skip test because it's short")
	}

	sender := mail.NewGmailSender("Amikom Pedia", "amikompedia@gmail.com", "djib wmwf kdnp hhek")

	subject := "A email from AmiPed Teams!"

	content := `
		<h1> Hello, this is a announce email from Amikom Pedia! </h1>
		<p> If you can read this email, it means that you are an eligible user to join Amikom Pedia. We are still in the process of developing our app with our teams. I hope you can patiently wait, and see you in the future. </p>
	`

	to := []string{"benediktus@students.amikom.ac.id"}
	attachFiles := []string{"../../README.md"}

	err := sender.SendEmail(subject, content, to, []string{}, []string{}, attachFiles)
	helper.PanicIfError(err)
	require.NoError(t, err)
}
