package tags

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"task-manager/internal/controllers/http/auth"
	"task-manager/internal/domain/entities"
	"task-manager/internal/domain/services"
	"task-manager/internal/domain/services/dto"
	"task-manager/internal/pkg"
)

type tagServer interface {
	CreateTag(tag *dto.Tag) (*entities.Tag, error)
	GetTagById(id string) (*entities.Tag, error)
	GetTagByName(name string) (*entities.Tag, error)
	GetTags(searchTerm string) ([]*entities.Tag, error)
	UpdateTag(id string, tag *dto.Tag) error
	DeleteTagById(id string) error
}

type tokenService interface {
	GenerateAccessToken(user *entities.User) (string, error)
	ParseAccessToken(tokenString string) (jwt.MapClaims, error)
	GenerateRefreshToken(id string) (string, error)
	CheckRefreshToken(refreshTokenString string) (string, error)
	DeleteRefreshToken(refreshTokenString string) error
}

type Controller struct {
	tagService tagServer
}

func NewTagsController(gin *gin.RouterGroup, tagService tagServer, tokenService tokenService) *Controller {
	controller := &Controller{
		tagService: tagService,
	}

	r := gin.Group("/tags", auth.CheckTokenMiddleware(tokenService))
	{
		r.GET("", controller.GetTags)
		r.GET("/:id", controller.GetTagById)
		r.POST("", controller.CreateTag)
		r.PATCH("/:id", controller.UpdateTag)
		r.DELETE("/:id", controller.DeleteTagById)
	}

	return controller
}

func (controller *Controller) CreateTag(ctx *gin.Context) {
	var tagDTO TagDTO
	validate := validator.New()

	if err := ctx.ShouldBind(&tagDTO); err != nil {
		logrus.Error("POST /tags | body parsing error :", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := validate.Struct(&tagDTO); err != nil {
		errorMessage := pkg.ExtractValidationErrors(err)

		logrus.Error("POST /tags | body validation error: ", errorMessage)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})

		return
	}

	tag, err := controller.tagService.CreateTag(ToDomainTagDTO(tagDTO))

	if errors.Is(err, services.ExistTagError) {
		logrus.Error("POST /tags | cannot create user :", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": services.ExistTagError.Error()})
		return
	}

	if err != nil {
		logrus.Error("POST /tags | cannot create user :", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id": tag.ID,
	})
}

func (controller *Controller) GetTagById(ctx *gin.Context) {
	id := ctx.Param("id")

	tag, err := controller.tagService.GetTagById(id)
	if err != nil {
		logrus.Error("GET /tags | cannot get tag")
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, toTagResp(tag))
}

func (controller *Controller) GetTags(ctx *gin.Context) {
	searchTerm := ctx.Request.URL.Query().Get("search")

	domainTags, err := controller.tagService.GetTags(searchTerm)
	if err != nil {
		logrus.Error("GET /tags | cannot get tags :", err)
		ctx.JSON(http.StatusOK, gin.H{
			"tags":  []TagResponse{},
			"count": 0,
		})

		return
	}

	var tags []*TagResponse

	for _, tag := range domainTags {
		tags = append(tags, toTagResp(tag))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"tags":  tags,
		"count": len(tags),
	})
}

func (controller *Controller) UpdateTag(ctx *gin.Context) {
	var tagDTO TagDTO
	validate := validator.New()

	if err := ctx.ShouldBind(&tagDTO); err != nil {
		logrus.Error("PATCH /tags | body parsing error :", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validate.Struct(&tagDTO); err != nil {
		errorMessage := pkg.ExtractValidationErrors(err)

		logrus.Error("PATCH /tags | body validation error: ", errorMessage)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	id := ctx.Param("id")

	if err := controller.tagService.UpdateTag(id, ToDomainTagDTO(tagDTO)); err != nil {
		logrus.Error("PATCH /tags | cannot update user :", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (controller *Controller) DeleteTagById(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := controller.tagService.DeleteTagById(id); err != nil {
		logrus.Error("DELETE /tags/:id: | cannot delete user: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
