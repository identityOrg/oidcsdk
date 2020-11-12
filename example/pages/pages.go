package pages

import (
	sdk "github.com/identityOrg/oidcsdk"
	"html/template"
	"net/http"
)

type PageRenderer struct {
}

func NewPageRenderer() *PageRenderer {
	return &PageRenderer{}
}

func (p *PageRenderer) DisplayLogoutConsentPage(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (p *PageRenderer) DisplayLogoutStatusPage(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (p *PageRenderer) DisplayErrorPage(err error, w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (p *PageRenderer) DisplayLoginPage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_ = template.Must(template.New("login").Parse(LoginPage)).Execute(writer, request.URL.String())
}

func (p *PageRenderer) DisplayConsentPage(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_, _ = writer.Write([]byte(ConsentPage))
}

const (
	LoginPage = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Login</title>
</head>
<body>
<form action="/login" method="post">
    <label for="u-name">Username</label>
    <input type="text" name="username" id="u-name">
    <label for="u-pass">Password</label>
    <input type="text" name="password" id="u-pass">
    <input type="submit">
    <input type="hidden" name="request" value="{{.}}">
</form>
</body>
</html>
`
	ConsentPage = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Login</title>
</head>
<body>
<form action="" method="post">
    <p>Consent Required</p>
</form>
</body>
</html>
`
)
