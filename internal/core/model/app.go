package model

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

type App struct {
	Timezone *time.Location
	Server   *http.Server
	Log      zerolog.Logger
}
