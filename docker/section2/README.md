# Manipulating Docker containers with docker client

## docker run hello-world

## Override command
`$ docker run busybox echo hi there`

- Can run echo command due to `busybox` image contains the command

## Life cycle of container
`$ docker run <image name>`

`$ docker run` = `docker create` + `docker start`

### `$ docker create`
  Just create the image. It means copied File snapshot

* can override startup command

### `$ docker start -a`

* `-a` show all the logs

### docker run busybox echo hi there
It creates container with file system snapshot and add start up command `echo hi there`
* `$ docker create busybox echo hi there`
* `$ docker start <created container id>`

## Command
### `$ docker system prune`
To clean all stopped containers, unused networks, dangling build cache and dangling images

### `$ docker logs <container id>`
Can logs output of container.  Then we don't need to run `docker start` with `-a` flag to display output


### `$ docker stop <container Id>`
Send `SIGTERM` so container process has time to do finishing tasks

* docker will wait 10 sec and will send `SIGKILL` in the end.

### `$ docker kill <container id>`
Send `SIGKILL` Stop immediately

### `$ docker exec -it <container id> <command>`
- `-it` flag allow tus to provide input to the container

#### `$ docker exec -it <container id> sh`
- sh is shell program like bash or powershell
- Most of containers included `sh`


## Docker container is isolated
