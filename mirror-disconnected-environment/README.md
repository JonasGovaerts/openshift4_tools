## Summary
Go script to mirror operator images for catalog indexes to a local registry. This is used when you are working in a different environment.
Script is based on the latest documentation of RedHat (https://docs.openshift.com/container-platform/latest/operators/admin/olm-restricted-networks.html#olm-pruning-index-image_olm-restricted-networks)

## Requirements
- `latest OC binary (https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/latest/openshift-client-linux.tar.gz)
- `latest OPM binary (https://mirror.openshift.com/pub/openshift-v4/x86_64/clients/ocp/latest/opm-linux.tar.gz)
-  Podman v1.9.3+
-  Access to a registry that supports multi arch images


### Parameters
**`-help`** shows the help menu

**`-creds`** location to authentication file

**`-operator`** operator index that you want to mirror

**`-packages`**  comma seperated list of packages that need to be mirrored

**`-local-operator`** target to push the modified operator index to

**`-list`** list packages available in the operator index

**`-loglevel`** set log level: debug, info, warn (default "info")
