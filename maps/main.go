package main

import "fmt"

// main is the entry point of the program.
func main() {
    // Print a header for the maps section.
    fmt.Println("Maps")

    // Create a map to store programming languages with their abbreviations.
    languages := make(map[string]string)
    
    // Add some key-value pairs to the map.
    languages["JS"] = "JavaScript"
    languages["PH"] = "PHP"
    languages["PY"] = "Python"
    
    // Print the entire map.
    fmt.Println("Languages", languages)
    
    // Retrieve and print the value associated with the key "JS".
    fmt.Println("Languages", languages["JS"])
    
    // Delete the key-value pair with the key "PH" from the map.
    delete(languages, "PH")
    
    // Print the map after deletion.
    fmt.Println("Languages after delete", languages)
    
    // Iterate over the map and print each key-value pair using the %v formatter.
    for key, value := range languages {
        fmt.Printf("Key for languages %v is value is %v\n", key, value)
    }
}
