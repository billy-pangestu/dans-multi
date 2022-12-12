package bootstrap

import (
	"be-user-scheme/pkg/jwe"
	"be-user-scheme/pkg/jwt"
	"be-user-scheme/pkg/pg"
	"be-user-scheme/usecase"

	"github.com/go-chi/chi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v7"
	validator "gopkg.in/go-playground/validator.v9"
)

// Bootup ...
type Bootup struct {
	R          *chi.Mux
	CorsDomain []string
	EnvConfig  map[string]string
	DB         *pg.MySQL
	Redis      *redis.Client
	Validator  *validator.Validate
	Translator ut.Translator
	ContractUC usecase.ContractUC
	Jwt        jwt.Credential
	Jwe        jwe.Credential
}
