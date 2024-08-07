package streamhandlers

import (
	"log"
	"net/http"

	streamingservice "olympy/api-gateway/genproto/stream_service"

	"github.com/gin-gonic/gin"
)

type StreamHandlers struct {
	client streamingservice.StreamingServiceClient
	logger *log.Logger
}

func NewStreamHandlers(client streamingservice.StreamingServiceClient, logger *log.Logger) *StreamHandlers {
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
// @Param request body streamingservice.StreamEventRequest true "Event details"
// @Success 200 {object} streamingservice.StreamEventResponse
// @Failure 400 {object} streamingservice.StreamEventResponse
// @Failure 500 {object} streamingservice.StreamEventResponse
// @Router /stream/send [post]
func (s *StreamHandlers) SendEvent(ctx *gin.Context) {
	var req streamingservice.StreamEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := s.client.StreamEvent(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
