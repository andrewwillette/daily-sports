Program that outputs current day NBA and NFL games.

# Install
* register for [api-nba](https://rapidapi.com/api-sports/api/api-nba/) and set environment variable `RAPID_API_TOKEN` to the rapid-api-key, the nfl rest-api does not require configuration
* git clone the repository
* execute `go install`
* if `$GOBIN` is configured to your `$PATH`, the executable will now be available as `daily-sports`
* executing `go run .` at the git root directory will also execute the program
