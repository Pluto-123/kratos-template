package service

import (
	"fmt"
	"kratos-template/internal/biz"
	"time"
)

var DefaultJobs map[string]JobFunc

type JobFunc func()

type JobService struct {
	uc *biz.HelloUsecase
}

func NewJobService(uc *biz.HelloUsecase) *JobService {
	job := &JobService{
		uc: uc,
	}
	return job
}

func (s *JobService) Init() {
	DefaultJobs = map[string]JobFunc{
		"one": s.DoMyWork,
	}
}
func (s *JobService) DoMyWork() {
	//s.uc.CreateGreeter(context.Background(), &biz.Greeter{})
	fmt.Printf("当前时间 %v \n", time.Now().Unix())
}
