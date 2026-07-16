# Gator 🐊

Gator is a command-line interface (CLI) application built in Go that acts as an RSS feed collector and viewer. It uses a PostgreSQL database to store feeds, subscriptions, and collected posts, allowing you to easily track and read your favorite RSS feeds directly from the terminal.

## Prerequisites

Before running Gator, ensure you have the following installed on your system:

* **Go:** [Download and install Go](https://go.dev/doc/install) (version 1.20 or higher is recommended).
* **PostgreSQL:** [Download and install PostgreSQL](https://www.postgresql.org/download/). You will need a running database instance to store the RSS data.

## Installation

You can install the `gator` CLI tool globally on your system using the `go install` command. 

If you are pulling from a remote repository:
```bash
go install [github.com/yourusername/gator@latest](https://github.com/yourusername/gator@latest)

(Note: Ensure your Go bin directory (usually ~/go/bin) is added to your system's PATH so you can run the gator command from anywhere).

If you have cloned the repository locally, navigate to the project directory and run:

Bash
go install .
Configuration
Gator requires a configuration file to connect to your PostgreSQL database.

Create a file named .gatorconfig.json in your home directory (e.g., ~/.gatorconfig.json).

Add your PostgreSQL database connection URL to the file in the following format:

JSON
{
  "db_url": "postgres://username:password@localhost:5432/gator_db?sslmode=disable"
}
Make sure to replace username, password, and gator_db with your actual PostgreSQL credentials and database name. You must create the database in PostgreSQL before running the program.

Usage
Once installed and configured, you can use the gator CLI to manage and view your RSS feeds. Below are a few of the primary commands:

Add a New Feed
To start tracking a new RSS feed, use the add command followed by the feed URL:

Bash
gator add [https://example.com/rss.xml](https://example.com/rss.xml)
List Subscriptions
To see all the feeds you are currently subscribed to:

Bash
gator list
Fetch Feeds
To manually trigger the collector to fetch the latest posts from all your subscribed feeds and store them in the database:

Bash
gator fetch
View Posts
To read the latest collected posts in your terminal:

Bash
gator view
You can run gator help at any time to see a full list of available commands and flags.
