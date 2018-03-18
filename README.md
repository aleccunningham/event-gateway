# api-glue
glue relevant apis together 


```Golang
type (    
        // Event is an API endpoint
        Event struct {
            // Event metadata
            EventInfo
            // URI identifier
            URI string `json:"uri" yaml:"uri"`
            // event type
            Type string `json:"type" yaml:"type"`
            // event kind
            Kind string `json:"kind" yaml:"kind"`
        }
        
        // EventInfo holds metadata regarding the endpoint
        EventInfo struct {
            // Endpoint URL
            Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`
            // Description 
            Description string `json:"description,omitempty" yaml:"description,omitempty"`
            // Status of current event for event provider (active, error, not active)
            Status string `json:"status,omitempty" yaml:"status,omitempty"`
        }
        
        // Gluer treats all events as the same
        Gluer interface {
            GetEvent(ctx context.Context, event string) (*Event, error)
            GetEvents(ctx context.Context, eventType, kind string) ([]Event, error)
            CreateEvent(ctx context.Context, eventType, kind, secret, context string, values map[string]string) (*Event, error)
            DeleteEvent(ctx context.Context, event, context string) error
        }    
)

```
