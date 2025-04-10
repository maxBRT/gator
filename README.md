# Gator - RSS Feed Aggregator

A command-line RSS feed aggregator written in Go.

## Prerequisites

- Go 1.24 or later
- PostgreSQL
- A running Postgres database server

## Installation

```bash
go install github.com/maxBRT/gator@latest
```

## Configuration

Create a `.gatorconfig.json` file in your home directory:

```json
{
  "db_url": "postgres://username:password@localhost:5432/dbname",
  "current_user_name": ""
}
```

## Usage

### User Management
```bash
# Register a new user
gator register <username>

# Login as existing user
gator login <username>

# List all users
gator users

# Reset database (removes all users)
gator reset
```

### Feed Management
```bash
# Add a new feed
gator addfeed <name> <url>

# List all feeds
gator feeds

# Follow a feed
gator follow <feed_url>

# List followed feeds
gator following

# Unfollow a feed
gator unfollow <feed_url>
```

### Content
```bash
# Browse posts (default shows 2 posts)
gator browse [limit]

# Start feed aggregation (runs continuously)
gator agg <duration>
# Example: gator agg 1m  # Fetch every minute
```


