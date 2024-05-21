# Telegram Bot

This project is a Telegram bot developed using Golang and PostgreSQL. The bot saves all messages sent in a chat and stores messages that contain a specific filter word in a separate table.

## Setup and Usage

### Starting the Services

1. **Add your telegram bot token:**
   - Go to docker-compose.yml and on 52nd line you will see the following: "TELEGRAM_BOT_TOKEN=TELEGRAM_TOKEN"
   - Replace "TELEGRAM_TOKEN" with your token.

2. **Start all services:**
    ```sh
    docker-compose up
    ```
   This command will pull the necessary images from Docker Hub and start all services.

3. **Send commands to the bot:**
    - **/start**: The bot will respond with a welcome message.
    - **/filter + word**: The bot will save the word and filter messages by it.
    - Any other message will be saved to one of the tables, depending on whether it contains the filter word or not.

### Checking the Database

1. **Connect to the `db-psql` container:**
    ```sh
    docker exec -it telegram_bot-db-psql-1 psql -U postgres
    ```
   2. **Check available databases and connect to the `telegram_records` database:**
       ```postgresql
       \l
       \c telegram_records
       ```
   3. **Select all unfiltered and filtered messages and exit psql shell:**
       ```postgresql
       SELECT * FROM messages;
       SELECT * FROM filtered_messages;
       \q
       ```

### Stopping All Services

```sh
docker-compose down
```

## Development Decisions

1. **Project Structure:**
The project is split into three main directories:
    - **`init` Directory:** Contains the `main.go` file which connects to a database and 
   initialises `MessageHandler` and `MessagesDB` structures as well as the Telegram bot.
    - **`handlers` Directory:** Contains functions for processing incoming messages and commands.
    - **`storage` Directory:** Works with the database.

2. **Handler and Storage Functions:**
    - Handler functions are encapsulated within the `MessageHandler` structure. This allows shared variables such as 
   `messagesDB` (for database operations) and `filterWord` (for filtering messages) to be easily accessed and managed.
    - Storage functions belong to the `MessagesDB` structure, to ensure that sql database can be accessed from the functions.

3. **Docker Compose:**
    - Docker compose is used and separate container is created to make migrations to the database