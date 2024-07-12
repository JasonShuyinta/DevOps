package com.dotjson.devops.spring.repository;

import org.bson.types.ObjectId;
import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.stereotype.Repository;

import com.dotjson.devops.spring.models.User;

@Repository
public interface UserRepository extends MongoRepository<User, ObjectId>{
    
}
