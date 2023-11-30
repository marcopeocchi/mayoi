# Bun ----------------------------------------------------------
FROM oven/bun:1 as ui

COPY ./cmd/api/ui /usr/src/mayoi

WORKDIR /usr/src/mayoi

RUN bun install
RUN bun run build
# --------------------------------------------------------------

# Go -----------------------------------------------------------
FROM golang AS build

COPY . /usr/src/mayoi
COPY --from=ui /usr/src/mayoi/dist /usr/src/mayoi/cmd/api/ui/dist

WORKDIR /usr/src/mayoi

RUN npm install
RUN npm run build

WORKDIR /usr/src/mayoi
RUN CGO_ENABLED=0 GOOS=linux go build -o mayoi cmd/api/main.go
# --------------------------------------------------------------

# Bin ----------------------------------------------------------
FROM scratch

VOLUME /config

WORKDIR /app

COPY --from=build /usr/src/mayoi/mayoi /app

ENV JWT_SECRET=secret

EXPOSE 6969
ENTRYPOINT [ "./mayoi" , "-d", "/config/mayoi.db", "-c", "/config/config.yml"]