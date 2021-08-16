package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	// This is a client that you need to create and initialize in Configure
	// It will be passed for each resource fetcher.
	logger hclog.Logger

	// Usually you store here your 3rd party clients and use them in the fetcher
	ThirdPartyClient interface{}
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

func Init(logger hclog.Logger) {

	// Use a utility function to retrieve all your environment variables
	opts, err := openstack.AuthOptionsFromEnv()

	// Once you have an opts variable, you can pass it in and get back a ProviderClient struct
	provider, err := openstack.AuthenticatedClient(opts)

	client, err := openstack.NewIdentityV3(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		panic(err)
	}

	var iTrue bool = true
	listOpts := projects.ListOpts{
		Enabled: &iTrue,
	}

	allPages, err := projects.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allProjects, err := projects.ExtractProjects(allPages)
	if err != nil {
		panic(err)
	}

	for _, project := range allProjects {
		logger.Info("project:", project)
	}

}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	providerConfig := config.(*Config)
	_ = providerConfig
	// Init your client and 3rd party clients using the user's configuration
	// passed by the SDK providerConfig
	client := Client{
		logger:           logger,
		ThirdPartyClient: nil,
	}
	logger.Info("Configure()")
	Init(logger)
	// Return the initialized client and it will be passed to your resources
	return &client, nil
}
