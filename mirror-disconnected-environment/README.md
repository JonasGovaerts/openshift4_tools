Go script to mirror operator images for catalog indexes to a local registry

Parameters:
- -h	show help
- -v	version of the catalog index that you want to mirror (e.g.: v4.x)
- -o	operators that you want to mirror (certified-operator-index,community-operator-index,redhat-operator-index)
- -l	list the packages available in the specified operator indexes - value "true"
- -r	local registry to mirror images to
- -a	file to authenticate to local and upstream registries
- -p	comma seperated list of packages that you want to mirror
