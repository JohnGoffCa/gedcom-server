#!/bin/bash

elm make frontend/Main.elm --output=assets/js/app.js  &&
go build -o gedcom-server ./backend/                  &&
./gedcom-server
