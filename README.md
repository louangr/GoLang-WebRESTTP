# ProjectRESTAPI (Louan)

- To build the project, run `./commands/build.cmd` and find the executable in `./cmd/restserver`
- To start the REST API, run `./commands/start.cmd`
  - The port number can be change in the first line of the `start.cmd` file
  - The .db file name can be change in the second line of the `start.cmd` file
- To generate swagger documentation, run `./commands/swaggerDocGeneration.cmd`
- Once the REST API is started, you can find the Swagger UI documentation behind this endpoint: [http://localhost:8000/swaggerui/](http://localhost:8000/swaggerui/)