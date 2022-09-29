package com.example.demojava.service;

import com.example.demojava.model.Users;

import java.util.List;

public interface UsersService {
    public List<Users> findAll();

    public Users getOneUser(Long id);

    public Users createUser(Users users);

    public Users updateUser(Users users);

    public boolean destroyUser(Long id) throws ClassNotFoundException;
}
