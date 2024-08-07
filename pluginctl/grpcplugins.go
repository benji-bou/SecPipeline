package pluginctl

import (
	context "context"
	"errors"

	"github.com/hashicorp/go-plugin"

	grpc "google.golang.org/grpc"
)

var (
	ErrJsonSchemaConvertion error = errors.New("failed to convert to json format")
	ErrJsonConvertion       error = errors.New("failed to convert to/from json")
)

// This is the implementation of plugin.GRPCPlugin so we can serve/consume this.
type SecPipelineGRPCPlugin struct {
	// GRPCPlugin must still implement the Plugin interface
	plugin.NetRPCUnsupportedPlugin
	Impl SecPluginable
	Name string
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
}

func (p SecPipelineGRPCPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterSecPipelinePluginsServer(s, &GRPCServer{
		Impl:        p.Impl,
		Name:        p.Name,
		PluginDataC: make(chan pluginServerDataC),
	})
	return nil
}

func (p SecPipelineGRPCPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		Name:   p.Name,
		client: NewSecPipelinePluginsClient(c),
	}, nil
}
