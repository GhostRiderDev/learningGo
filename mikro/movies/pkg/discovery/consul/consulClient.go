package consul

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/ghostriderdev/movies/pkg/discovery"
)

func RegisterService(port int, host, serviceName, serverAddr string) (*Registry, context.CancelFunc, error) {
    registry, err := NewRegistry(serverAddr)
    if err != nil {
        return nil, nil, err
    }

    ctx, cancel := context.WithCancel(context.Background())

    instanceID := discovery.GenerateInstanceID(serviceName)

    if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("%s:%d", host, port)); err != nil {
        cancel()
        return nil, nil, err
    }

    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
                    log.Printf("%s", "Failed to report healthy state: "+err.Error())
                }
                time.Sleep(1 * time.Second)
            }
        }
    }()

    return registry, cancel, nil
}