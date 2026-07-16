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
```

(Note: Ensure your Go bin directory (usually ~/go/bin) is added to your system's PATH so you can run the gator command from anywhere).
If you have cloned the repository locally, navigate to the project directory and run:


```Bash
go install .
```

## Configuration
Gator requires a configuration file to connect to your PostgreSQL database.
Create a file named .gatorconfig.json in your home directory (e.g., ~/.gatorconfig.json).
Add your PostgreSQL database connection URL to the file in the following format:


```JSON
{
  "db_url": "postgres://username:password@localhost:5432/gator_db?sslmode=disable"
}
```

Make sure to replace username, password, and gator_db with your actual PostgreSQL credentials and database name. You must create the database in PostgreSQL before running the program.

## Usage

Gator is a multi-user system, meaning you must register and log in to manage your specific feed subscriptions. Below are the available commands:
User Management
Register a new user:
```Bash
gator register "<username>"
```

Log in as an existing user:
(This sets your active user in the configuration file)
```Bash
gator login "<username>"
```

List all users:
```Bash
gator users
```

Reset the database:
(Warning: This clears all database records. Useful for development/testing)
```Bash
gator reset
```

Feed Management
Note: You must be logged in to add, follow, or unfollow feeds.
Add a new RSS feed:
(This adds the feed to the system and automatically subscribes you to it)
```Bash
gator addfeed "<name>" "<url>"
```

List all available feeds:
(Shows all feeds added to the system by any user)
```Bash
gator feeds
```

Follow an existing feed:
```Bash
gator follow "<url>"
```

View your followed feeds:
```Bash
gator following
```

Unfollow a feed:
```Bash
gator unfollow "<url>"
```

## Aggregation & Reading

Run the feed aggregator:
(This kicks off the worker that reaches out to the web and fetches new posts for all feeds in the database)
```Bash
gator agg <time_between_requests>
```

Example: gator agg 1m (runs every 1 minute)
Browse your feed posts:
(Requires login. Shows the most recently fetched posts from feeds you follow)
```Bash
gator browse "<limit>"
```

Example: gator browse 10 (shows the 10 most recent posts)
