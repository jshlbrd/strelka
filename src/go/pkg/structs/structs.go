package structs

import (
        "time"

        "google.golang.org/grpc"

        "github.com/target/strelka/src/go/api/strelka"
)

// defines structures used in configuration files
type ConfConn struct {
        Server          string                      // required
        Cert            string                      // required
        Timeout         time.Duration               // required
}

type ConfThroughput struct {
        Concurrency     int                         // required
        Chunk           int                         // required
        Delay           time.Duration               // optional
}

type ConfFiles struct {
        Patterns        []string                    // required
        Delete          bool                        // optional
}

type ConfRedis struct {
        Addr            string                      // required
        Db              int                         // required
}

// determines what action the client takes with responses, defaults to discarding messages
// only one of these is valid at a time
type ConfResponse struct {
        Log             string                      // optional, loads rpc.LogResponses
        Report          time.Duration               // optional, loads rpc.ReportResponses
}

// defines structures of configuration files
type FileShot struct {
        Client          string                      // optional
        Conn            ConfConn                    // required
        Throughput      ConfThroughput              // required
        Files           ConfFiles                   // required
        Response        ConfResponse                // optional
}

type FileStream struct {
        Client          string                      // optional
        Conn            ConfConn                    // required
        Throughput      ConfThroughput              // required
        Files           ConfFiles                   // required
        Response        ConfResponse                // optional
        Staging         string                      // required
        Delta           time.Duration               // required
}

type Frontend struct {
        Server              string                  // required
        Cache               ConfRedis               // required
        Coordinator         ConfRedis               // required
        Response            ConfResponse            // optional
}

type Redis struct {
        Coordinator         ConfRedis               // required
}

// defines options used when sending scan requests
type Options struct {
        Conn                *grpc.ClientConn        // required
        Timeout             time.Duration           // required
}

type ScanFileRequest struct {
        Request             *strelka.Request        // required
        Attributes          *strelka.Attributes     // required
        Chunk               int                     // required
        Delay               time.Duration           // optional
        Delete              bool                    // optional, only use if files must be deleted!
}
