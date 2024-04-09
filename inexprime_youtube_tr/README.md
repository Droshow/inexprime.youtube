## Project Structure
The project has the following structure:
- `cmd/inexprime_youtube_tr/main.go`: The entry point of the application.
- `pkg/youtube`: Contains the code to interact with the YouTube API.
- `pkg/calendar`: Contains the code to interact with the Google Calendar API.
- `internal`: Contains private application and library code.
- `README.md`: The project's README file.
- `.gitignore`: Specifies which files and directories to ignore in Git.
- `go.mod` and `go.sum`: Handle the project's dependencies.

# for IDE purposes
/inexprime_youtube_tr
  /cmd
    /inexprime_youtube_tr
      main.go
  /pkg
    /youtube
      (code to interact with YouTube API)
    /calendar
      (code to interact with Google Calendar API)
  /internal
    (private application and library code)
  README.md
  .gitignore
  go.mod
  go.sum

# Google Calendar Integration in Go

This project integrates with Google Calendar API to create events. It uses OAuth 2.0 for authorization.

## Setup

1. **Set up OAuth 2.0 credentials**

   Go to the Google Cloud Console, create a new project, and set up OAuth 2.0 credentials. You'll get a client ID and client secret.

2. **Configure your Go project**

   Use the client ID and client secret to configure the OAuth 2.0 flow in your Go project.

## Implementation

1. **Implement OAuth 2.0 Authorization Flow**

   Write code to handle the following:

   - Redirection to the Google consent page
   - Obtaining the authorization code
   - Exchanging the authorization code for an access token

2. **Create a Google Calendar service**

   Use the access token to create a service client that can interact with the Google Calendar API.

3. **Create calendar events**

   Define event details and make API calls to insert these events into the user's calendar.

## Usage

After setting up and implementing the project, you can create Google Calendar events by running the Go application and following the prompts.