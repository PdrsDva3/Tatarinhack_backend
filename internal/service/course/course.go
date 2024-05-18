package courseserv

import (
	"Tatarinhack_backend/internal/entities"
	"Tatarinhack_backend/internal/repository"
	"Tatarinhack_backend/internal/service"
	"context"
)

func InitCourseService(CourseRepo repository.CourseRepo) service.CourseService {
	return CourseService{CourseRepo: CourseRepo}
}

type CourseService struct {
	CourseRepo repository.CourseRepo
}

func (crs CourseService) GetMe(ctx context.Context, courseID int) (*entities.Course, error) {
	out, err := crs.CourseRepo.GetByID(ctx, courseID)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (crs CourseService) Create(ctx context.Context, course entities.CourseBase) (int, error) {
	id, err := crs.CourseRepo.Create(ctx, course)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (crs CourseService) AddTest(ctx context.Context, courseID int, testID int) error {
	err := crs.CourseRepo.AddT(ctx, testID, courseID)
	if err != nil {
		return err
	}
	return nil
}
