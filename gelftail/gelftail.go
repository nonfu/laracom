package main

import (
    "encoding/json"
    "flag"
    "github.com/nonfu/laracom/gelftail/aggregator"
    "github.com/nonfu/laracom/gelftail/transformer"
    "github.com/sirupsen/logrus"
    "io/ioutil"
    "net"
    "os"
    "sync"
)

var authToken = ""
var port *string

func init() {
    data, err := ioutil.ReadFile("token.txt")
    if err != nil {
        msg := "Cannot find token.txt that should contain our Loggly token"
        logrus.Errorln(msg)
        panic(msg)
    }
    authToken = string(data)

    port = flag.String("port", "12202", "UDP port for the gelftail")
    flag.Parse()
}

func main()  {
    logrus.Println("Starting Gelf-tail server...")

    ServerConn := startUDPServer(*port)   // Remember to dereference the pointer for our "port" flag
    defer ServerConn.Close()

    var bulkQueue = make(chan []byte, 1)  // Buffered channel to put log statements ready for LaaS upload into

    go aggregator.Start(bulkQueue, authToken)          // Start goroutine that'll collect and then upload batches of log statements
    go listenForLogStatements(ServerConn, bulkQueue)   // Start listening for UDP traffic

    logrus.Infoln("Started Gelf-tail server")

    wg := sync.WaitGroup{}
    wg.Add(1)
    wg.Wait()              // Block indefinitely
}

func startUDPServer(port string) *net.UDPConn {
    ServerAddr, err := net.ResolveUDPAddr("udp", ":"+port)
    checkError(err)

    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    checkError(err)

    return ServerConn
}

func checkError(err error) {
    if err != nil {
        logrus.Println("Error: ", err)
        os.Exit(0)
    }
}

func listenForLogStatements(ServerConn *net.UDPConn, bulkQueue chan[]byte) {
    buf := make([]byte, 8192)                        // Buffer to store UDP payload into. 8kb should be enough for everyone, right Bill? :D
    var item map[string]interface{}                  // Map to put unmarshalled GELF json log message into
    for {
        n, _, err := ServerConn.ReadFromUDP(buf)     // Blocks until data becomes available, which is put into the buffer.
        if err != nil {
            logrus.Errorf("Problem reading UDP message into buffer: %v\n", err.Error())
            continue                                 // Log and continue if there are problms
        }

        err = json.Unmarshal(buf[0:n], &item)        // Try to unmarshal the GELF JSON log statement into the map
        if err != nil {                              // If unmarshalling fails, log and continue. (E.g. filter)
            logrus.Errorln("Problem unmarshalling log message into JSON: " + err.Error())
            item = nil
            continue
        }

        // Send the map into the transform function
        processedLogMessage, err := transformer.ProcessLogStatement(item)
        if err != nil {
            logrus.Printf("Problem parsing message: %v", string(buf[0:n]))
        } else {
            bulkQueue <- processedLogMessage          // If processing went well, send on channel to aggregator
        }
        item = nil
    }
}