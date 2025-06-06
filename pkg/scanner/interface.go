package scanner

import (
	"github.com/solaire/genie/pkg/models"
	"github.com/solaire/genie/pkg/scanner/battlenet"
	"github.com/solaire/genie/pkg/scanner/ea"
	"github.com/solaire/genie/pkg/scanner/epic"
	"github.com/solaire/genie/pkg/scanner/gog"
	"github.com/solaire/genie/pkg/scanner/steam"
	"github.com/solaire/genie/pkg/scanner/ubisoft"
)

type Scanner interface {
	Name() string
	Detect() bool
	ScanGames() ([]models.Game, error)
}

var ScannerList = []Scanner{
	&steam.Scanner{},
	&gog.Scanner{},
	&ubisoft.Scanner{},
	&epic.Scanner{},
	&ea.Scanner{},
	&battlenet.Scanner{},
}
