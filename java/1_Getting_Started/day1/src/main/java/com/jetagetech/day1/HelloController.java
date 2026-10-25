package com.jetagetech.day1;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

// @RestController = ye class HTTP requests handle karegi aur JSON return karegi
@RestController
public class HelloController {

    // GET /health endpoint
    @GetMapping("/health")
    public Map<String, String> health() {
        // Map.of se banaya JSON, Spring khud convert kar dega
        return Map.of(
            "status", "ok",
            "service", "backend-from-first-principles"
        );
    }

    // GET /users/{id} - {id} URL se milega
    @GetMapping("/users/{id}")
    public Map<String, Object> getUser(@PathVariable String id) {
        // Hardcoded data (abhi database nahi hai)
        return Map.of(
            "id", id,
            "name", "Yatan",
            "email", "yatan@jetagetech.com"
        );
    }
}
