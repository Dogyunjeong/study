## DOCKER

### Docker process

#### Build Image
- build with CLI
- build with docker file

#### Run Image

### Docker file

#### build [#reference](https://docs.docker.com/engine/reference/commandline/build/)
```
$ docker build --tag whenry/fedora-jboss:latest -t whenry/fedora-jboss:v2.1 .
```
- `--tag` can be a `-t` 


## Docker compose

- Separate CLI that gets installed along with docker
- Used to start up multiple Docker containers at the same time
- Automates some of the long-winded arguments we were passing to 'docker run'

### Run services with Docker compose
- `docker-compose up` is like to `docker run myimage`
- `docker-compose up --build` is like to `docker build . \ docker run myimage`

#### more

- Launch in background
    `docker-compose up -d`
- Stop containers
    `docker-compose down`

### Service maintenance

#### Check status

- `docker-compose ps`: it must run in the same folder of `docker-compose up`

#### Restart Policies
- "no": Never attempt to restart this . container if it stops or crashes
- always: If this container stops "for any reason" always attempt to restart it
- on-failure: only restart if the container stops with an error code
- unless-stopped: always restart unless we (the developers) forcibly stop it