FROM golang:1.25

WORKDIR /code

# Dependencies first so editing source code doesn't re-download the module
# graph on every rebuild.
COPY go.mod go.sum ./
RUN go mod download

# Then the source. `builds/` is excluded via .dockerignore: the blog workspace
# is purely a runtime concern (bind-mounted by compose, scaffolded on first
# start by entrypoint.sh), so `docker compose build` never needs it to exist
# and never creates it.
COPY . .
RUN chmod +x entrypoint.sh
RUN go install

EXPOSE 8000
ENTRYPOINT ["/code/entrypoint.sh"]
CMD ["ink", "preview", "builds"]
