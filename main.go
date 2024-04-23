package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

// Create functions to print colored text.
var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

// A simple program that connects to a redis server and pings it.
func main() {
	// Print a greeting message.
	printWithPadding("Connecting to the redis server...")

	// Create a new redis client.
	client := createRedisClient()

	// Close the connection to the redis server when the main function finishes.
	defer client.Close()

	// Ping the redis server to check if it is working.
	pong, err := client.Ping().Result()
	if err != nil {
		message := red("ERROR: " + err.Error()) // Construct an error message.
		printWithPadding(message)               // Print an error message, if it exists.
	} else {
		// Print the response from the redis server.
		message := "PING: " + green(pong) // Construct the message.
		printWithPadding(message)         // Should print "PING: PONG".
	}
}

// Creates a new redis client.
func createRedisClient() (client *redis.Client) {
	// Get environment variables.
	host, foundHost := os.LookupEnv("REDIS_HOST")

	// Print if the REDIS_HOST environment variable is found.
	fmt.Println("Found host: " + colorBoolean(foundHost))

	// If the REDIS_HOST environment variable is not set, use "localhost" as the default value.
	if !foundHost {
		host = "localhost"
	}

	// Print the host value.
	fmt.Println("Host: " + host)

	// Create a new redis client.
	client = redis.NewClient(&redis.Options{
		//  Host "redis" is the name of the redis service in the docker-compose file.
		//  Host "localhost" is used when running the program outside of a container.
		Addr:     host + ":6379", // Port 6379 points to the redis server address.
		Password: "",
		DB:       0,
	})

	return
}

// Print a message with newlines above and below.
func printWithPadding(message string) {
	fmt.Printf("\n%s\n\n", message)
}

// Print a boolean value in green if true and red if false.
func colorBoolean(in bool) string {
	// Convert the boolean value to a string.
	string := strconv.FormatBool(in)

	if in {
		return green(string)
	} else {
		return red(string)
	}
}
