# Timezones calculator Microservice

This is a simple CRUD program built with Golang's `Mux` and a `Mongo` database from a Docker instance. The program will interact and manage basic operations on timezones records.

## Provisioning

For the persistence layer, a Mongo 3.6 Docker image is used; all the data is found inside the `timezones.json`.

To perform a database connection and insert records, you should do the following:

### Docker Setup

1. Pull the image: `docker pull mongo:3.6`
2. Create a container: `docker run --name [your_instance_name] -d -p 27017:27017 mongo:3.6`
3. Accessing the container: `docker exec -it [your_instance_name] /bin/bash`

### Inserting the records

1. Run: `go run main.go`

### Interacting with the database

After the third step in the `Docker Setup` section.

1. Run: `mongo`

Inside the shell

1. `show dbs`
  - At this point, it should appear the `timezones` doc database

2. `use timezones`
3. `show collections`
  - It should appear in the `timezones` collection

4. `tz = db.timezones`
5. `tz.find()`
  - A bunch of records will appear

6. Alternatively, you can run `tz.count()`


## Running script

There are two scripts.

1. The `first.go` script is a simple PoC code that demonstrates the handler functions and performs static conversions of timezones based on an in-memory data structure.

2. However, the complete API is inside the `/api` directory. You will need to install the dependencies manually and set up the correct path for each one.

## Usage

Execute: `go run api/main.go` and follow the REST pattern for the `timezones` resource. You can check the file directly for more information.

## Credits
[David Lares S](https://davidlares.com)

## License
[MIT](https://opensource.org/licenses/MIT)
