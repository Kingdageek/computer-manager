## Computer Manager

Simple backend service for system admins to manage computers in their organizations.

### General Features

- Computers can be fetched, created, updated, and deleted
- Computers can be assigned and unassigned to employees using employee codes (The update endpoint handles both)
- Computers assigned to multiple employees can be fetched or filtered out via the get all computers endpoint
- System admins will be notified if an employee has been assigned to 3 (this is configurable) or more computers

### To Build

- Ensure you have `docker compose` installed
- Ensure that ports 8000, 8081, and 54322 are free on your computer
- Go inside the project directory and run `cp .env.example .env`
- Run `docker compose up`
- The app should be running on `localhost:8000`
- The admin notification service should be running on `localhost:8081`
- The database should be auto-created and the `computers` table should auto-migrate

### Potential Improvements

- **Proper pagination** in the get all computers endpoint preferably cursor-based pagination for performance
- **Caching:** for the get computer by ID & get computer results, a caching layer would reduce latency after the first request
- **Better error handling:** If notifying an admin fails, the error is just logged. A potential alternative could be using queues with some retry mechanism for failed tasks
- **Improved validation:** validating the IP and MAC addresses sent are actually valid and not just random strings
- **Adequate Test suite:** both integration tests for the REST endpoints & unit tests
- **Better migration with rollbacks** to make development smoother
