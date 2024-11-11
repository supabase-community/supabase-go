An isomorphic Go client for Supabase.

## Features

- [ ] Integration with [Supabase.Realtime](https://github.com/supabase-community/realtime-go)
  - Realtime listeners for database changes
- [x] Integration with [Postgrest](https://github.com/supabase-community/postgrest-go)
  - Access your database using a REST API generated from your schema & database functions
- [x] Integration with [Auth](https://github.com/supabase-community/auth-go)
  - User authentication, including OAuth, ***email/password***, and native sign-in
- [x] Integration with [Supabase Storage](https://github.com/supabase-community/storage-go)
  - Store files in S3 with additional managed metadata
- [x] Integration with [Supabase Edge Functions](https://github.com/supabase-community/functions-go)
  - Run serverless functions on the edge

## Quickstart

1. To get started, create a new project in the [Supabase Admin Panel](https://app.supabase.io).
2. Grab your Supabase URL and Supabase Public Key from the Admin Panel (Settings -> API Keys).
3. Initialize the client!

*Reminder: `supabase-go` has some APIs that require the `service_key` rather than the `public_key` (for instance: the administration of users, bypassing database roles, etc.). If you are using the `service_key` **be sure it is not exposed client side.** Additionally, if you need to use both a service account and a public/user account, please do so using a separate client instance for each.*

## Documentation

### Get Started

First of all, you need to install the library:

```sh
  go get github.com/supabase-community/supabase-go
```

Then you can use

```go
  client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{})
  if err != nil {
    fmt.Println("cannot initalize client", err)
  }
  data, count, err := client.From("countries").Select("*", "exact", false).Execute()
```

### Use authenticated client

```go

 client, err := supabase.NewClient(API_URL, API_KEY, &supabase.ClientOptions{})
 if err != nil {
  fmt.Println("cannot initalize client", err)
 }
 client.SignInWithEmailPassword(USER_EMAIL, USER_PASSWORD)

```
