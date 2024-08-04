package streamhandlers

import (
	"log"
	"net/http"

	streamingservice "olympy/api-gateway/genproto/streaming_service"

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
}

// SendEvent godoc
// @Summary Send an event to the streaming service
// @Description This endpoint sends an event to the streaming service.
// @Accept json
// @Produce json
// @Param request body streamingservice.StreamEventRequest true "Event details"
// @Success 200 {object} streamingservice.StreamEventResponse
// @Failure 400 {object} streamingservice.Message
// @Failure 500 {object} streamingservice.Message
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

// StreamEvents godoc
// @Summary Stream events in real-time
// @Description This endpoint streams events in real-time to registered users.
// @Accept json
// @Produce json
// @Param event_id path string true "Event ID"
// @Success 200 {object} streamingservice.StreamEventResponse
// @Failure 400 {object} streamingservice.Message
// @Failure 500 {object} streamingservice.Message
// @Router /stream/stream/{event_id} [get]
func (s *StreamHandlers) StreamEvents(ctx *gin.Context) {
	eventID := ctx.Param("event_id")

	req := &streamingservice.StreamEventRequest{EventId: eventID}

	stream, err := s.client.StreamEvent(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for {
		event, err := stream.Recv()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.SSEvent("event", event)
	}
}
