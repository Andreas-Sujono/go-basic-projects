# Go Monorepo for NodeJs Developer
based on NestJs structure

- entities: the actual database schema
- domains: abstraction of database entities that will be used accross applications
- repositories: typeorm implementation to run query to the schema
- usecases: like service in nestjs which is one level implementation on top of repositories
- handlers: like controllers in nestjs which handles request validation, response returning, and call at least 1 usecases
- routes: routes declaration, 1 route per 1 handler
- others
    - middlewares
    - guards
    - pipe

## libraries
- github.com/joho/godotenv: read from env variable