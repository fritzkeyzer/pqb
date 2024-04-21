# PQB (Postgresql Query Builder)

Contains a set of utilities to assist in crafting postgresql queries.

The intention of this package is not to be a ORM, but a set of composable functions that
create a comfortable abstraction for slightly more complex queries.
It is aimed at go developers who prefer to write their own sql, but want
to avoid writing error-prone queries with positional arguments.
Or if you need to create dynamic queries with conditional where clauses etc.

`pqb` shines at insert/update/upsert statements and prevents you from wrangling and
manually aligning large numbers of positional arguments.

## Compatibility

`pqb` only generates sql and is intended to be used with `database/sql` or other drivers eg: `pgx`.
This package does not handle any "unmarshalling" of query results - there are plenty of good options this.
Therefore, it's completely compatible with struct scanners eg: `pgxscan`.

PQB is currently used in production systems alongside tools like [sqlc](https://docs.sqlc.dev/en/stable/index.html)
and sql drivers like `database/sql` and [pgx](https://github.com/jackc/pgx).

Sqlc tips:
- Use sqlc to generate precise structs that represent your data schema
- Use `emit_db_tags: true` to add struct tags. These are automatically detected by `pqb` 
- Use `sqlc` to generate type-safe code for simple queries and `pqb` for more complex, 
dynamic queries or inconvenient insert update queries

## Feedback or contributions are welcome!

Feel free to create a PR or issue if you have any suggestions or improvements.
This is still a WIP and I will add more features and test as I need them.

## Usage 

See the example files for self-explanatory demonstrations of how various features of this module can be used.

## Notes

The primary goal of the code is to be as clear and readable as possible.
For any reasonable query, the time spent executing the query will be orders of magnitude
larger than the time spent in the query builder.

If the query builder is not expressive enough, it's possible use only parts that are useful,
or fall back to writing the query manually.
