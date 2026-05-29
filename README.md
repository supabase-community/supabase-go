An isomorphic Go client for Supabase.

## Features

- [ ] Integration with [Supabase.Realtime](https://github.com/supabase-community/realtime-go)
  - Realtime listeners for database changes
- [x] Integration with [Postgrest](https://github.com/supabase-community/postgrest-go)
  - Access your database using a REST API generated from your schema & database functions
- [x] Integration with [Gotrue](https://github.com/supabase-community/gotrue-go)
  - User authentication, including OAuth, **_email/password_**, and native sign-in
- [x] Integration with [Supabase Storage](https://github.com/supabase-community/storage-go)
  - Store files in S3 with additional managed metadata
- [x] Integration with [Supabase Edge Functions](https://github.com/supabase-community/functions-go)
  - Run serverless functions on the edge

## Quickstart

1. To get started, create a new project in the [Supabase Admin Panel](https://app.supabase.io).
2. Grab your Supabase URL and Supabase Public Key from the Admin Panel (Settings -> API Keys).
3. Initialize the client!

_Reminder: `supabase-go` has some APIs that require the `service_key` rather than the `public_key` (for instance: the administration of users, bypassing database roles, etc.). If you are using the `service_key` **be sure it is not exposed client side.** Additionally, if you need to use both a service account and a public/user account, please do so using a separate client instance for each._

## Documentation

### Installation

Install the library to your go project:

```sh
  go get github.com/supabase-community/supabase-go
```

### Quick start

```go
  client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{})
  if err != nil {
    fmt.Println("Failed to initalize the client: ", err)
  }
```

### Client configuration

#### Basic Client

```go
client, err := supabase.NewClient(url, key, nil)
```

#### Client with custom options

```go
options := &supabase.ClientOptions{
    Headers: map[string]string{
        "X-Custom-Header": "custom-value",
    },
    Schema: "custom_schema", // defaults to "public"
}

client, err := supabase.NewClient(url, key, options)
```

### Querying data

```go
  // ...

  data, count, err := client.From("countries").Select("*", "exact", false).Execute()
```

For more see [postgrest-go Query Builder documentation](https://pkg.go.dev/github.com/supabase-community/postgrest-go#QueryBuilder)

### Authentication

The client provides comprehensive authentication features through the integrated GoTrue client.

#### Email/Password Authentication

```go
// Sign in with email and password
session, err := client.SignInWithEmailPassword("user@example.com", "password")
if err != nil {
    log.Fatal("Sign in failed:", err)
}

fmt.Printf("User ID: %s\n", session.User.ID)
fmt.Printf("Access Token: %s\n", session.AccessToken)

```

#### Phone/Password Authentication

```go
// Sign in with phone and password
session, err := client.SignInWithPhonePassword("+1234567890", "password")
if err != nil {
    log.Fatal("Sign in failed:", err)
}
```

### Token Management

#### Manual Token Refresh

```go
// Refresh an expired token
newSession, err := client.RefreshToken(session.RefreshToken)
if err != nil {
    log.Fatal("Token refresh failed:", err)
}
```

#### Automatic Token Refresh

```go
// Enable automatic token refresh in the background
client.EnableTokenAutoRefresh(session)

// The client will automatically:
// - Refresh tokens before they expire (at 75% of expiry time)
// - Retry failed refreshes with exponential backoff
// - Update all service clients with new tokens
```
