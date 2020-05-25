# MyFridge Server App

## Running the app
Requirements:
- Golang v1.14 or newer, go [here](https://golang.org/doc/install) for more info on how to install it.
- Postgresql v12 or newer, go [here](https://www.postgresql.org/download/) for more info on how to install it.

To run (locally with postgres instance):
1. First, run `go get` to install dependencies
2. Ensure your local postgres service is running, and create a database using the psql console or your IDE tools (TODO: automate this step).
3. Copy-paste the `template.env` file into the project root, and rename the copy to `.env`.
4. Then fill in the respective values in the `.env` according to the postgres db you created in step 2.
5. Next run `bash migrate.sh up` to run the necessary database migrations.
6. Then run `bash start.sh`, voila! The server should now be running if all went well.