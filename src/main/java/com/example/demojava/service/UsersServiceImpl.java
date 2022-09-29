package com.example.demojava.service;

import com.example.demojava.model.Users;
import com.example.demojava.repository.UsersRepository;
import lombok.Getter;
import lombok.Setter;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class UsersServiceImpl implements UsersService {

    @Autowired
    @Getter
    @Setter
    private UsersRepository repository;

    @Override
    public List<Users> findAll() {
        return repository.findAll();
    }

    @Override
    public Users getOneUser(Long id) {
        return repository.findById(id).get();
    }

    @Override
    public Users createUser(Users users) {
        return repository.save(users);
    }

    @Override
    public Users updateUser(Users users) {
        return repository.save(users);
    }

    public boolean destroyUser(Long id) throws ClassNotFoundException {
            repository.deleteById(id);
            return true;
    }
}
