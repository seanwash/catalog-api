# Catalog API

## Brief

> The API should be able to process all CRUD style requests and store its data
in a relational database. The API should allow for the manipulation of Artists,
Albums, Songs and Genres at a minimum. You may choose to have read-only data,
but at least 1 field for each object should be editable. Feel free to add any
features that you think will make your API better to consume.

## Setup

1. Clone repository.
1. Ensure that Docker is installed & running.
1. Obtain `.env` file. For development, an `.env.sample` file has been provided with values suitable for this demo.
1. Run `$ docker-compose up -d` from the project root. This will build and start both the Web and Postgres services, keeping them running in the background. If you'd like to see logs, you can skip `-d`, but you'll have to run the following commands in another terminal.
1. Run `$ docker-compose run --rm web make db.up`. This will run the initial migrations. Note that `--rm` will remove the container that was created to run the command. This is fine since the container we want to keep is already running the web instance.
1. Run `$ docker-compose run --rm web go run cmd/seeds/main.go`. This will pre-seed the database with some albums from one of my favorite artists, [Plini](https://www.youtube.com/watch?v=oNFXnFAKqAE).
1. Import the Insomnia export (attached via email) and fire away!

## Using the API

The attached Insomnia workspace will include a number of sample queries and mutations for you to try out. Don't forget to check out the documentation pane!

One thing to note is that all mutations are protected and can be accessed by sending an `API-KEY` header along with the request. The Insomnia workspace should have this pre configured, so give the mutations a try with and without the header.

Additionally, a live API is available at https://catalog-api.onrender.com. The exported Insomnia workspace contains a production environment as well as a development environment.

## Database Design

It seems like these days you don't see any albums being released without multiple featured artists, and often times an album will cover the span of multiple genres. Because of this, tracks are central to the heart of this database design. 

### Notable features:
- Tracks can have multiple artists/features through `track_artists`.
- Tracks belong to multiple albums through `track_albums`, facilitating Spotify style album singles, compilation albums, as well as a traditional album release.
- Tracks can have multiple genres through `track_genres`, ensuring that searching and cataloging tracks by genre do not restrict them to the genre of their album.

### Additional features:
- When a track or its relations are removed, the joins `track_{artists,albums,genres}` are automatically removed ensuring that stale data isn't left around.
- Both genres and albums require a unique name via a unique index of their lower cased names.

### Trade offs
One trade off of this design is that most queries will involve at least one join which over time and with growth will impact performance. If the app grows and performance does become an issue, one could:
1. Throw a cache in front of the database.
2. Utilize document store tool like Elasticsearch to create documents of the tracks and their relationships so that joins are avoided altogether for most read/search queries. A drawback of this approach is that this additional data source would need to be kept up to date with the database.

## Go Libraries and ORM

The most important dependencies are:

### Sqlboiler (ORM)

[Sqlboiler](https://github.com/volatiletech/sqlboiler) utilizes introspection and code generation to build a type safe ORM tailored to your relational database. I chose to use Sqlboiler because:
1. It uses a schema first approach. The goal of this test wasn't to see whether I could implement a good ORM myself, but to see if I could design a flexible schema around tracks, albums, artists, and genres.
2. It generates a type safe API, and you're able to see all of the generated code right in the project. Because of this, there aren't any levels of indirection that can cause confusion and the auto completion is superb!

### Gqlgen + Chi (Web Interface)

[Gqlgen](https://gqlgen.com/) also utilizes code generation to build a type safe GraphQL interface for your app. Once configured, on generate it will analyze a `schema.graphql` file and generate the appropriate types and queries, as well as stub out any required resolvers required to facilitate the schema. Ultimately I chose to use Gqlgen because:
1. It seemed relatively stable and mature with an active community/maintainer.
2. It uses a declarative and code generation approach. You describe what you want and only have to implement the pieces to make that happen.

[Chi](https://github.com/go-chi/chi) is a lightweight router that works very well in tandem with Go's own `http` package. Ultimately I chose to use Chi because provides some useful middleware out of the box while still using the standard HTTP handlers. 

### Goose

[Goose](https://github.com/pressly/goose) powers the database migration system. I've been spoiled in the past by frameworks that have database migrations built in (Rails, Elixir's Ecto), so I went looking for a lib to handle this for me instead of building one from scratch. I picked Goose because I had read a few favorable reviews of it online.

## Project Structure

- `api` - Contains middleware related to the Chi router.
- `cmd` - Contains the `main` package and entry points for the HTTP server and database seeds.
- `db` - Contains the database migrations and `db.Connection` which is used to create a new connection to Postgres.
- `graph` - Contains the `schema.graphql` file as well as all of the resolvers and code generated by Gqlgen.
- `models` - Contains the models and code generated by Sqlboiler.
- `modex` - As recommended by Sqlboiler, contains a few helper functions for models. This is a separate package/directory so that when the models are generated, the custom helpers aren't removed.

## Things to Improve Upon

### Dataloader

It's really easy to introduce a ton of N+1 queries into a GraphQL service without a tool like Dataloader. In a nutshell, Dataloader is tool that keeps track of which resources have been requested and instead of making a query for each time that resource was requested, it batches those requests into one query. The batched query is similar to `select * from thing where id in [resourceIdsFromRequsts]`. While I've implemented this pattern in other languages, I'm still pretty new to Go and the recommended package [dataloaden](https://github.com/vektah/dataloaden) is a little too opaque for me at the moment. I would love to ship  a product without N+1s, but seeing as this is a small and also a low traffic API we should be ok! 

### Tests

I would love to have included a whole slew of tests for this API. No excuses, but I just couldn't get there in time.

### Idiomatic Go

Since I'm still new to Go, I'm certain that much of the code I did write could be written in a more terse or organized way. Because of this, I tried not to be fancy or utilize any concepts that I didn't feel that I had a firm grasp on. My hope is that the code base is simple and organized enough that a seasoned Go programmer can still find their way around and make edits and improvements should they have to.
