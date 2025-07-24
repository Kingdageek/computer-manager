## Computer Manager

Simple golang backend service for system admins to manage computers in their organizations.

### General Features

- Computers can be fetched, created, updated, and deleted
- Computers can be assigned and unassigned to employees using employee codes (The update endpoint handles both)
- Computers assigned to multiple employees can be fetched or filtered out via the get all computers endpoint
- System admins will be notified if an employee has been assigned to 3 (this is configurable) or more computers

### To Build & Run

- Ensure you have `docker compose` installed
- Ensure that ports 8000, 8081, and 54322 are free on your computer
- Go inside the project directory and run `cp .env.example .env`
- Run `docker compose up`
- The app should be running on `localhost:8000`
- The admin notification service should be running on `localhost:8081`
- Postgres should be running on `localhost:54322`
- Try the APIs on swagger by visiting: `localhost:8000/swagger/index.html`
- The database should be auto-created and the `computers` table should auto-migrate

If you get any error during build related to the `computer-manager/docs` package not being part of `$GOROOT`, go to the [routes.go](./internal/api/routes/routes.go) file and comment out this line in the import: `_ "computer-manager/docs"`. Re-run `docker compose up` then uncomment it to see the swagger page.

### Potential Improvements

- **Authentication & Authorization** to verify the identity of the sysadmin and to ensure they have the proper permission to perform any action
- **Proper pagination** in the get all computers endpoint preferably cursor-based pagination for performance
- **Caching:** for the get computer by ID & get computer results, a caching layer would reduce latency after the first request
- **Better error handling:** If notifying an admin fails, the error is just logged. A potential alternative could be using queues with some retry mechanism for failed tasks
- **Improved validation:** validating the IP and MAC addresses sent are actually valid and not just random strings
- **Adequate Test suite:** both integration tests for the REST endpoints & unit tests
- **Better migration with rollbacks** to make development smoother

### Other Notes

- The get all computers endpoint expects employee codes to be added as separate query params like so `employee_codes=mmu&employee_codes=ckm` but the swagger
  definition uses a comma-separated list like so `employee_codes=mmu,ckm`, this doesn't work.
