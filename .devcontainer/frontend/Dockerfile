FROM node:16-alpine as build-stage

RUN apk add --update &&  apk add git && apk add bash

# RUN mkdir -p /frontend/app

WORKDIR /app

RUN npm install -g @vue/cli

# RUN npm install

# RUN npm create vite@latest