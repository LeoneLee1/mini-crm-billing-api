package updatepassword

import (
	"context"
	"errors"
	hashingpassword "mini-crm-billing-api/source/common/glob_utils/hashing_password"

	"gorm.io/gorm"
)

// update implements [Usecase].
func (u *usecaseImpl) update(ctx context.Context, id string, req updateRequest) error {
	user, err := u.repo.findByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}

	if req.NewPassword != req.ConfirmPassword {
		return errors.New("New password and confirm password do not match")
	}

	if req.NewPassword == req.OldPassword {
		return errors.New("New password cannot be the same as the old password")
	}

	if !hashingpassword.VerifyPassword(req.OldPassword, user.Password) {
		return errors.New("Old password is incorrect")
	}

	hashedPassword, err := hashingpassword.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	if err := u.repo.updatePassword(ctx, id, hashedPassword); err != nil {
		return err
	}

	return nil
}
