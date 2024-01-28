# Specify hash to prevent hijacking attacks
FROM node:12-alpine@sha256:0e96b3d6c9e2350819e00b9148b3f28de38a3d5cc72f6cae6029aceec50ead5f

# Create app directory
WORKDIR /app

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
COPY ./mock/package*.json ./
RUN npm ci

# Bundle add app source, using the dockerignore to prevent unwanted files to be copied
COPY ./mock .

# Expose mock server port
EXPOSE 3000

# Start mock server
ENTRYPOINT npm start