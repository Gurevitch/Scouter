📦 project-root
│── 📂 cmd
│   └── main.go                 # Application entry point
│
│── 📂 pkg                      # Core application logic (modular & reusable)
│   ├── 📂 db                   # Database connection
│   │   └── db.go
│   │
│   ├── 📂 repository           # Data access layer (DB queries)
│   │   └── user_repository.go
│   │
│   ├── 📂 service              # Business logic layer
│   │   └── user_service.go
│   │
│   ├── 📂 handler              # HTTP request handlers
│   │   └── user_handler.go
│   │
│   ├── 📂 router               # API routing
│   │   └── router.go
│   │
│   ├── 📂 models               # Data models (structs)
│   │   └── user.go
│   │
│   ├── 📂 config               # Configuration files (env, JSON, YAML)
│   │   └── config.go
│   │
│   ├── 📂 utils                # Utility functions (helpers)
│   │   └── hash.go             # Password hashing
│   │
│── 📂 internal                 # Internal app logic (if needed)
│
│── 📂 scripts                  # DevOps & database migration scripts
│
│── 📂 docs                     # Documentation & API specs
│
│── 📂 test                     # Unit & integration tests
│   └── user_service_test.go
│
│── go.mod                      # Go module file
│── go.sum                      # Go dependencies



1. get full team and save in DB
2.create simple front page and display
3.create a code block and understand of the design pattern
4.find a way to save the entire league


https://www.capology.com/club/tottenham/salaries/