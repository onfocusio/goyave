package auth

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/System-Glitch/goyave/v3"
	"github.com/System-Glitch/goyave/v3/config"
	"github.com/System-Glitch/goyave/v3/database"
	"github.com/System-Glitch/goyave/v3/lang"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// BasicAuthenticator implementation of Authenticator with the Basic
// authentication method.
type BasicAuthenticator struct{}

var _ Authenticator = (*BasicAuthenticator)(nil) // implements Authenticator

// Authenticate fetch the user corresponding to the credentials
// found in the given request and puts the result in the given user pointer.
// If no user can be authenticated, returns an error.
//
// The database request is executed based on the model name and the
// struct tags `auth:"username"` and `auth:"password"`.
// The password is checked using bcrypt. The username field should unique.
func (a *BasicAuthenticator) Authenticate(request *goyave.Request, user interface{}) error {
	username, password, ok := request.BasicAuth()

	if !ok {
		return fmt.Errorf(lang.Get(request.Lang, "auth.no-credentials-provided"))
	}

	columns := FindColumns(user, "username", "password")

	result := database.GetConnection().Where(columns[0].Name+" = ?", username).First(user)
	notFound := errors.Is(result.Error, gorm.ErrRecordNotFound)

	if result.Error != nil && !notFound {
		panic(result.Error)
	}

	pass := reflect.Indirect(reflect.ValueOf(user)).FieldByName(columns[1].Field.Name)

	if notFound || bcrypt.CompareHashAndPassword([]byte(pass.String()), []byte(password)) != nil {
		return fmt.Errorf(lang.Get(request.Lang, "auth.invalid-credentials"))
	}

	return nil
}

//--------------------------------------------

func init() {
	config.Register("auth.basic.username", config.Entry{
		Value:            nil,
		Type:             reflect.String,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})
	config.Register("auth.basic.password", config.Entry{
		Value:            nil,
		Type:             reflect.String,
		IsSlice:          false,
		AuthorizedValues: []interface{}{},
	})
}

// BasicUser a simple user for config-based basic authentication.
type BasicUser struct {
	Name string
}

type basicUserAuthenticator struct{}

var _ Authenticator = (*basicUserAuthenticator)(nil) // implements Authenticator

// Authenticate check if the request basic auth header matches the
// "auth.basic.username" and "auth.basic.password" config entries.
func (a *basicUserAuthenticator) Authenticate(request *goyave.Request, user interface{}) error {
	username, password, ok := request.BasicAuth()

	if !ok ||
		username != config.GetString("auth.basic.username") ||
		password != config.GetString("auth.basic.password") {
		return fmt.Errorf(lang.Get(request.Lang, "auth.invalid-credentials"))
	}
	user.(*BasicUser).Name = username
	return nil
}

// ConfigBasicAuth create a new authenticator middleware for
// config-based Basic authentication. On auth success, the request
// user is set to a "BasicUser".
// The user is authenticated if the "auth.basic.username" and "auth.basic.password" config entries
// match the request's Authorization header.
func ConfigBasicAuth() goyave.Middleware {
	return Middleware(&BasicUser{}, &basicUserAuthenticator{})
}
