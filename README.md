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

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

Note, enum values are always validated and all unused variables are silently ignored.

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8080*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MoviesApi* | [**MoviesDeleteDelete**](docs/MoviesApi.md#moviesdeletedelete) | **Delete** /movies/delete | Delete a movie
*MoviesApi* | [**MoviesGet**](docs/MoviesApi.md#moviesget) | **Get** /movies | Get all movies
*MoviesApi* | [**MoviesPost**](docs/MoviesApi.md#moviespost) | **Post** /movies | Add a new movie
*MoviesApi* | [**MoviesSearchGet**](docs/MoviesApi.md#moviessearchget) | **Get** /movies/search | Search movies by part of title
*MoviesApi* | [**MoviesSortGet**](docs/MoviesApi.md#moviessortget) | **Get** /movies/sort | Get movies with sorting
*MoviesApi* | [**MoviesUpdatePut**](docs/MoviesApi.md#moviesupdateput) | **Put** /movies/update | Update an existing movie


## Documentation For Models

 - [Movie](docs/Movie.md)






## Author

anpuchkov

