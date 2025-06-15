package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"itops-assignment/backend/internal/model"
	"itops-assignment/backend/internal/service"
	"itops-assignment/backend/internal/util"

	"github.com/gorilla/mux"
)

// Issue handler 구성
type IssueHandlers struct {
	issueService service.IssueService
}

func NewIssueHandlers(issuesServicee service.IssueService) *IssueHandlers {
	return &IssueHandlers{issueService: issuesServicee}
}

// 이슈 요청 생성
type CreateIssueRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      *uint  `json:"userId"`
}

// 업데이트 이슈 생성
type UpdateIssueReqeust struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	UserID      *uint   `json:"userId"`
}

// GET /issue 조회
func (h *IssueHandlers) CreateIssue(w http.ResponseWriter, r *http.Request) {
	var req CreateIssueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendErrorResponse(w, util.NewError(http.StatusBadRequest, "Invaild Error"))
		return
	}

	issue, err := h.issueService.CreateIssue(req.Title, req.Description, req.UserID)
	if err != nil {
		util.SendErrorResponse(w, err) // service errors are already APIError type
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// 201 HTTP STATUS 반환
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(issue)
}

// 이슈 조회 (/issue)
func (h *IssueHandlers) GetIssues(w http.ResponseWriter, r *http.Request) {
	statusFilter := r.URL.Query().Get("status")

	issues, err := h.issueService.GetAllIssues(statusFilter)
	if err != nil {
		util.SendErrorResponse(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]model.Issue{"issues": issues})
}

// 한 개의 이슈 탐색 (/issue/id)
func (h *IssueHandlers) GetIssueByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		util.SendErrorResponse(w, util.NewError(http.StatusBadRequest, "Invalid issue ID format"))
		return
	}
	issue, err := h.issueService.GetIssueByID(uint(id))
	if err != nil {
		util.SendErrorResponse(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(issue)
}

// update /issue/id : update Request
func (h *IssueHandlers) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		util.SendErrorResponse(w, util.NewError(http.StatusBadRequest, "Invalid issue ID format"))
		return
	}
	// UPDATE Issue Request
	var req UpdateIssueReqeust
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		util.SendErrorResponse(w, util.NewError(http.StatusBadRequest, "Invalid issue Request Body "))
		return
	}
	serviceUpdateReq := service.IssueUpdateRequest{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		UserID:      req.UserID,
	}

	// err 탐지
	issue, err := h.issueService.UpdateIssue(uint(id), serviceUpdateReq)
	if err != nil {
		util.SendErrorResponse(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(issue)
}
