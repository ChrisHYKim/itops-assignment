package repository

import (
	"fmt"
	"itops-assignment/backend/internal/model"
	"sync"
	"sync/atomic"
)

// 이슈 생성, 이슈 로그, 전체 로그, 업데이트 로그, 삭제 로그 구성
type IssueRepository interface {
	CreateIssue(issue model.Issue) (model.Issue, error)
	GetIssueByID(id uint) (*model.Issue, error)
	GetAllIssues(statusFilter string) ([]model.Issue, error)
	UpdateIssue(issue model.Issue) (model.Issue, error)
	DeleteIssue(id uint) error
}
type InMemoryIssueRepository struct {
	issues map[uint]model.Issue
	// Read/Write Mutex 구성
	mu          sync.RWMutex
	nextIssueId uint32
}

// InMemoryIssue 탐색
func NewInMemoryIssueRepository() *InMemoryIssueRepository {
	return &InMemoryIssueRepository{
		issues:      make(map[uint]model.Issue),
		nextIssueId: 0, // 초기 0 지정
	}
}

// 이슈 크레커 생성
func (r *InMemoryIssueRepository) CreateIssue(issue model.Issue) (model.Issue, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	newID := atomic.AddUint32(&r.nextIssueId, 1) // Increment ID atomically
	issue.ID = uint(newID)
	r.issues[issue.ID] = issue
	return issue, nil
}

func (r *InMemoryIssueRepository) GetIssueByID(id uint) (*model.Issue, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	issue, ok := r.issues[id]
	if !ok {
		return nil, fmt.Errorf("issue with ID %d not found", id)
	}
	return &issue, nil
}

// 전체 이슈 발생 상황 탐지
func (r *InMemoryIssueRepository) GetAllIssues(statusFilter string) ([]model.Issue, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	//  filter issue 정의
	var filteredIssue []model.Issue
	for _, issue := range r.issues {
		if statusFilter == "" || issue.Status == statusFilter {
			filteredIssue = append(filteredIssue, issue)
		}
	}
	// filterissue 및 null 반환
	return filteredIssue, nil
}

// 업데이트 중에 발생한 이슈 로그 기록 진행
func (r *InMemoryIssueRepository) UpdateIssue(updatedIssue model.Issue) (model.Issue, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// update issue 발생 시 에러 발생시 출력
	if _, OK := r.issues[updatedIssue.ID]; !OK {
		return model.Issue{}, fmt.Errorf("issue with ID %d not found for update", updatedIssue.ID)
	}
	r.issues[updatedIssue.ID] = updatedIssue
	return updatedIssue, nil
}

// 삭제 로그 기록
func (r *InMemoryIssueRepository) DeleteIssue(id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.issues[id]; !ok {
		return fmt.Errorf("issue with ID %d not found for deletion", id)
	}
	delete(r.issues, id)
	return nil
}
