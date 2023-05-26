package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"goyave.dev/goyave/v4"
	"goyave.dev/goyave/v4/config"
	"goyave.dev/goyave/v4/util/testutil"
)

func prepareJWTServiceTest() (*testutil.TestServer, *JWTService) {
	server := testutil.NewTestServerWithConfig(config.LoadDefault(), nil)
	service := &JWTService{}
	server.RegisterService(service)
	return server, service
}

func TestJWTService(t *testing.T) {

	t.Run("GenerateToken", func(t *testing.T) {
		server, service := prepareJWTServiceTest()
		server.Config().Set("auth.jwt.secret", "secret")
		server.Config().Set("auth.jwt.expiry", 20)

		now := time.Now()
		expiry := time.Duration(20) * time.Second

		tokenString, err := service.GenerateToken("johndoe")
		if !assert.NoError(t, err) {
			return
		}
		parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(server.Config().GetString("auth.jwt.secret")), nil
		})

		assert.NoError(t, err)
		assert.True(t, parsedToken.Valid)
		assert.Equal(t, jwt.SigningMethodHS256, parsedToken.Method)

		assert.NoError(t, parsedToken.Claims.Valid())
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if assert.True(t, ok) {
			assert.Equal(t, "johndoe", claims["sub"])
			assert.GreaterOrEqual(t, float64(now.Unix()), claims["nbf"])
			assert.True(t, time.Unix(int64(claims["exp"].(float64)), 0).After(now))
			assert.Equal(t, int64(expiry.Seconds()), int64(claims["exp"].(float64)-claims["nbf"].(float64)))
		}
	})

	t.Run("GenerateTokenWithClaims_HS256", func(t *testing.T) {
		server, service := prepareJWTServiceTest()
		server.Config().Set("auth.jwt.secret", "secret")
		server.Config().Set("auth.jwt.expiry", 20)

		now := time.Now()
		expiry := time.Duration(20) * time.Second

		srcClaims := jwt.MapClaims{
			"sub":         "johndoe",
			"customClaim": "customValue",
		}
		tokenString, err := service.GenerateTokenWithClaims(srcClaims, jwt.SigningMethodHS256)
		if !assert.NoError(t, err) {
			return
		}
		parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(server.Config().GetString("auth.jwt.secret")), nil
		})

		assert.NoError(t, err)
		assert.True(t, parsedToken.Valid)
		assert.Equal(t, jwt.SigningMethodHS256, parsedToken.Method)

		assert.NoError(t, parsedToken.Claims.Valid())
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if assert.True(t, ok) {
			assert.Equal(t, "johndoe", claims["sub"])
			assert.Equal(t, "customValue", claims["customClaim"])
			assert.GreaterOrEqual(t, float64(now.Unix()), claims["nbf"])
			assert.True(t, time.Unix(int64(claims["exp"].(float64)), 0).After(now))
			assert.Equal(t, int64(expiry.Seconds()), int64(claims["exp"].(float64)-claims["nbf"].(float64)))
		}
	})

	t.Run("GenerateTokenWithClaims_RSA", func(t *testing.T) {
		rootDir := testutil.FindRootDirectory()
		server, service := prepareJWTServiceTest()
		server.Config().Set("auth.jwt.rsa.public", rootDir+"resources/rsa/public.pem")
		server.Config().Set("auth.jwt.rsa.private", rootDir+"resources/rsa/private.pem")
		server.Config().Set("auth.jwt.expiry", 20)

		now := time.Now()
		expiry := time.Duration(20) * time.Second

		srcClaims := jwt.MapClaims{
			"sub":         "johndoe",
			"customClaim": "customValue",
		}
		tokenString, err := service.GenerateTokenWithClaims(srcClaims, jwt.SigningMethodRS256)
		if !assert.NoError(t, err) {
			return
		}
		parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return service.GetKey("auth.jwt.rsa.public")
		})

		assert.NoError(t, err)
		assert.True(t, parsedToken.Valid)
		assert.Equal(t, jwt.SigningMethodRS256, parsedToken.Method)

		assert.NoError(t, parsedToken.Claims.Valid())
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if assert.True(t, ok) {
			assert.Equal(t, "johndoe", claims["sub"])
			assert.Equal(t, "customValue", claims["customClaim"])
			assert.GreaterOrEqual(t, float64(now.Unix()), claims["nbf"])
			assert.True(t, time.Unix(int64(claims["exp"].(float64)), 0).After(now))
			assert.Equal(t, int64(expiry.Seconds()), int64(claims["exp"].(float64)-claims["nbf"].(float64)))
		}
	})

	t.Run("GenerateTokenWithClaims_ECDSA", func(t *testing.T) {
		rootDir := testutil.FindRootDirectory()
		server, service := prepareJWTServiceTest()
		server.Config().Set("auth.jwt.ecdsa.public", rootDir+"resources/ecdsa/public.pem")
		server.Config().Set("auth.jwt.ecdsa.private", rootDir+"resources/ecdsa/private.pem")
		server.Config().Set("auth.jwt.expiry", 20)

		now := time.Now()
		expiry := time.Duration(20) * time.Second

		srcClaims := jwt.MapClaims{
			"sub":         "johndoe",
			"customClaim": "customValue",
		}
		tokenString, err := service.GenerateTokenWithClaims(srcClaims, jwt.SigningMethodES256)
		if !assert.NoError(t, err) {
			return
		}
		parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return service.GetKey("auth.jwt.ecdsa.public")
		})

		assert.NoError(t, err)
		assert.True(t, parsedToken.Valid)
		assert.Equal(t, jwt.SigningMethodES256, parsedToken.Method)

		assert.NoError(t, parsedToken.Claims.Valid())
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if assert.True(t, ok) {
			assert.Equal(t, "johndoe", claims["sub"])
			assert.Equal(t, "customValue", claims["customClaim"])
			assert.GreaterOrEqual(t, float64(now.Unix()), claims["nbf"])
			assert.True(t, time.Unix(int64(claims["exp"].(float64)), 0).After(now))
			assert.Equal(t, int64(expiry.Seconds()), int64(claims["exp"].(float64)-claims["nbf"].(float64)))
		}
	})

	t.Run("GenerateTokenWithClaims_Unsupported", func(t *testing.T) {
		server, service := prepareJWTServiceTest()
		server.Config().Set("auth.jwt.expiry", 20)

		_, err := service.GenerateTokenWithClaims(nil, jwt.SigningMethodPS256)
		assert.Error(t, err)
	})

	t.Run("GetPrivateKey", func(t *testing.T) {
		// TODO
	})
}

