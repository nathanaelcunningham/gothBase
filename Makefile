# Load env from file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

app:
	air
css:
	npx tailwindcss -i ./views/css/app.css -o ./public/css/styles.css --watch
templ:
	templ generate --watch --proxy=http://localhost:${PORT} --open-browser=false
all:
	make -j3 app css templ
