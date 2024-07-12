package com.dotjson.devops.spring.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import com.dotjson.devops.spring.models.User;
import java.util.List;
import java.util.ArrayList;
import java.time.LocalDateTime;

import com.dotjson.devops.spring.repository.UserRepository;

@Service
public class UserService {
    
    @Autowired
    private UserRepository userRepository;

    public List<User> getAllUsers() {
        return userRepository.findAll();
    }

    public void saveMultipleUsers(int n) {
        List<User> users = new ArrayList<>();

        for (int i = 0; i < n; i++) {
            User user = new User();
            user.setUsername("username" + i);
            user.setEmail("user" + i + "@example.com");
            user.setPassword("password" + i); // In a real application, ensure passwords are hashed
            user.setDateOfBirth(LocalDateTime.now().minusYears(20).minusDays(i)); // Example DOB
            user.setCreationTimestamp(LocalDateTime.now());
            user.setUpdateTimestamp(LocalDateTime.now());

            users.add(user);
        }

        userRepository.saveAll(users);
    }

}
