# README

## About

MovieWatcher is an application written in Vue and Go to be able to look up movies and shows and then watch them in the App.

## Env

Add a file named .env to the root of the project, this file should include `TMDB_API=your_api_key_here` to be able to access the TMDB api for searching.

## Live Development

This project uses Node version 24.14.0, run `nvm use` to apply the node version specified in the nvmrc-file.

You also need wails to be able to run this program, which comes with a couple dependencies.
Run `go install github.com/wailsapp/wails/v2/cmd/wails@latest` to install the Wails CLI through go.
after this is done you can run the `wails doctor` command to see if you have all the neccessary dependencies.

You can read more in the wails installation guide here: https://wails.io/docs/gettingstarted/installation

To run in live development mode, run `wails dev` in the project directory. 
### If you are using a never version of Linux you might need to use webkit2_41 and must then specify this in your command using `-tags`
example: `wails dev -tags webkit2_41`

This will run a Vite development server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
### If you are using a never version of Linux you might need to use webkit2_41 and must then specify this in your command using `-tags`
example: `wails build -tags webkit2_41`

You can also specify which platform you're building for using the `--platform` flag
example: `wails build -tags webkit2_41 --platform windows`
