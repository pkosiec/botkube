package plugin

import (
	"github.com/kubeshop/botkube/internal/httpx"
	"github.com/kubeshop/botkube/pkg/config"
	"github.com/sirupsen/logrus"
)

type IncomingWebhookServer struct {
	log logrus.FieldLogger
	cfg config.IncomingWebhook
}

func NewIncomingWebhookServer(log logrus.FieldLogger, cfg config.IncomingWebhook) *httpx.Server {
	return &IncomingWebhookServer{log: log, cfg: cfg}

	return httpx.NewServer(log, addr, router)
}
