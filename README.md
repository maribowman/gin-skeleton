# gin-skeleton

Sample microservice skeleton including the [gin-gonic](https://github.com/gin-gonic/gin) HTTP web framework with a [Prometheus](https://prometheus.io/) monitoring integration, the [Postgres](https://www.postgresql.org/) open-source database using the [gorm](https://gorm.io/docs/) ORM library and the [resty](https://github.com/go-resty/resty) HTTP/REST client library. The configuration management was inspired by [Spring Boot's externalized configurations](https://docs.spring.io/spring-boot/docs/current/reference/html/features.html#features.external-config) leveraging [viper](https://github.com/spf13/viper) for the actual integration. <br>
For extended flexibility, simplified maintenance and better testability the microservice's structure is based on the clean architecture concept by Unclebob.

> The center of your application is not the database. Nor is it one or more of the frameworks you may be using. **The center of your application is the use cases of your application**  -  _Unclebob_ ([source](https://blog.8thlight.com/uncle-bob/2012/05/15/NODB.html "NODB"))

<p align="center">
  <img src="https://github.com/mattia-battiston/clean-architecture-example/blob/master/docs/images/clean-architecture-diagram-1.png" width="600">
  <img src="https://github.com/mattia-battiston/clean-architecture-example/blob/master/docs/images/clean-architecture-diagram-2.png" width="600">
</p>
