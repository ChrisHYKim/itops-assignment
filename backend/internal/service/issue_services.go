package service

import (
	"fmt"
	"itops-assignment/backend/internal/model"
	"itops-assignment/backend/internal/repository"
	"itops-assignment/backend/internal/util"
	"strings"
	"time"
)

// 이슈 서비스 인터페이스 구성
type IssueService interface {
	CreateIssue(title, description string, userID *uint) (model.Issue, error)
	GetIssueByID(id uint) (*model.Issue, error)
	GetAllIssues(statusFilter string) ([]model.Issue, error)
	UpdateIssue(id uint, updateReq IssueUpdateRequest) (model.Issue, error)
}
type issueService struct {
	repo repository.IssueRepository
}

// Issue 생성
func NewIssueService(repo repository.IssueRepository) IssueService {
	return &issueService{repo: repo}
}

type IssueUpdateRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Status      *string `json:"status"`
	UserID      *uint   `json:"userId"`
}

func (s *issueService) CreateIssue(title, description string, userID *uint) (model.Issue, error) {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)
	var validationErr []string
	if title == "" {
		validationErr = append(validationErr, "title canot be empty")
	}
	if description == "" {
		validationErr = append(validationErr, "description cannot be empty")
	}
	// title 및 description 유효성 체크 진행
	if len(validationErr) > 0 {
		return model.Issue{}, util.NewError(400, strings.Join(validationErr, ", "))
	}
	newIssue := model.Issue{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if userID != nil {
		user := model.GetUserByID(*userID)
		if user == nil {
			return model.Issue{}, util.NewError(400, fmt.Sprintf("user with ID %d not found", *userID))
		}
		newIssue.User = user
		newIssue.Status = "IN_PROGRESS"
	} else {
		newIssue.Status = "PENDING"
	}
	createdIssue, err := s.repo.CreateIssue(newIssue)
	if err != nil {
		return model.Issue{}, util.NewError(500, fmt.Sprintf("failed to create issue: %v", err))
	}
	return createdIssue, nil
}

// Issue 체크
func (s *issueService) GetIssueByID(id uint) (*model.Issue, error) {
	issue, err := s.repo.GetIssueByID(id)
	if err != nil {
		return nil, util.NewError(404, fmt.Sprintf("issue with ID %d not found", id))
	}
	return issue, nil
}
func (s *issueService) GetAllIssues(statusFilter string) ([]model.Issue, error) {
	if statusFilter != "" && !model.IsValidStatus(statusFilter) {
		return nil, util.NewError(400, fmt.Sprintf("invalid status filter: %s", statusFilter))
	}

	issues, err := s.repo.GetAllIssues(statusFilter)
	if err != nil {
		return nil, util.NewError(500, fmt.Sprintf("failed to retrieve issues: %v", err))
	}
	return issues, nil
}

// upodate issue
func (s *issueService) UpdateIssue(id uint, updateReq IssueUpdateRequest) (model.Issue, error) {
	existingIssue, err := s.repo.GetIssueByID(id)
	if err != nil {
		return model.Issue{}, fmt.Errorf("issue with ID %d not found", id)
	}
	if existingIssue.Status == "COMPLTED" || existingIssue.Status == "CANCELLED" {
		return model.Issue{}, fmt.Errorf("cannot update a %s issue", existingIssue.Status)
	}

	// 새로운 값을 적용
	updateIssue := *existingIssue
	// title 값이 존재하거나 Description 존재할 경우
	if updateReq.Title != nil {
		updateIssue.Title = *updateReq.Title
	}
	if updateReq.Description != nil {
		updateIssue.Description = *updateReq.Description
	}
	// userID 변경 진행
	assignesChaned := false
	if updateReq.UserID != nil {
		assignesChaned = true
		if *updateReq.UserID == 0 {
			updateIssue.User = nil
		} else {
			user := model.GetUserByID(*updateReq.UserID)
			if user == nil {
				return model.Issue{}, fmt.Errorf("assigned user failed %d not found", *&updateReq.UserID)
			}
			updateIssue.User = user
		}
	}
	// 상태 변경 진행
	statusChanged := false
	if updateReq.Status != nil { // status 필드가 요청에 포함된 경우
		if !model.IsValidStatus(*updateReq.Status) {
			return model.Issue{}, fmt.Errorf("invalid status provided: %s", *updateReq.Status)
		}
		updateIssue.Status = *updateReq.Status
		statusChanged = true
	}
	// PENDING 상태 반환 시, 특정 사용자를 제거한다
	if assignesChaned && updateIssue.User == nil {
		updateIssue.Status = "PENDING"
	}
	if existingIssue.User == nil && updateIssue.User != nil && !statusChanged {
		if existingIssue.Status == "PENDING" { // 담당자 할당 시 PENDING에서 IN_PROGRESS로 자동 변경
			updateIssue.Status = "IN_PROGRESS"
		}
	}
	if updateIssue.User == nil && (updateIssue.Status == "IN_PROGRESS" || updateIssue.Status == "COMPLETED") {
		return model.Issue{}, fmt.Errorf("cannot change status to %s without an assignee", updateIssue.Status)
	}
	updateIssue.UpdatedAt = time.Now()

	// 변경된 이슈 저장소에 반영
	resultIssue, err := s.repo.UpdateIssue(updateIssue)
	if err != nil {
		return model.Issue{}, fmt.Errorf("failed to update issue in repository: %w", err)
	}
	return resultIssue, nil
}
