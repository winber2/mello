### Getting Started

Make sure you have node installed, then go into the client folder and run:

`npm install`

This will install all the node dependencies for react. If you want to run just the client server locally, run:

`npm run start`

Currently this local webpack server expects the go server to run on `localhost:8000`. If you want the go server to also host the built version of the frontend, you can build the react code using:

`npm run build`

This will put the build js bundles in the `client/public/` folder. The one caveat is that you will need to add `public/` before the `index.{some_hash}.js` src attribute in the `script` tag inside the the `index.html` file generated inside `/public`. This will likely be different once we have deploy scripts however. You can run the go server using:

`go run server.go`
