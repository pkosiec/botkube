package source

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/kubeshop/botkube/internal/httpx"
	"github.com/kubeshop/botkube/pkg/config"
	"github.com/kubeshop/botkube/pkg/multierror"
)

func NewIncomingWebhookServer(log logrus.FieldLogger, cfg *config.Config, dispatcher *Dispatcher, startedSources map[string][]StartedSource) *httpx.Server {
	addr := fmt.Sprintf(":%d", cfg.IncomingWebhook.Port)
	router := incomingWebhookRouter(log, cfg, dispatcher, startedSources)

	log.Info("Starting server on %q...", addr)
	return httpx.NewServer(log, addr, router)
}

const (
	sourceNameVarName = "sourceName"
)

func incomingWebhookRouter(log logrus.FieldLogger, cfg *config.Config, dispatcher *Dispatcher, startedSources map[string][]StartedSource) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(fmt.Sprintf("/sources/v1/{%s}", sourceNameVarName), func(writer http.ResponseWriter, request *http.Request) {
		sourceName, ok := mux.Vars(request)[sourceNameVarName]
		if !ok {
			http.Error(writer, "sourceName is required", http.StatusBadRequest)
			return
		}
		logger := log.WithFields(logrus.Fields{
			"sourceName": sourceName,
		})
		logger.Debugf("Handling incoming webhook request...")

		sourcePlugins, ok := startedSources[sourceName]
		if !ok {
			http.Error(writer, fmt.Sprintf("source %q not found", sourceName), http.StatusNotFound)
			return
		}

		payload, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, fmt.Sprintf("while reading request body: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		defer request.Body.Close()

		multiErr := multierror.New()
		for _, src := range sourcePlugins {
			logger.WithFields(logrus.Fields{
				"pluginName":               src.PluginName,
				"isInteractivitySupported": src.IsInteractivitySupported,
			}).Debug("Dispatching message...")

			err := dispatcher.DispatchSingle(SinglePluginDispatch{
				PluginDispatch: PluginDispatch{
					ctx:                      request.Context(),
					sourceName:               sourceName,
					sourceDisplayName:        src.SourceDisplayName,
					pluginName:               src.PluginName,
					pluginConfig:             src.PluginConfig,
					isInteractivitySupported: src.IsInteractivitySupported,
					cfg:                      cfg,
					pluginContext:            config.PluginContext{},
				},
				payload: payload,
			})
			if err != nil {
				multiErr = multierror.Append(multiErr, fmt.Errorf(`while dispatching message for "%s.%s": %w`, sourceName, src.PluginName, err))
			}
		}

		if multiErr.ErrorOrNil() != nil {
			wrappedErr := fmt.Errorf("while dispatching message: %w", multiErr)
			http.Error(writer, wrappedErr.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods(http.MethodPost)
	return router
}
