package pages

import (
	sdk "github.com/identityOrg/oidcsdk"
	"html/template"
	"net/http"
	"net/url"
)

type PageRenderer struct {
}

func NewPageRenderer() *PageRenderer {
	return &PageRenderer{}
}

type LCModel struct {
	CSRFToken string
	Params    url.Values
}

func (p *PageRenderer) DisplayLogoutConsentPage(writer http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	params := LCModel{
		CSRFToken: "asdd",
		Params:    r.Form,
	}
	writer.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_ = template.Must(template.New("logout_consent").Parse(LogoutConsent)).Execute(writer, params)
}

func (p *PageRenderer) DisplayLogoutStatusPage(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_, _ = writer.Write([]byte(LogoutStatus))
}

func (p *PageRenderer) DisplayErrorPage(err error, writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set(sdk.HeaderContentType, sdk.ContentTypeHtml)
	writer.WriteHeader(200)
	_ = template.Must(template.New("error").Parse(ErrorPage)).Execute(writer, err)
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
	LogoutConsent = `
<!doctype html>
<html lang="en">
<head>
    <title>Error</title>
</head>
<body>
<h2>Logout has been initiated without id_token</h2>
<form action="/oauth2/logout" method="post">
    <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
    {{range $i, $v := .Params}}
    <input type="hidden" name="{{$i}}" value="{{$v}}">
    {{end}}
    <input type="submit" value="Logout">
</form>
</body>
</html>
`
	LogoutStatus = `
<!doctype html>
<html lang="en">
<head>
    <title>Error</title>
</head>
<body>
<h2>You have been logged out</h2>
</body>
</html>
`
	ErrorPage = `
<!doctype html>
<html lang="en">
<head>
    <title>Error</title>
</head>
<body>
<h2>An error has occurred</h2>
{{if .Name}}
<h2>{{.Name}}</h2>
<h3>{{.Description}}</h3>
<h4>{{.Hint}}</h4>
{{else}}
<h3>{{.}}</h3>
{{end}}
</body>
</html>
`
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
