# Stage 1: Build the application (uses Gradle)
FROM gradle:7.6.0-jdk17 AS build

COPY . /app
WORKDIR /app
RUN gradle bootJar

# Stage 2: Slim runtime image (only necessary files)
FROM openjdk:17-slim
EXPOSE 8080
WORKDIR /app
COPY --from=build /app/build/libs/spring-0.0.1-SNAPSHOT.jar app.jar
ENTRYPOINT ["java", "-jar", "app.jar"]
