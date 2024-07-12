package com.dotjson.devops.spring;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

import com.dotjson.devops.spring.service.UserService;

import io.github.cdimascio.dotenv.Dotenv;

import com.dotjson.devops.spring.models.User;
import java.util.List;

import lombok.extern.slf4j.Slf4j;

@SpringBootApplication
@RestController
@Slf4j
public class Application {

	@Autowired
	private UserService userService;

	public static void main(String[] args) {
		Dotenv dotenv = Dotenv.configure().load();

		// Set environment variables
		System.setProperty("MONGODB_URI", dotenv.get("MONGODB_URI"));
		SpringApplication.run(Application.class, args);
	}

	@GetMapping("/spring")
	public String helloWorld() {
		log.info("Calling helloWorld from Spring");
		return "Hello world!";
	}

	@GetMapping("/user")
	public List<User> getUsers() {
		return userService.getAllUsers(); 
	}

	@PostMapping("/user/{n}")
	public String saveMultipleUsers(@PathVariable int n) {
		userService.saveMultipleUsers(n);
		return n+" users created";
	}
}
