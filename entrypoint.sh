#!/bin/sh
set -e

BUILDS=/code/builds
TEMPLATE=/code/template

# `docker compose up` bind-mounts ./builds from the host. On a bare clone the
# folder doesn't exist (it's gitignored), so Docker creates it empty right
# before we run -- seed it from the bundled template so `docker compose up`
# yields a working blog with no manual setup. Never touches an existing blog.
# `docker compose build` doesn't reach this point at all, which is why it
# leaves no ./builds behind.
if [ ! -f "$BUILDS/config.yml" ]; then
	echo "No blog found in ./builds, scaffolding one from the bundled template..."
	mkdir -p "$BUILDS"
	cp -R "$TEMPLATE/." "$BUILDS/"
	scaffolded=1
fi

# The template ships its own theme, but heal it independently in case it's
# ever missing (e.g. someone deletes just this folder) so the blog never fails
# to build for a missing theme.
if [ ! -f "$BUILDS/theme/config.yml" ]; then
	echo "No theme found in ./builds/theme, restoring it from the bundled template..."
	rm -rf "$BUILDS/theme"
	cp -R "$TEMPLATE/theme" "$BUILDS/theme"
	scaffolded=1
fi

# Anything scaffolded above was written as root but is bind-mounted back to
# the host: hand it to the host user (override with PUID/PGID) so it stays
# editable without sudo.
if [ -n "$scaffolded" ] && [ "${PUID:-1000}" != "0" ]; then
	chown -R "${PUID:-1000}:${PGID:-1000}" "$BUILDS"
fi

exec "$@"
