package updatepassword

import (
	"context"
	"errors"
	hashingpassword "mini-crm-billing-api/source/common/glob_utils/hashing_password"
	"mini-crm-billing-api/source/features/public/auth/update_password/body"
)

// update implements [Usecase].
func (u *usecaseImpl) Update(ctx context.Context, userID string, req body.UpdateRequest) error {
	if req.NewPassword != req.ConfirmPassword {
		return errors.New("New password and confirm password do not match")
	}

	if req.NewPassword == req.OldPassword {
		return errors.New("New password cannot be the same as the old password")
	}

	user, err := u.repo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if !hashingpassword.VerifyPassword(req.OldPassword, user.Password) {
		return errors.New("Old password is incorrect")
	}

	hashedPassword, err := hashingpassword.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	if err := u.repo.UpdatePassword(ctx, userID, hashedPassword); err != nil {
		return err
	}

	return nil
}
