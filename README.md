# Tango-Sync

Integration between tango tienda API and tienda nube API to automatize the synchronization
of stock, images, etc.

## First steps

Tango-Sync has a PostgreSQL database so need to create a new database with tangosync name to run the project. 
The tables are created automatically. 

Also to run the project, need to set env vars that will be injected
into repositories.

- PORT
- TANGO_ACCESS_TOKEN
- TN_AUTHENTICATION
- TN_USER_AGENT
- TN_NUMBER