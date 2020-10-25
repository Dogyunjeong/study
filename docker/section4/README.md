## Access docker node server service via web browser

### How to install what we need

- Use base image which include node
- Install by our self with `RUN`

### How to add files into image

`COPY` [path to folder to copy from on local(relative to build context)] [place to copy stuff]


### How to map inside container port

`docker run -p 8080:8080 <image name>`
docker run -p [port in local network]:[port inside container] <image name>


### How to specify working directory inside of container

`WORKDIR /usr/app`
- any following command will be executed relative to this path in the container


### **Docker build cache mechanism**

- Validate each file changes and if it is changed rebuild since that command line


## Docker networks between containers

### Options
#### Use Docker CLI's Networking Features
- Painful to re run it due to we have to set it up all that time
#### Use Docker Compose
- Docker compose is separate tool from CLI
- Used to start up multiple Docker containers at the same time
- Automates some of the long-winded arguments we were passing to 'docker run'