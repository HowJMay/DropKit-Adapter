package health

import (
	"github.com/DropKit/DropKit-Adapter/configs"
	"github.com/DropKit/DropKit-Adapter/pkg/logging"
)

var logger logging.Logger

func init() {
	logger, _ = logging.NewLogger(configs.Configs.Adapter.LoggerLevel, "health")
}
