package v1

import (
	fmt "fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	profiles_v1 "github.com/videocoin/cloud-api/profiles/v1"
)

// CheckErrs used to ensure a job is valid before passing to transcoder
func (s *Stream) CheckErrs() []error {
	var errs = make([]error, 0)

	err := fmt.Errorf("invalid profile")
	for _, p := range []profiles_v1.ProfileId{profiles_v1.ProfileIdFHD, profiles_v1.ProfileIdHD, profiles_v1.ProfileIdSD} {
		if p == s.ProfileId {
			err = nil
		}
	}

	if err != nil {
		errs = append(errs, err)
	}

	if s.ClientAddress == "" {
		errs = append(errs, fmt.Errorf("invalid client address"))
	}

	if s.StreamContractId <= 0 {
		errs = append(errs, fmt.Errorf("invalid stream id"))
	}

	if s.StreamContractAddress == "" {
		errs = append(errs, fmt.Errorf("invalid stream address"))
	}

	return errs
}

func (j *Job) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()

	err := scope.SetColumn("id", uuid.String())
	if err != nil {
		return err
	}

	return scope.SetColumn("created_at", time.Now().Unix())
}

func (j *Job) BeforeSave(scope *gorm.Scope) error {
	if err := scope.SetColumn("updated_at", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}

func (j *Job) BeforeUpdate(scope *gorm.Scope) error {
	if err := scope.SetColumn("updated_at", time.Now().Unix()); err != nil {
		return err
	}

	return nil
}