func TestJWTAuthenticator(t *testing.T) {

	t.Run("success_hs256", func(t *testing.T) {
		server, user := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{})

		// No need to register the JWTService, it should be done automatically
		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateToken(user.Email)
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Equal(t, user.ID, request.User.(*TestUser).ID)
			assert.Equal(t, user.Name, request.User.(*TestUser).Name)
			assert.Equal(t, user.Email, request.User.(*TestUser).Email)
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("success_rsa", func(t *testing.T) {
		rootDir := testutil.FindRootDirectory()
		server, user := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.rsa.public", rootDir+"resources/rsa/public.pem")
		server.Config().Set("auth.jwt.rsa.private", rootDir+"resources/rsa/private.pem")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{SigningMethod: jwt.SigningMethodRS256})

		// No need to register the JWTService, it should be done automatically
		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateTokenWithClaims(jwt.MapClaims{"sub": user.Email}, jwt.SigningMethodRS256)
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Equal(t, user.ID, request.User.(*TestUser).ID)
			assert.Equal(t, user.Name, request.User.(*TestUser).Name)
			assert.Equal(t, user.Email, request.User.(*TestUser).Email)
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("success_ecdsa", func(t *testing.T) {
		rootDir := testutil.FindRootDirectory()
		server, user := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.ecdsa.public", rootDir+"resources/ecdsa/public.pem")
		server.Config().Set("auth.jwt.ecdsa.private", rootDir+"resources/ecdsa/private.pem")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{SigningMethod: jwt.SigningMethodES256})

		// No need to register the JWTService, it should be done automatically
		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateTokenWithClaims(jwt.MapClaims{"sub": user.Email}, jwt.SigningMethodES256)
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Equal(t, user.ID, request.User.(*TestUser).ID)
			assert.Equal(t, user.Name, request.User.(*TestUser).Name)
			assert.Equal(t, user.Email, request.User.(*TestUser).Email)
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("invalid_token", func(t *testing.T) {
		server, _ := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{})

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer invalidtoken")
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Nil(t, request.User)
			assert.Fail(t, "middleware passed despite failed authentication")
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		body, err := testutil.ReadJSONBody[map[string]string](resp.Body)
		_ = resp.Body.Close()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, map[string]string{"error": server.Lang.GetDefault().Get("auth.jwt-invalid")}, body)
	})

	t.Run("token_not_valid_yet", func(t *testing.T) {
		server, user := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{})

		// No need to register the JWTService, it should be done automatically
		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateTokenWithClaims(jwt.MapClaims{
			"sub": user.Email,
			"nbf": time.Now().Add(time.Hour).Unix(),
		}, jwt.SigningMethodHS256)
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Nil(t, request.User)
			assert.Fail(t, "middleware passed despite failed authentication")
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		body, err := testutil.ReadJSONBody[map[string]string](resp.Body)
		_ = resp.Body.Close()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, map[string]string{"error": server.Lang.GetDefault().Get("auth.jwt-not-valid-yet")}, body)
	})

	t.Run("token_expired", func(t *testing.T) {
		server, user := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{})

		// No need to register the JWTService, it should be done automatically
		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateTokenWithClaims(jwt.MapClaims{
			"sub": user.Email,
			"exp": time.Now().Add(-time.Hour).Unix(),
		}, jwt.SigningMethodHS256)
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Nil(t, request.User)
			assert.Fail(t, "middleware passed despite failed authentication")
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		body, err := testutil.ReadJSONBody[map[string]string](resp.Body)
		_ = resp.Body.Close()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, map[string]string{"error": server.Lang.GetDefault().Get("auth.jwt-expired")}, body)
	})

	t.Run("unknown_user", func(t *testing.T) {
		server, _ := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{})

		// No need to register the JWTService, it should be done automatically
		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateToken("notjohndoe@example.org")
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Nil(t, request.User)
			assert.Fail(t, "middleware passed despite failed authentication")
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		body, err := testutil.ReadJSONBody[map[string]string](resp.Body)
		_ = resp.Body.Close()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, map[string]string{"error": server.Lang.GetDefault().Get("auth.invalid-credentials")}, body)
	})

	t.Run("error_no_table", func(t *testing.T) {
		cfg := config.LoadDefault()
		cfg.Set("database.connection", "sqlite3")
		cfg.Set("database.name", "testjwtauthenticator_no_table.db")
		cfg.Set("database.options", "mode=memory")
		cfg.Set("auth.jwt.secret", "secret")
		cfg.Set("app.debug", false)
		server := testutil.NewTestServerWithConfig(cfg, nil)

		// No need to register the JWTService, it should be done automatically
		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateToken("johndoe@example.org")
		if !assert.NoError(t, err) {
			return
		}

		authenticator := Middleware[*TestUser](&JWTAuthenticator{})
		authenticator.Init(server.Server)
		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		assert.Panics(t, func() {
			user := &TestUserPromoted{}
			_ = authenticator.Authenticate(request, &user)
		})
	})

	t.Run("unexpected_method_hmac", func(t *testing.T) {
		rootDir := testutil.FindRootDirectory()
		server, _ := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.rsa.public", rootDir+"resources/rsa/public.pem")
		server.Config().Set("auth.jwt.rsa.private", rootDir+"resources/rsa/private.pem")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{SigningMethod: jwt.SigningMethodHS256})

		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateTokenWithClaims(jwt.MapClaims{"sub": "johndoe@example.org"}, jwt.SigningMethodRS256)
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Nil(t, request.User)
			assert.Fail(t, "middleware passed despite failed authentication")
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		body, err := testutil.ReadJSONBody[map[string]string](resp.Body)
		_ = resp.Body.Close()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, map[string]string{"error": server.Lang.GetDefault().Get("auth.jwt-invalid")}, body)
	})

	t.Run("unexpected_method_rsa", func(t *testing.T) {
		server, _ := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{SigningMethod: jwt.SigningMethodRS256})

		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateToken("johndoe@example.org")
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Nil(t, request.User)
			assert.Fail(t, "middleware passed despite failed authentication")
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		body, err := testutil.ReadJSONBody[map[string]string](resp.Body)
		_ = resp.Body.Close()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, map[string]string{"error": server.Lang.GetDefault().Get("auth.jwt-invalid")}, body)
	})

	t.Run("unexpected_method_ecdsa", func(t *testing.T) {
		server, _ := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{SigningMethod: jwt.SigningMethodES256})

		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateToken("johndoe@example.org")
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Nil(t, request.User)
			assert.Fail(t, "middleware passed despite failed authentication")
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		body, err := testutil.ReadJSONBody[map[string]string](resp.Body)
		_ = resp.Body.Close()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, map[string]string{"error": server.Lang.GetDefault().Get("auth.jwt-invalid")}, body)
	})

	t.Run("unsupported_method", func(t *testing.T) {
		server, _ := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{SigningMethod: jwt.SigningMethodPS256})

		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateToken("johndoe@example.org")
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		assert.Panics(t, func() {
			user := &TestUser{}
			_ = authenticator.Authenticate(request, &user)
		})
	})

	t.Run("no_auth", func(t *testing.T) {
		server, _ := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{})

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Nil(t, request.User)
			assert.Fail(t, "middleware passed despite failed authentication")
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		body, err := testutil.ReadJSONBody[map[string]string](resp.Body)
		_ = resp.Body.Close()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, map[string]string{"error": server.Lang.GetDefault().Get("auth.no-credentials-provided")}, body)
	})

	t.Run("optional_success", func(t *testing.T) {
		server, user := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{Optional: true})

		// No need to register the JWTService, it should be done automatically
		service := &JWTService{}
		service.Init(server.Server)

		token, err := service.GenerateToken(user.Email)
		if !assert.NoError(t, err) {
			return
		}

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer "+token)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Equal(t, user.ID, request.User.(*TestUser).ID)
			assert.Equal(t, user.Name, request.User.(*TestUser).Name)
			assert.Equal(t, user.Email, request.User.(*TestUser).Email)
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		_ = resp.Body.Close()
	})

	t.Run("optional_invalid_token", func(t *testing.T) {
		server, _ := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{Optional: true})

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		request.Request().Header.Set("Authorization", "Bearer invalidtoken")
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Nil(t, request.User)
			assert.Fail(t, "middleware passed despite failed authentication")
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		body, err := testutil.ReadJSONBody[map[string]string](resp.Body)
		_ = resp.Body.Close()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, map[string]string{"error": server.Lang.GetDefault().Get("auth.jwt-invalid")}, body)
	})

	t.Run("optional_no_auth", func(t *testing.T) {
		server, _ := prepareAuthenticatorTest()
		server.Config().Set("auth.jwt.secret", "secret")
		authenticator := Middleware[*TestUser](&JWTAuthenticator{Optional: true})

		request := server.NewTestRequest(http.MethodGet, "/protected", nil)
		resp := server.TestMiddleware(authenticator, request, func(response *goyave.ResponseV5, request *goyave.RequestV5) {
			assert.Nil(t, request.User)
			response.Status(http.StatusOK)
		})
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		_ = resp.Body.Close()
	})
}
