# Go API client

API for managing movies

## Overview

- API version: 1.0.0
- Package version: 1.0.0

## Installation

Install the following dependencies:

```shell
go mod download
go mod tidy
```

Put the package under your project folder and add the following in import

## Authorization for users

Simple authorization and register with 2 spaces. For authorize, use 
/login page. For register use /register page.
```json
{
  "username": "your-username",
  "password": "your-password"
}
```


## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

Note, enum values are always validated and all unused variables are silently ignored.

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080*

| Class       | Method                                                   | HTTP request              | Description                    |
|-------------|----------------------------------------------------------|---------------------------|--------------------------------|
| *MoviesApi* | [**MoviesDelete**](docs/MoviesApi.md#MoviesDelete)       | **Delete** /movies/delete | Delete a movie                 |
| *MoviesApi* | [**MoviesGet**](docs/MoviesApi.md#MoviesGet)             | **Get** /movies           | Get all movies                 |
| *MoviesApi* | [**MoviesPost**](docs/MoviesApi.md#MoviesPost)           | **Post** /movies          | Add a new movie                |
| *MoviesApi* | [**MoviesSearchGet**](docs/MoviesApi.md#MoviesSearchGet) | **Get** /movies/search    | Search movies by part of title |
| *MoviesApi* | [**MoviesSortGet**](docs/MoviesApi.md#MoviesSortGet)     | **Get** /movies/sort      | Get movies with sorting        |
| *MoviesApi* | [**MoviesUpdatePut**](docs/MoviesApi.md#MoviesUpdatePut) | **Put** /movies/update    | Update an existing movie       |

## Documentation For Models







## Author

anpuchkov

