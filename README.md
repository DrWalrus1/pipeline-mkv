https://makemkv.com/developers/usage.txt <- Link for makemkvcon documentation

MkvToolNix documentation
* https://mkvtoolnix.download/doc/mkvmerge.html
* https://mkvtoolnix.download/doc/mkvextract.html
* https://mkvtoolnix.download/doc/mkvinfo.html
* https://mkvtoolnix.download/doc/mkvpropedit.html

## Useful commands

* create docker buildx to handle multiple targets `docker buildx create --use desktop-linux`
* docker build for arm64 and amd64 `docker buildx build . -t thedrwalrus/makemkv --platform linux/amd64,linux/arm64 --push`
* command used to get outputs.txt `makemkvcon -r --cache=1 info disc:0`
* command example to rip with progress output `makemkvcon -r --progress=-stdout mkv disc:0 4 ./`
* debug in container:
```bash
docker pull thedrwalrus/makemkv && \
docker run --rm -it \
--device=/dev/sr0:/dev/sr0 \
--device=/dev/sg0:/dev/sg0 \
--mount type=bind,src=/mnt/bigboy/Plex,dst=/code/Plex \
-p 9090:8080 \
--name sandbox \
thedrwalrus/makemkv /bin/bash
```
* debug in the container `docker pull thedrwalrus/makemkv && docker run --rm -it --device=/dev/sr0:/dev/sr0 --device=/dev/sg0:/dev/sg0 -p 9090:8080 --name sandbox thedrwalrus/makemkv /bin/bash`

Changes to master will create/push a new Docker image 

## TODO
* Make UI
* when api starts up we need to run makemkv info and load all detectable disc locations. these only reload on restart and connecting devices seems rare and worth the restart
* Need to figure out a way for the UI to show the a backup locations. We want to limit 
* The UI needs an advanced mode where the user can manually pass in device info to read from /dev/...
* need to create a responsive progress bar for when ripping the movie
* maybe allow browser notifications for notifying the user when the rip is done
