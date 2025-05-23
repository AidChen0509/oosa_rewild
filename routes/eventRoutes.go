package routes

import (
	"oosa_rewild/internal/middleware"
	"oosa_rewild/pkg/repository"

	"github.com/gin-gonic/gin"
)

func EventRoutes(r gin.IRouter) gin.IRouter {
	repo := repository.EventRepository{}
	repoMessageBoard := repository.EventMessageBoardRepository{}
	repoAnnouncement := repository.EventAnnouncementRepository{}
	//repoDetail := repository.PocketListItemsRepository{}
	repoReferenceLinks := repository.EventReferenceLinksRepository{}
	repoSchedule := repository.EventScheduleRepository{}
	repoAccounting := repository.EventAccountingRepository{}
	repoParticipants := repository.EventParticipantsRepository{}
	repoInvitation := repository.EventInvitationMessageRepository{}

	main := r.Group("/event")
	{
		main.GET("", repo.Retrieve)
		main.POST("", middleware.AuthMiddleware(), repo.Create)
		// main.GET("/references", repo.Options)
	}

	detail := main.Group("/:id")
	{
		detail.GET("", repo.Read)
		detail.PUT("", middleware.AuthMiddleware(), repo.Update)
		detail.DELETE("", middleware.AuthMiddleware(), repo.Delete)
	}

	messageBoard := detail.Group("/message-board", middleware.AuthMiddleware())
	{
		messageBoard.GET("", repoMessageBoard.Retrieve)
		messageBoard.POST("", repoMessageBoard.Create)
		messageBoard.GET("/:messageBoardId", repoMessageBoard.Read)
		messageBoard.PUT("/:messageBoardId", repoMessageBoard.Update)
		messageBoard.DELETE("/:messageBoardId", repoMessageBoard.Delete)
		messageBoard.POST("/:messageBoardId/pin", repoMessageBoard.Pin)
	}

	announcement := detail.Group("/announcement", middleware.AuthMiddleware())
	{
		announcement.GET("", repoAnnouncement.Retrieve)
		announcement.POST("", repoAnnouncement.Create)
		announcement.GET("/:messageBoardId", repoAnnouncement.Read)
		announcement.PUT("/:messageBoardId", repoAnnouncement.Update)
		announcement.DELETE("/:messageBoardId", repoAnnouncement.Delete)
	}

	referenceLinks := detail.Group("/reference-links", middleware.AuthMiddleware())
	{
		referenceLinks.GET("", repoReferenceLinks.Retrieve)
		referenceLinks.POST("", repoReferenceLinks.Create)
		referenceLinks.GET("/:referenceLinkId", repoReferenceLinks.Read)
		referenceLinks.PUT("/:referenceLinkId", repoReferenceLinks.Update)
		referenceLinks.DELETE("/:referenceLinkId", repoReferenceLinks.Delete)
	}

	schedule := detail.Group("/schedule", middleware.AuthMiddleware())
	{
		schedule.GET("", repoSchedule.Retrieve)
		schedule.POST("", repoSchedule.Create)
		schedule.DELETE("", repoSchedule.DeleteAll)
		schedule.GET("/:scheduleId", repoSchedule.Read)
		schedule.PUT("/:scheduleId", repoSchedule.Update)
		schedule.DELETE("/:scheduleId", repoSchedule.Delete)
	}

	accounting := detail.Group("/accounting", middleware.AuthMiddleware())
	{
		accounting.GET("", repoAccounting.Retrieve)
		accounting.POST("", repoAccounting.Create)
		accounting.GET("/:accountingId", repoAccounting.Read)
		accounting.PUT("/:accountingId", repoAccounting.Update)
		accounting.DELETE("/:accountingId", repoAccounting.Delete)
	}

	participants := detail.Group("/participants", middleware.AuthMiddleware())
	{
		participants.GET("", repoParticipants.Retrieve)
		participants.POST("", repoParticipants.Create)
		// participants.PUT("/:accountingId", repoParticipants.Update)
		participants.DELETE("/:participantId", repoParticipants.Delete)
	}
	main.GET("/participants/:participantId", repoParticipants.Read)

	invitation := detail.Group("/invitation", middleware.AuthMiddleware())
	{
		invitation.PUT("", repoInvitation.Update)
	}

	join := detail.Group("/join", middleware.AuthMiddleware())
	{
		join.POST("", repoInvitation.Join)
	}

	return r
}
