package resources

import (
	"github.com/cloudquery/cq-provider-openstack/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "openstack",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"openstack.projects": OpenstackProject(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
