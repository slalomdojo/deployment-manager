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

        if len(nodeCount) > 0 {
            cmd := exec.CommandContext(r.Context(), "/bin/bash", "deploy.sh", "gke", nodeCount[0])
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

func vmHandler(w http.ResponseWriter, r *http.Request) {
    vmCount := r.URL.Query()["vm_count"]

    if len(vmCount) > 0 {
        cmd := exec.CommandContext(r.Context(), "/bin/bash", "deploy.sh", "vm", vmCount[0])
        cmd.Stderr = os.Stderr
        out, err := cmd.Output()

        if err != nil {
            w.WriteHeader(500)
        }

        w.Write(out)
    } else {
        w.WriteHeader(500)
        w.Write([]byte("Error: please provide vm_count"))
    }
}

func main() {
        http.HandleFunc("/", rootHandler)
        http.HandleFunc("/templates/gke", gkeHandler)
        http.HandleFunc("/templates/vm", vmHandler)

        port := os.Getenv("PORT")

        if port == "" {
            port = "8080"
        }

        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
