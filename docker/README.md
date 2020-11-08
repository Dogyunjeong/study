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

## Development work flow
**development** -> **testing** -> **deployment** -> **development** ...


### Use local files as a container volume

#### using docker-compose configuration

attaching volumes.

## Docker Build

- keyword: `base image`

### Multi build process (Multi step docker file)

can build react project and use buildt source code to serve with nginx.
So there will be two step. build react project in first docker build and build nginx container with the built static file to serve


## Continuous Integration with travis CI and aws

Once there is new pull request on master, travis ci will do jobs and dploy in AWS

### Travis is CI tool

Deploying is only happen when master branch is updated
Test is happens when new pull request is created onto master.

1. Connect travis and github repository
2. setup `.trvis.yml` file
    1. Set up environment
    2. Add script to test
3. Setup deploying
    1. create aws elasticbeanstalk
        - About the platform: `docker running on 64bit amazon linux` use `DockerFile` to build image and `linux2` use docker-compose. **be aware about this**
    2. create IAM for travis ci in aws
    3. added `ACCESS_KEY` and `SECRET_KEY` into travis ci environment variable
    4. configure `.travis.yml` for `deploy`
