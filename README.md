# golang-url-shortener
URL shortener application using Golang. 

Currently: allows creating a custom, non-persistent alias to a URL. Accessing the alias redirects a user to the desired URL.

Currently overwrites existing entries, does not check for uniqueness.

### Creating an alias:

Send a POST request to `http://127.0.0.1:8080/addUrl`

Query parameters:
- alias: the desired short URL for an application
- original: the original URL to access using the alias

### Accessing an alias:

Send a GET request to `http://127.0.0.1:8080/<alias>`

### Next steps:

1. Adding constraints to aliases, ensure aliases are unique and cannot be overwritten unless explicitly mentioned
2. Adding persistence to aliases, adding a database
3. Enabling deleting or updating aliases using dedicated endpoints
4. (Maybe) adding users and ensuring specific aliases for users only

