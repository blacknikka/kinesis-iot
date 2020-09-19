FROM node:12-alpine3.11
WORKDIR /app

ENV LANG=C.UTF-8
ENV TZ=Asia/Tokyo

COPY ./src /app
RUN yarn install && yarn build
