package supabase_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/supabase-community/auth-go/types"

	supa "github.com/supabase-community/supabase-go"
)

const (
	API_URL       = "https://your-company.supabase.co"
	API_KEY       = "your-api-key"
	TEST_EMAIL    = "test@example.com"
	TEST_PASSWORD = "test123456"
	TEST_PHONE    = "+1234567890"
)

func TestFrom(t *testing.T) {
	client, err := supa.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("countries").Select("*", "exact", false).Execute()
	fmt.Println(string(data), err, count)
}

func TestRpc(t *testing.T) {
	client, err := supa.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	result := client.Rpc("hello_world", "", nil)
	fmt.Println(result)
}

func TestStorage(t *testing.T) {
	client, err := supa.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	result, err := client.Storage.GetBucket("bucket-id")
	fmt.Println(result, err)
}

func TestFunctions(t *testing.T) {
	client, err := supa.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	result, err := client.Functions.Invoke("hello_world", map[string]interface{}{"name": "world"})
	fmt.Println(result, err)
}

// use the new auth package to test auth functions
func TestAuth(t *testing.T) {
	client, err := supa.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Fatalf("cannot initialize client: %v", err)
	}

	t.Run("SignUp with Email", func(t *testing.T) {
		signupReq := types.SignupRequest{
			Email:    TEST_EMAIL,
			Password: TEST_PASSWORD,
		}

		resp, err := client.Auth.Signup(signupReq)
		if err != nil {
			t.Errorf("SignUp failed: %v", err)
		}
		if resp.User.Email != TEST_EMAIL {
			t.Errorf("Expected email %s, got %s", TEST_EMAIL, resp.User.Email)
		}
	})

	t.Run("SignIn with Email", func(t *testing.T) {
		tokenReq := types.TokenRequest{
			GrantType: "password",
			Email:     TEST_EMAIL,
			Password:  TEST_PASSWORD,
		}

		resp, err := client.Auth.Token(tokenReq)
		if err != nil {
			t.Errorf("SignIn failed: %v", err)
		}
		if resp.User.Email != TEST_EMAIL {
			t.Errorf("Expected email %s, got %s", TEST_EMAIL, resp.User.Email)
		}
	})

	t.Run("SignIn with Phone", func(t *testing.T) {
		tokenReq := types.TokenRequest{
			GrantType: "password",
			Phone:     TEST_PHONE,
			Password:  TEST_PASSWORD,
		}

		resp, err := client.Auth.Token(tokenReq)
		if err != nil {
			t.Errorf("Phone SignIn failed: %v", err)
		}
		if resp.User.Phone != TEST_PHONE {
			t.Errorf("Expected phone %s, got %s", TEST_PHONE, resp.User.Phone)
		}
	})

	t.Run("Get User", func(t *testing.T) {
		// First sign in to get a token
		tokenReq := types.TokenRequest{
			GrantType: "password",
			Email:     TEST_EMAIL,
			Password:  TEST_PASSWORD,
		}

		authResp, err := client.Auth.Token(tokenReq)
		if err != nil {
			t.Errorf("SignIn failed: %v", err)
		}

		// Use the token to get user info
		client.UpdateAuthSession(authResp.Session)
		user, err := client.Auth.GetUser()
		if err != nil {
			t.Errorf("Get user failed: %v", err)
		}
		if user.Email != TEST_EMAIL {
			t.Errorf("Expected email %s, got %s", TEST_EMAIL, user.Email)
		}
	})

	t.Run("Refresh Token", func(t *testing.T) {
		// First sign in to get a token
		tokenReq := types.TokenRequest{
			GrantType: "password",
			Email:     TEST_EMAIL,
			Password:  TEST_PASSWORD,
		}

		authResp, err := client.Auth.Token(tokenReq)
		if err != nil {
			t.Errorf("SignIn failed: %v", err)
		}

		// Wait a moment before refreshing
		time.Sleep(1 * time.Second)

		// Refresh the token
		refreshReq := types.TokenRequest{
			GrantType:    "refresh_token",
			RefreshToken: authResp.RefreshToken,
		}

		newResp, err := client.Auth.Token(refreshReq)
		if err != nil {
			t.Errorf("Token refresh failed: %v", err)
		}
		if newResp.AccessToken == authResp.AccessToken {
			t.Error("Expected new access token to be different")
		}
	})

	t.Run("Logout", func(t *testing.T) {
		err := client.Auth.Logout()
		if err != nil {
			t.Errorf("Logout failed: %v", err)
		}
	})
}
