# Specify a base image
FROM node:alpine

WORKDIR /usr/app

# Install some dependency
COPY ./package.json ./
RUN npm install
COPY ./ ./

# Default command
CMD ["npm", "start"]
