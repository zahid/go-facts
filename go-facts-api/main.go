package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

/* The Fact model */
type Fact struct {
	Text  string `json:"text"`
	Image string `json:"image"`
}

/* The Facts model is a list of Fact structs */
type Facts []Fact

/* Global fact list taken from https://www.buzzfeed.com/tomphillips/42-incredibly-weird-facts-youll-want-to-tell-people-down-the */
var facts = Facts{
	Fact{Text: "The longest time between two twins being born is 87 days"},
	Fact{Text: "The world’s deepest postbox is in Susami Bay in Japan. It’s 10 metres underwater."},
	Fact{Text: "In 2007, an American man named Corey Taylor tried to fake his own death in order to get out of his cell phone contract without paying a fee. It didn’t work."},
	Fact{Text: "The oldest condoms ever found date back to the 1640s (they were found in a cesspit at Dudley Castle), and were made from animal and fish intestines."},
	Fact{Text: "In 1923, jockey Frank Hayes won a race at Belmont Park in New York despite being dead — he suffered a heart attack mid-race, but his body stayed in the saddle until his horse crossed the line for a 20–1 outsider victory."},
	Fact{Text: "Everyone has a unique tongue print, just like fingerprints."},
	Fact{Text: "Most Muppets are left-handed. (Because most Muppeteers are right-handed, so they operate the head with their favoured hand.)"},
	Fact{Text: "Female kangaroos have three vaginas."},
	Fact{Text: " It costs the U.S. Mint almost twice as much to mint each penny and nickel as the coins are actually worth. Taxpayers lost over $100 million in 2013 just through the coins being made."},
	Fact{Text: "Light doesn’t necessarily travel at the speed of light. The slowest we’ve ever recorded light moving at is 38 mph."},
}

func main() {
	log.Print("Starting")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fact", GetRandomFact)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetRandomFact(w http.ResponseWriter, r *http.Request) {
	log.Print("responding to request")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	randomFact := facts[rand.Intn(len(facts))]
	randomFact.Image = fmt.Sprintf("/images/background-%v", rand.Intn(len(facts)))
	json.NewEncoder(w).Encode(randomFact)
}
