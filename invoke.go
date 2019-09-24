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

        var nc string
        var ec string
        var dbc string

        if len(nodeCount) > 0 {
            nc = nodeCount[0]
        } else {
            nc = "null"
        }

        if len(env) > 0 {
            ec = env[0]
        } else {
            ec = "null"
        }

        if len(database) > 0 {
            dbc = database[0]
        } else {
            dbc = "null"
        }

        if len(nodeCount) > 0 {
            cmd := exec.CommandContext(r.Context(), "/bin/bash", "deploy.sh", nc, ec, dbc)
            cmd.Stderr = os.Stderr
            out, err := cmd.Output()

            if err != nil || cmd.ProcessState.ExitCode() > 0 {
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
