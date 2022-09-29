package com.example.demojava.controller;

import com.example.demojava.model.Users;
import com.example.demojava.service.UsersService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.crossstore.ChangeSetPersister;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.server.ResponseStatusException;

import java.util.List;

@RestController
@RequestMapping("/api")
public class UsersController {

    @Autowired
    private UsersService usersService;

    @GetMapping("/users")
    public ResponseEntity<List<Users>> getAllUsers() {
        List<Users> response = usersService.findAll();
        return new ResponseEntity<>(response, HttpStatus.OK);
    }

    @GetMapping("/users/{id}")
    public Users getProductById(@PathVariable Long id) {
        try {
            return usersService.getOneUser(id);
        } catch (Exception e) {
            throw new ResponseStatusException(
                    HttpStatus.NOT_FOUND, "Not Found", e);
        }
    }

    @PostMapping("/users")
    public Users createProduct(@RequestBody Users users) {
        return usersService.createUser(users);
    }

    @PutMapping("/users/{id}")
    public Users updateProduct(@PathVariable("id") Long id, @RequestBody Users users) {
        try {
            users.setId(id);
            return usersService.updateUser(users);
        } catch (Exception e) {
            throw new ResponseStatusException(
                    HttpStatus.NOT_FOUND, "Not Found", e);
        }
    }

    @DeleteMapping("/users/{id}")
    public void destroyUser(@PathVariable Long id) {
        try {
            usersService.destroyUser(id);
        } catch (Exception e) {
            throw new ResponseStatusException(
                    HttpStatus.NOT_FOUND, "Not Found", e);
        }
    }
}
