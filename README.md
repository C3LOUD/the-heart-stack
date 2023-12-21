# The HEART Stack

**A High-Performance, Enterprise-Ready Stack with an Emphasis on Developer Productivity**

## Technologies

[Htmx](https://htmx.org/), [Echo](https://echo.labstack.com/), [Air](https://github.com/cosmtrek/air), [Redis](https://redis.io/), [Templ](https://templ.guide/)

Also leveraging:
- [Tailwindcss](https://tailwindcss.com/)
- [Goose](https://pressly.github.io/goose/)
- [Sqlite](https://www.sqlite.org/index.html)
- [Sqlc](https://docs.sqlc.dev/en/stable/index.html#)
- [Sqlfluff](https://docs.sqlfluff.com/en/stable/index.html)

---

## Motivation

In contemporary frontend development, managing data synchronization between the backend and frontend has become overly complex. The reliance on technologies like Redux to maintain in-memory state duplication of the database is cumbersome. The frontend state often lags behind the actual database state, requiring additional mechanisms to simulate real-time updates. The HEART Stack addresses these issues by emphasizing a return to a Single Source of Truth (SSOT).

Another pain point is the redundant validation of data on the client side, even when the server has already performed permission validation. The HEART Stack encourages developers to streamline this process by responding only with data that the requesting client has the privilege to access. This approach eliminates the need to duplicate validation efforts on both the client and server sides.

Most importantly, Javascript suck, don't write it. 

---

## Contributions Welcome

We welcome contributions from the community to enhance and evolve the HEART Stack. Join us in building a robust, efficient, and developer-friendly stack for enterprise applications. Feel free to contribute by submitting bug reports, feature requests, or code improvements. Your input is vital to the continuous improvement of the HEART Stack.