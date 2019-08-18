package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "os/exec"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("You successfully deployed this application to Cloud Run! Now let the fun begin!"))
}

func gkeHandler(w http.ResponseWriter, r *http.Request) {
        nodeCount := r.URL.Query()["node_count"]
        env := r.URL.Query()["env"]
        database := r.URL.Query()["database"]

        if len(nodeCount) > 0 {
            cmd := exec.CommandContext(r.Context(), "/bin/bash", "deploy.sh", nodeCount[0], env[0], database[0])
            cmd.Stderr = os.Stderr
            out, err := cmd.Output()

            if err != nil {
                w.WriteHeader(500)
            }

            w.Write(out)
        } else {
            w.WriteHeader(500)
            w.Write([]byte("Error: please provide node_count"))
        }
}

func main() {
        http.HandleFunc("/", rootHandler)
        http.HandleFunc("/templates/gke", gkeHandler)

        port := os.Getenv("PORT")

        if port == "" {
            port = "8080"
        }

        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
