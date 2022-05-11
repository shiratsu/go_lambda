FROM node:latest

WORKDIR /usr/src/app

COPY front/package*.json ./
RUN npm cache verify
RUN npm install