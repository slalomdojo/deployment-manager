package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
        cmd := exec.CommandContext(r.Context(), "/bin/bash", "deploy.sh")
        cmd.Stderr = os.Stderr
        out, err := cmd.Output()

        if err != nil {
            w.WriteHeader(500)
        }

        w.Write(out)
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
        http.HandleFunc("/", handler)
        http.HandleFunc("/templates/gke", gkeHandler)
        http.HandleFunc("/templates/vm", vmHandler)

        port := os.Getenv("PORT")

        if port == "" {
            port = "8080"
        }

        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
