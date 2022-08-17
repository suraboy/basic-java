package com.example.basicjava

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.data.annotation.Id
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@SpringBootApplication
class BasicJavaApplication

fun main(args: Array<String>) {
	runApplication<BasicJavaApplication>(*args)
}

@RestController
class MessageResource {
	@GetMapping
	fun index(): List<Message> = listOf(
		Message("1", "Hello!"),
		Message("2", "Sirichai!"),
		Message("3", "Boy!"),
	)
}

data class Message(val id: String?, val text: String)