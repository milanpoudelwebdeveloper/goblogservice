package templates

import "fmt"

// verifyEmailTemplate generates an email template for account verification.
func VerifyEmailTemplate(clientUrl, token string) string {
	return fmt.Sprintf(`
<!doctype html>
<html>
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Verify Your Account</title>
  </head>
  <body style="margin: 0; padding: 0; background-color: #f7f7f7; font-family: Arial, sans-serif">
    <table
      align="center"
      border="0"
      cellpadding="0"
      cellspacing="0"
      width="100%%"
      style="
        max-width: 600px;
        margin: 30px auto;
        background-color: #ffffff;
        padding: 20px;
        border-radius: 10px;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      "
    >
      <tr>
        <td
          align="center"
          style="background-color: #4caf50; color: white; padding: 10px; border-top-left-radius: 10px; border-top-right-radius: 10px"
        >
          <h1 style="margin: 0; font-size: 24px">Welcome to Code With Milan!</h1>
        </td>
      </tr>
      <tr>
        <td style="padding: 20px; line-height: 1.6; color: #333">
          <p>Dear User,</p>
          <p>
            Thank you for signing up with Code With Milan. To complete your registration, please verify your account by clicking the button
            below:
          </p>
          <p style="text-align: center">
            <a
              href="%s/verify-account?token=%s"
              style="
                display: inline-block;
                background-color: #4caf50;
                color: white;
                padding: 10px 20px;
                text-decoration: none;
                border-radius: 5px;
                font-size: 16px;
              "
              >Verify Account</a
            >
          </p>
          <p>If you did not sign up for this account, please disregard this email.</p>
          <p>Best regards,<br />Code With Milan</p>
        </td>
      </tr>
      <tr>
        <td style="text-align: center; color: #777; font-size: 12px; margin-top: 20px">
          <p>&copy; 2024 Code With Milan. All rights reserved.</p>
        </td>
      </tr>
    </table>
  </body>
</html>
`, clientUrl, token)
}

// forgotPasswordTemplate generates an email template for password reset.
func ForgotPasswordTemplate(clientUrl, token string) string {
	return fmt.Sprintf(`
<!doctype html>
<html>
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Verify your account to reset your password</title>
  </head>
  <body style="margin: 0; padding: 0; background-color: #f7f7f7; font-family: Arial, sans-serif">
    <table
      align="center"
      border="0"
      cellpadding="0"
      cellspacing="0"
      width="100%%"
      style="
        max-width: 600px;
        margin: 30px auto;
        background-color: #ffffff;
        padding: 20px;
        border-radius: 10px;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      "
    >
      <tr>
        <td
          align="center"
          style="background-color: #4caf50; color: white; padding: 10px; border-top-left-radius: 10px; border-top-right-radius: 10px"
        >
          <h1 style="margin: 0; font-size: 24px">Welcome to Code With Milan!</h1>
        </td>
      </tr>
      <tr>
        <td style="padding: 20px; line-height: 1.6; color: #333">
          <p>Dear User,</p>
          <p>
            We received a request to reset your password. To complete the process, please click the button below to verify your account first:
          </p>
          <p style="text-align: center">
            <a
              href="%s/verify-password-reset?token=%s"
              style="
                display: inline-block;
                background-color: #4caf50;
                color: white;
                padding: 10px 20px;
                text-decoration: none;
                border-radius: 5px;
                font-size: 16px;
              "
              >Password Reset Link</a
            >
          </p>
          <p>If you did not request to reset your password, please disregard this email.</p>
          <p>Best regards,<br />Code With Milan</p>
        </td>
      </tr>
      <tr>
        <td style="text-align: center; color: #777; font-size: 12px; margin-top: 20px">
          <p>&copy; 2024 Code With Milan. All rights reserved.</p>
        </td>
      </tr>
    </table>
  </body>
</html>
`, clientUrl, token)
}
