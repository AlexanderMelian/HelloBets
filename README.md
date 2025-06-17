Hello bets is an alternative to 1xbet and similar platforms

This is my current database diagram

```mermaid
erDiagram
    User {
        long id
        string username
        string password
        string email
        decimal money  
        boolean enabled
        datetime createdAt
        datetime updatedAt
    }

    Bet {
        long id
        long userId
        long quotaId
        decimal quantity
        string status
        datetime createdAt
    }

    Quota {
        long id
        decimal percentage
        long marketId
        datetime createdAt
    }

    Event {
        long id
        string name
        date eventDate
        string status
        boolean isOpen
        datetime createdAt
        datetime updatedAt
    }

    Market {
        long id
        string type
        long eventId
        datetime createdAt
        datetime updatedAt
    }  

    Transaction {
        long id
        long userId
        decimal quantity
        int type
        int status
        date date
        datetime createdAt
    }

    User ||--o{ Bet : makes
    Event ||--o{ Market : has
    Market ||--o{ Quota : defines
    Quota ||--o{ Bet : associated_with
    User ||--o{ Transaction : performs

```

The first step is to build a functional monolith, conduct performance tests and scale.

## How to Run
To get started with the project, use the following commands:
```sh
make build   # Build the application
make up      # Start the application
```

## Types of transactions
- [ ] Deposit
- [ ] Transfer For Accounts
- [ ] Withdraw
- [ ] Bet
- [ ] Cancel Bet


## TODO
- [ ] Implement all database tables (`ALL_TABLES`)
- [ ] Add rate limiting to API endpoints (`RATE_LIMITER`)
- [ ] Enforce password attempt limits (`PASSWORD_TRY_LIMIT`)
- [ ] Add controller-level input validations (`CONTROLLER_VALIDATIONS`)
- [ ] Validate `iss` and `aud` claims in JWTs (`ISS_AUD_VALIDATION_TO_JWT`)
- [ ] Standardize and improve generic error responses (`GENERIC_ERRORS_RETURN`)
