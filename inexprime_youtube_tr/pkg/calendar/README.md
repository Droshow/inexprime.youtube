# This will be code relevant to retrieve of Calendar tracking
## Overview

The code setup for Google OAuth2 login involves several components:

1. **api_server.go**: This is where a new Fiber application is set up and two routes are defined: `/google_login` and `/google_callback`. These routes are handled by the `GoogleLogin` and `GoogleCallback` functions in the `controllers` package, respectively.

2. **config.go**: This file defines a `Config` struct that holds the configuration for Google OAuth2 login. The configuration is loaded from environment variables in the `LoadConfig` function.

3. **google_login.go**: This file contains the `GoogleLogin` function which generates a URL for the Google login page and redirects the user to this URL.

4. **google_callback.go**: This file contains the `GoogleCallback` function which handles the callback from Google after the user has logged in and authorized the application. It exchanges the authorization code for an access token and uses that token to fetch the user's data from Google.