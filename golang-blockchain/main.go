package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"golang-blockchain/blockchain"

	"github.com/gorilla/mux"
)

type BlockSuccess struct {
	Message string `json:"message"`
}
type ServerSetup struct {
	Status string
}

func getBlocks(w http.ResponseWriter, r *http.Request) {

	var tmpRecords []blockchain.Block
	iter := blockchain.InitBlockChain().Iterator()

	for {
		block := iter.Next()

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		tmpRecords = append(tmpRecords, *block)
		if len(block.PrevHash) == 0 {
			break
		}
	}
	setupHeader(w)
	json.NewEncoder(w).Encode(tmpRecords)
	iter.Database.Close()
}

func viewCurrentBlock(w http.ResponseWriter, r *http.Request) {

	iter := blockchain.InitBlockChain()
	setupHeader(w)
	json.NewEncoder(w).Encode(iter)
	iter.Database.Close()
}

func checkServer(w http.ResponseWriter, r *http.Request) {

	var newEmployee = ServerSetup{Status: "Server is in running state"}
	setupHeader(w)
	json.NewEncoder(w).Encode(newEmployee)
}

func createBlock(w http.ResponseWriter, r *http.Request) {
	setupHeader(w)
	var inst = blockchain.InitBlockChain()
	data := r.FormValue("data")
	inst.AddBlock(data)
	json.NewEncoder(w).Encode(BlockSuccess{Message: "Block has been added"})
	inst.Database.Close()
}

func setupHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", checkServer)
	router.HandleFunc("/addblock", createBlock).Methods("POST")
	router.HandleFunc("/getblocks", getBlocks).Methods("GET")
	router.HandleFunc("/viewCurrentBlock", viewCurrentBlock).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
