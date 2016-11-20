package main

import "fmt"

func location(name, city string) (name, country string) {
    switch city {
        case "New York", "LA", "Chicago":
            country = "North America"
        default:
            country = "Unknown"
    }

    return city, country
}

func main(){
    name, country := location("Matt","LA")
    fmt.Printf("%s lives in %s", name, country)
}
