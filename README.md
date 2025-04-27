# Personal Site

[13 April 2025]

Creating a simple personal website.

The goal, right now, is to create a backend with Go and a frontend with React.

[14 April 2025]

Based on the guide I'm using to build this site, Let's Go (Alex Edwards), I will not be using React, but instead HTML templating.

## Things to Learn
- HTML templates
  - base template
  - the '.', dynamic data you want to pass and display in a template
  - Setting up proper configurations for a web app via the CLI
  - Dependency Injection

- Model-View-Controller Pattern
  - *Models* are the data structures and business logic 
  - *Views* are HTML (dynamic) templates that display info
  - *Controllers* are handlers that process user requests, selecting the correct

### Dependency Injection

### Loggers
At a high level, there are three core components associated with structured logging: the Logger struct, the Record struct, and the Handler interface.

The Logger is a public interface that users interact with to create logs. It provides methods such as Debug, Info, Error, and Fatal for different severity levels. These methods are the main entry points for an application to generate log entries.

When a logging method is called, the Logger creates a Record (sometimes called LogEntry); this contains information for the log event:

- Timestamp when the method was called
- Message text
- Severity level
- Structured fields in the form of key-value pairs (optional).

The structure fields, while optional, allow developers to include additional data to the logs, including users, measured values (such as system resources used and time taken), status information, and more.

Once the Record is created, it's passed to one or more Handlers. A Handler determines what happens to the log information -- where it's stored, how it's formatted, and which logs are processed based on filtering rules. A Logger can have multiple Handlers, each processing the Record differently depending on the context. For example, a development environment might use a ConsoleHandler to displays barebones data in the terminal, while a FileHandler writes comprehensive JSON-formatted logs to a file for later analysis.

### Logging -- An In-Depth Examination

[27 April 2025]

I've decided to explore logging in more detail. While a simple logger would be perfectly acceptable for a small project like mine -- a one- or two-page personal website -- I'm interested in learning more about logging leading practices. So I'll be exploring Go's `slog` package and building a slightly more sophisticated custom logger.

#### Types
There are three main components to a Logger, which I've mentioned above but will repeat here:

- **Logger**: (struct) the public interface that allows users to implement logging
- **Record**: (struct) the specific instance of a log, with details. 
- **Handler**: (interface) determines how the logs are formatted and where they're displayed or saved

__References__
- https://betterstack.com/community/guides/logging/logging-in-go/



