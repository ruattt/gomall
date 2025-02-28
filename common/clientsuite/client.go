package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	// "github.com/hertz-contrib/obs-opentelemetry/provider"
	// "github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

type CommonClientSuite struct {
	DestServiceName    string
	DestServiceAddr    string
	CurrentServiceName string
	TracerProvider     *tracesdk.TracerProvider
	RegistryAddr       string
}

func (s CommonClientSuite) Options() []client.Option {
	opts := []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
	}

	// _ = provider.NewOpenTelemetryProvider(provider.WithSdkTracerProvider(s.TracerProvider), provider.WithEnableMetrics(false))
	// opts = append(opts,
	// 	client.WithSuite(tracing.NewClientSuite()),
	// )

	r, err := consul.NewConsulResolver(s.RegistryAddr)
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))



	return opts
}
