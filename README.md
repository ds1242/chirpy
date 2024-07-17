# Chirpy Web Server
This is a HTTP server that is similar to a clone of a certin social media site.  Users can be created and updated.  Then a user can create, update, delete, a chirp.  Chirps can be displayed with optional params.

There is also a webhook for updating a user.

## Installation
A .env file is need to run this program.  In your .env add JWT_SECRET= and PolkaAPIKey=

## Running the program
Run the following command ```go build -o out && ./out``` an optional ```--debug``` flag can be added if you want to clear the local json database.


## Future Plans
    - Need to clean up handler funcs
    - More documentation on API endpoints
    - Add more in line documentation

