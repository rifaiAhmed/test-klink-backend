package api

import (
	"backend-test/helpers"
	"backend-test/internal/interfaces"
	"backend-test/internal/models"
	"backend-test/internal/models/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IMemberHandler struct {
	MemberService interfaces.IMemberService
}

func NewMemberHandler(service interfaces.IMemberService) *IMemberHandler {
	return &IMemberHandler{
		MemberService: service,
	}
}

// CreateMember POST /members
func (api *IMemberHandler) CreateMember(c *gin.Context) {
	var req dto.CreateMemberRequest
	log := helpers.Logger

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, "failed to parse request", nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, "failed to validate request", err.Error())
		return
	}

	member := models.Member{
		Nama:         req.Nama,
		JenisKelamin: req.JenisKelamin,
		NoKtp:        req.NoKtp,
		TempatLahir:  req.TempatLahir,
		TanggalLahir: req.TanggalLahir,
		NoHp:         req.NoHp,
		Email:        req.Email,
		NoRekening:   req.NoRekening,
		ManagerID:    req.ManagerID,
	}

	created, err := api.MemberService.CreateMember(c.Request.Context(), &member, req.PaketID)
	if err != nil {
		log.Error("failed to create member: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, "failed to create member", err.Error())
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, "member created successfully", created)
}

// GetMemberByID GET /members/:id
func (api *IMemberHandler) GetMemberByID(c *gin.Context) {
	id := c.Param("id")
	log := helpers.Logger

	member, err := api.MemberService.GetMemberByID(c.Request.Context(), id)
	if err != nil {
		log.Error("failed to get member: ", err)
		helpers.SendResponseHTTP(c, http.StatusNotFound, "member not found", err.Error())
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, "success", member)
}

// GetAllMembers GET /members
func (api *IMemberHandler) GetAllMembers(c *gin.Context) {
	log := helpers.Logger

	limit := 10
	skip := 0
	sortBy := "created_at"
	sortType := "desc"
	search := ""

	if val := c.Query("limit"); val != "" {
		fmt.Sscanf(val, "%d", &limit)
	}
	if val := c.Query("skip"); val != "" {
		fmt.Sscanf(val, "%d", &skip)
	}
	if val := c.Query("sortby"); val != "" {
		sortBy = val
	}
	if val := c.Query("sorttype"); val != "" {
		sortType = val
	}
	if val := c.Query("search"); val != "" {
		search = val
	}

	objComponent := models.ComponentServerSide{
		Limit:    limit,
		Skip:     skip,
		SortBy:   sortBy,
		SortType: sortType,
		Search:   search,
	}

	members, total, err := api.MemberService.GetAllMembers(c.Request.Context(), objComponent)
	if err != nil {
		log.Error("failed to get members: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, "failed to get members", err.Error())
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, "success", gin.H{
		"total": total,
		"data":  members,
	})
}

// UpdateMember PUT /members/:id
func (api *IMemberHandler) UpdateMember(c *gin.Context) {
	id := c.Param("id")
	var req models.Member
	log := helpers.Logger

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, "failed to parse request", nil)
		return
	}

	req.ID = id

	member, err := api.MemberService.UpdateMember(c.Request.Context(), &req)
	if err != nil {
		log.Error("failed to update member: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, "failed to update member", err.Error())
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, "member updated successfully", member)
}

// DeleteMember DELETE /members/:id
func (api *IMemberHandler) DeleteMember(c *gin.Context) {
	id := c.Param("id")
	log := helpers.Logger

	if err := api.MemberService.DeleteMember(c.Request.Context(), id); err != nil {
		log.Error("failed to delete member: ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, "failed to delete member", err.Error())
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, "member deleted successfully", nil)
}
