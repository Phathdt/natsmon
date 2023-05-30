FROM node:18.16-alpine as fe-builder
WORKDIR /app
COPY frontend/package.json /app/package.json
COPY frontend/yarn.lock /app/yarn.lock
RUN yarn install
COPY frontend /app
RUN yarn build

FROM golang:1.20.4-alpine as builder
WORKDIR /app
COPY backend/go.* /app/
RUN go mod download
COPY backend /app
RUN go build -ldflags '-w -s' -a -o natsmon main.go

FROM alpine:3.17
WORKDIR /app
RUN chown nobody:nobody /app
USER nobody:nobody
COPY --from=builder --chown=nobody:nobody ./app/natsmon .
COPY --from=fe-builder --chown=nobody:nobody ./app/apps/web/dist /app/public
COPY --from=builder --chown=nobody:nobody ./app/run.sh .

ENTRYPOINT sh run.sh
