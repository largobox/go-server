package main

import (
  "log"
  "net/http"

  "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
  CheckOrigin: func(r *http.Request) bool {
    return true
  },
}

func handler(w http.ResponseWriter, r *http.Request) {
  conn, err := upgrader.Upgrade(w, r, nil)

  if err != nil {
    log.Println(err)

    return
  }

  defer conn.Close()

  for {
    messageType, p, err := conn.ReadMessage()

    if err != nil {
      log.Println(err)

      break
    }

    switch messageType {
      case websocket.TextMessage:
        text := string(p)

        log.Printf("Received message: %s\n", text)

        err = conn.WriteMessage(websocket.TextMessage, []byte("some-string-from-GO-server"))

        if err != nil {
          log.Println(err)
          break
        }
    }
  }
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":3000", nil))
}
