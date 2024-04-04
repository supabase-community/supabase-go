// This is basic example for postgrest-go library usage.
// For now this example is represent wanted syntax and bindings for library.
// After core development this test files will be used for CI tests.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var (
	RestUrl = `http://localhost:3000`
	headers = map[string]string{}
	schema  = "public"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	projectURL := os.Getenv("SUPABASE_URL")
	anonKey := os.Getenv("SUPABASE_ANON_KEY")
	email := os.Getenv("TESTUSER")
	password := os.Getenv("TESTUSERPASSWORD")

	client, _ := supabase.NewClient(projectURL, anonKey, nil)
	client.SignInWithEmailPassword(email, password)

	//
	rooms, _, err := client.From("rooms").Select("*", "", false).ExecuteString()
	if err != nil {
		panic(err)
	}
	fmt.Println(rooms)

}
