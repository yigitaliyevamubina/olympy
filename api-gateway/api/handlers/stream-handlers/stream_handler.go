package streamhandlers

import (
    "log"
    "net/http"

    streamingservice "olympy/api-gateway/genproto/stream_service"

    "github.com/gin-gonic/gin"
    "github.com/k0kubun/pp"
	"github.com/gin-gonic/gin"
)

type StreamHandlers struct {
    client streamingservice.StreamingServiceClient
    logger *log.Logger
}

func NewStreamHandlers(client streamingservice.StreamingServiceClient, logger *log.Logger) *StreamHandlers {
    return &StreamHandlers{
        client: client,
        logger: logger,
    }
	if client == nil {
		logger.Fatal("Client is nil during initialization") // Use Fatal to terminate the application if the client is nil
	}
	return &StreamHandlers{
		client: client,
		logger: logger,
	}
}

// SendEvent godoc
// @Summary Send an event to the streaming service
// @Description This endpoint sends an event to the streaming service.
// @Accept json
// @Produce json
// @Tags Live Streaming
// @Security ApiKeyAuth
// @Param request body streamingservice.StreamEventRequest true "Event details"
// @Success 200 {object} streamingservice.StreamEventResponse
// @Failure 400 {object} streamingservice.StreamEventResponse
// @Failure 500 {object} streamingservice.StreamEventResponse
// @Router /stream/send [post]
func (s *StreamHandlers) SendEvent(ctx *gin.Context) {
    var req streamingservice.StreamEventRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	var req streamingservice.StreamEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    // Debugging logs
    s.logger.Println("StreamHandlers struct:", s)
    s.logger.Printf("StreamHandlers.client is nil: %v\n", s.client == nil)

    if s.client == nil {
        s.logger.Println("StreamHandlers.client is nil, aborting.")
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: client is not initialized"})
        return
    }

    pp.Println(req)
    resp, err := s.client.StreamEvent(ctx, &req)
    if err != nil {
        s.logger.Println("Error calling StreamEvent:", err)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    s.logger.Println("StreamEvent successful")
    ctx.JSON(http.StatusOK, resp)
}
