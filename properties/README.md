- server:
  - host: hostName (localhost)
  - port: port (8080)

- database:
  - port: dbPort
  - user: dbUsername
  - password: dbPassword
  - database_name: authentication
  - address: dbAddress

- redis:
  - env: REDIS_DSN
  - port: redisPort (6379)
  - address: redisAddress (localhost)

- kafka:
  - producer:
    - otp:
      - bootstrap-server: kafkaProducerAddress (localhost:9092)
      - topic: otp
  - consumer:
    - otp:
      - bootstrap-server: kafkaConsumerAddress (localhost:9092)
      - topic: otp
      - group: authentication