## Dockerfile

 - `FROM` : specify base image
 - `RUN` : execute command when container is built
 - `CMD`: execute command when container is start

 ## `docker build .`
Tell docker client to build image with dockerfile

### Docker build use caching
The command order should be same

### `docker build -t stephengrider/redis:lates .`
- `-t` tags the image
- `stephengrider` is docker id
- `redis` is Repo/Project name
- `latest` is version. It should be actual version number

### docker commit
- Can take a snapshot out of running container
#### example
`$ docker run -it alpine sh`
and run `# apk add --update redis`
`$ docker commit -c 'CMD ["redis-server"]' <container id>`
Then it will return new image


## Example [#](https://www.udemy.com/course/docker-and-kubernetes-the-complete-guide/learn/lecture/11436708#overview)

`FROM alpine`
  - download alpine image

`RUN apk add --update redis`
  - Get image from previous step
  - Create a container out of it
  - Run `apk add --update redis` in it
  - Take snapshot of that container's FS
  - Shut down the temporary container
  - Get image ready for next instruction
`CMD ["redis-server"]`
  - Get image from last step
  - Create a container out of it
  - Tell container it should run `redis-server` when started
  - Shut down that temporary container
  - Get image ready for next instruction

