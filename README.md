# TCG Games Core API

This project handles searching across data from multiple trading card games (TCGs). It is initially implemented for **MTG** and **Lorcana**, using a dependency injection and interface-based structure to make it easy to expand and support more games with different data types.

## Adding More Games

To add a new game to this project, the most important step is to add a struct that implements the `Game` interface inside the `app/games` directory. After you do that, the game is recognized by the application. For security reasons, we strictly retrieve all fields from the request and pass them to the database query, instead of passing all fields blindly. Therefore, once you add a new game, you should also add the new filters in the `CardFilters` type.

## Getting Started

To start this project, simply use the following command and wait for it to run on the port specified in your `.env` file:

```
docker compose up
```

## Postman

A file named `postman.json` at the root of this project can be imported into Postman to view details about the implemented endpoint.

## Improvements

### Database & Search

I'm using **PostgreSQL** and its JSON data type to store each game’s specific data. By applying appropriate indexing (such as GIN indexes for JSON fields and B-tree indexes for numeric fields), we can optimize partial text search, range queries, and performance for large datasets.

If more advanced full-text search or higher query performance is needed, integrating **Elasticsearch** would be a significant enhancement.

### Caching

I have **Redis** included in our Docker setup. However, due to the time constraints (this task was limited to around four hours), I did not integrate caching. In a production scenario, caching responses for a specific period could greatly improve performance under heavy load.

### Security & Extensibility

Additional improvements include sanitizing frontend data to prevent potential security issues. We could also introduce a more generalized filtering structure—supporting strings, numbers, ranges, and arrays—so that adding a new game wouldn’t require updates to the filters. This would accelerate future integrations with new TCGs.
