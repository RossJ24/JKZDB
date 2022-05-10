package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
)

var shardConfigPath = flag.String("shard-config", "../shard-config.json", "JSON file representing shard configuration.")

func main() {
	flag.Parse()
	fmt.Println(*shardConfigPath)
	configFile, err := os.ReadFile(*shardConfigPath)
	if err != nil {
		log.Fatalf("Could not parse shard configuration file.")
	}
	os.Chdir("..")
	var config map[string][]map[string]string
	json.Unmarshal(configFile, &config)
	var wg sync.WaitGroup
	for _, sh := range config["shards"] {
		wg.Add(1)
		go func(sh map[string]string) {
			println("--port=" + sh["port"])
			cmd := exec.Command("go", "run", "db/db_server.go", "--port="+sh["port"])
			err := cmd.Start()
			errBuf := &bytes.Buffer{}
			cmd.Stderr = errBuf

			out, err := cmd.Output()
			if err != nil {
				fmt.Println(errBuf.String())
			}
			fmt.Println(string(out))
			wg.Done()
		}(sh)
	}

	wg.Wait()
}
