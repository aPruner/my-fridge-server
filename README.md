# MyFridge Server App

## Running the app
Requirements:
- Golang v1.14 or newer, go [here](https://golang.org/doc/install) for more info on how to install it.
- Postgresql v12 or newer, go [here](https://www.postgresql.org/download/) for more info on how to install it.

### To run locally with postgres instance (Recommended for local development):
1. Copy-paste the `template.env` file into the project root, and rename the copy to `.env`. If you'd like, feel free to change the values in .env after doing copying the template.
2. Ensure your local postgres service is running, then run `bash setup.sh`. This will create your local postgres instance.
3. Run `bash migrate.sh up` to run the necessary database migrations.
4. Run `bash start.sh`. Voila! The server should now be running assuming all went well.

### To run locally with docker and docker-compose (Not recommended for local development):
1. Ensure you have `docker` and `docker-compose` installed! Find how to install `docker` [here](https://docs.docker.com/get-docker/), and `docker-compose` [here](https://docs.docker.com/compose/install/).
2. In an ideal world, it should be as simple as running `docker-compose up`. The app should now be running at `localhost:3000`.