TODO : integrate client openstack (https://github.com/gophercloud/gophercloud)

See documentation at https://github.com/cloudquery/cloudquerydoc/blob/master/developers/developing-new-provider.md

# prerequirises

## Dependencies

```bash
go get github.com/gophercloud/gophercloud
go get github.com/tools/godep
godep save ./...
```
## Openstack Environment variables

* OS_PROJECT_DOMAIN_ID
* OS_PROJECT_ID
* OS_REGION_NAME
* OS_USER_DOMAIN_NAME
* OS_PROJECT_NAME
* OS_IDENTITY_API_VERSION
* OS_PASSWORD
* OS_AUTH_URL
* OS_USERNAME
* OS_CACERT
* OS_INTERFACE