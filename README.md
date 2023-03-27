# sqlcommenter-hertz

hertz middleware for https://github.com/google/sqlcommenter

> sqlcommenter is a suite of middlewares/plugins that enable your ORMs to augment SQL statements before execution, with comments containing information about the code that caused its execution. This helps in easily correlating slow performance with source code and giving insights into backend database performance. In short it provides some observability into the state of your client-side applications and their impact on the database’s server-side.

## Usage

```go
    ...
    h.Use(sqlcommenterhertz.SQLCommenterMiddleware())
```
