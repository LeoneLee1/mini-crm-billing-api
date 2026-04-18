package transactioncreate

import (
	httpresputils "mini-crm-billing-api/source/common/glob_utils/http_resp_utils"
	jwtutils "mini-crm-billing-api/source/common/glob_utils/jwt_utils"
	"mini-crm-billing-api/source/features/public/transaction/transaction_create/body"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Impl(c *gin.Context) {
	var req body.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errMSG := "invalid request body: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
	}

	userID, ok := jwtutils.GetCurrentUserID(c)
	if !ok {
		errMSG := "Unauthorized"
		httpresputils.HttpResponseUnAuth(c, &errMSG)
		return
	}
	userIDParse, err := uuid.Parse(userID)
	if err != nil {
		errMSG := "invalid user id"
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}
	createdBy := userIDParse

	transaction, err := h.usecase.create(c.Request.Context(), &req, createdBy)
	if err != nil {
		errMSG := "Failed to create transaction: " + err.Error()
		httpresputils.HttpRespBadRequest(c, &errMSG)
		return
	}

	msg := "Transaction created successfully"
	httpresputils.HttpRespCreated(c, transaction, &msg)
}
