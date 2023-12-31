# Node ------------------------------------------------------------------------
FROM node:20-slim AS ui
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
COPY . /usr/src/mayoi

WORKDIR /usr/src/mayoi/cmd/web/ui

RUN rm -rf node_modules

RUN pnpm install
RUN pnpm run build
# -----------------------------------------------------------------------------

# Go --------------------------------------------------------------------------
FROM golang:alpine AS build

COPY . /usr/src/mayoi
COPY --from=ui /usr/src/mayoi/cmd/web/ui/dist /usr/src/mayoi/cmd/web/ui/dist 

WORKDIR /usr/src/mayoi

RUN CGO_ENABLED=0 GOOS=linux go build -o mayoi cmd/web/main.go
# -----------------------------------------------------------------------------

# Bin -------------------------------------------------------------------------
FROM scratch

VOLUME /config

WORKDIR /app

COPY --from=build /usr/src/mayoi/mayoi /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV JWT_SECRET=secret

EXPOSE 6969
ENTRYPOINT [ "./mayoi" , "-d", "/config/mayoi.db", "-c", "/config/config.yml"]
