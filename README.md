https://makemkv.com/developers/usage.txt <- Link for makemkvcon documentation

## Useful commands

* create docker buildx to handle multiple targets `docker buildx create --use desktop-linux`
* docker build for arm64 and amd64 `docker buildx build . -t thedrwalrus/makemkv --platform linux/amd64,linux/arm64 --push`
* command used to get outputs.txt `makemkvcon -r --cache=1 info disc:0`
* command example to rip with progress output `makemkvcon -r --progress=-stdout mkv disc:0 4 ./`

Changes to master will create/push a new Docker image 
