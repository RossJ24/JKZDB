package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

var shardConfigPath = flag.String("shard-config", "../shard-config.json", "JSON file representing shard configuration.")

func main() {
	flag.Parse()
	configFile, err := os.ReadFile(*shardConfigPath)
	if err != nil {
		log.Fatalf("Could not parse shard configuration file.")
	}
	var config map[string][]map[string]string
	json.Unmarshal(configFile, &config)
	var wg sync.WaitGroup
	for _, sh := range config["shards"] {
		wg.Add(1)
		go func(sh map[string]string) {
			println("--port=" + sh["port"])
			var cmd *exec.Cmd
			os := runtime.GOOS
			switch os {
			case "windows":
				cmd = exec.Command("stop_server.bat", sh["port"])
			case "darwin":
				cmd = exec.Command("bash", "stop_server.sh", sh["port"])
			case "linux":
				cmd = exec.Command("bash", "stop_server.sh", sh["port"])
			default:
				break
			}
			errBuf := &bytes.Buffer{}
			cmd.Stderr = errBuf
			out, err := cmd.Output()

			if err != nil {
				fmt.Println(errBuf.String())
				log.Fatalf(err.Error())
			}
			fmt.Println(string(out))

			if os == "windows" {
				cmd = exec.Command("db_file_cleanup.bat", "..\\jkzdb-"+sh["port"]+".db", "..\\jkzdb-"+sh["port"]+".db.lock")
			} else {
				cmd = exec.Command("bash", "db_file_cleanup.sh", "../jkzdb-"+sh["port"]+".db")
			}
			errBuf = &bytes.Buffer{}
			cmd.Stderr = errBuf

			out, err = cmd.Output()
			if err != nil {
				fmt.Println(errBuf.String())
				log.Fatalf(err.Error())
			}
			fmt.Println(string(out))
			wg.Done()
		}(sh)
	}

	wg.Wait()
}
