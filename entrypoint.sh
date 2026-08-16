#!/bin/sh
set -e

# On a bare clone (or if ./builds was deleted) there is no blog to serve yet.
# Bootstrap one from the bundled template so `docker compose up` works with
# no manual setup. Never touches an existing blog.
if [ ! -f /code/builds/config.yml ]; then
	echo "No blog found in ./builds, scaffolding one from the bundled template..."
	mkdir -p /code/builds
	cp -r /code/template/. /code/builds/
	# Written as root, but bind-mounted back to the host: hand it to the
	# host user (override with PUID/PGID) so it's editable without sudo.
	chown -R "${PUID:-1000}:${PGID:-1000}" /code/builds
fi

# The theme is vendored into this repo, but heal it independently in case
# it's ever missing (e.g. someone deletes just this folder) so the blog
# never fails to build for a missing theme.
if [ ! -f /code/builds/ink-theme-dark/dark/config.yml ]; then
	echo "Theme ink-theme-dark not found, cloning it..."
	rm -rf /code/builds/ink-theme-dark
	git clone --depth 1 https://github.com/InkProject/ink-theme-dark.git /code/builds/ink-theme-dark
	rm -rf /code/builds/ink-theme-dark/.git
	chown -R "${PUID:-1000}:${PGID:-1000}" /code/builds/ink-theme-dark
fi

exec "$@"
