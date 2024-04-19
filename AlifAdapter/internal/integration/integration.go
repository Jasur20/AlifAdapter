package integration

import "strconv"

type PaymentStatus int

const (
	PaymentStatusSuccessfully               PaymentStatus = 200 // успешно
	PaymentStatusConversionError            PaymentStatus = 285 // произошла оошибка при конвертации
	PaymentStatusCourseCurrencyHasChanged   PaymentStatus = 286 // курс валют изменился
	PaymentStatusInvalidRequest             PaymentStatus = 400 // неверный запрос
	PaymentStatusNotAuthorized              PaymentStatus = 401 // не авторизован
	PaymentStatusRecepienWasNotFound        PaymentStatus = 402 // Получатель не найден
	PaymentStatusNoAccess                   PaymentStatus = 403 // нет доступа
	PaymentStatusPaymentNotFound            PaymentStatus = 404 // Платеж не найден
	PaymentStatusMethodNotAllowed           PaymentStatus = 405 // Метод не разрешен
	PaymentStatusReConfirmationPayment      PaymentStatus = 406 // Повторное подтверждение платежа
	PaymentStatusRepeatedVerificationReques PaymentStatus = 409 // Повторный запрос проверки
	PaymentStatusInvalidAccount             PaymentStatus = 410 // неверный аккаунт получателя
	PaymentStatusAmountSmall                PaymentStatus = 411 // Сумма слишком мала
	PaymentStatusAmountLarge                PaymentStatus = 412 // Сумма слишком велика
	PaymentStatusIncorrectTransferAmount    PaymentStatus = 413 // Неверная сумма перевода
	PaymentStatusInvalidRequestIdentidier   PaymentStatus = 414 // Неверный идентификатор запроса
	PaymentStatusClientStopList             PaymentStatus = 415 // Клиент в стоп-листе
	PaymentStatusInternalServerError        PaymentStatus = 500 // Внутренняя ошибка сервера
	PaymentStatusTemporaryErrorRepeatLater  PaymentStatus = 503 // Временная ошибка. Повторите запрос позже
	PaymentStatusPaymentPending             PaymentStatus = 520 // Платеж в ожидании
	PaymentStatusPaymentChecking            PaymentStatus = 521 // Платеж на проверке
)

const(
	ACCEPTED =0
	SUCCESS=1
	PENDING=2
	FAILED=3
	CANCELED=4
)


type Adapter interface {
	PreCheck(account string, serviceID string) (status int64, description string, rawInfo map[string]string, err error)
	//Payment(account, amount, trnID, notifyRoute string) (status int64, description string, paymentID int64, err error)
	Payment(account, serviceID, amount, trnID, notifyRoute string) (status int64, description string, paymentID string, err error)
	PostCheck(trnID string) (status int64, description string, err error)
}

type Integration interface {
	ReceiverInfo(operID string, req *GetReceiverInfoRequestBody) *GetReceiverInfoResponseBody
	Payment(operID string, req *PaymentRequestBody) *PaymentResponseBody
	PostCheck(operID string, req *PostCheckRequestBody) *PostCheckResponseBody
}

type integration struct {
	adapter Adapter
}

func (i integration) ReceiverInfo(operID string, req *GetReceiverInfoRequestBody) *GetReceiverInfoResponseBody {
	var resp GetReceiverInfoResponseBody
	status, desc, rawInfo, err := i.adapter.PreCheck(req.Account, req.ProviderServiceID)
	if err != nil {
		resp.Status.Code = int(FAILED)
		resp.Status.Message = err.Error()
		return &resp
	}

	switch status {
	case 1:
		resp.Status.Code = int(SUCCESS)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	default:
		resp.Status.Code = int(FAILED)
		resp.Status.Message = desc
		resp.ReceiverInfo = rawInfo
	}
	return &resp
}


func (i integration) Payment(operID string, req *PaymentRequestBody) *PaymentResponseBody {
	var resp PaymentResponseBody

	amount, err := strconv.ParseFloat(req.ReceiverAmount, 64)

	if err != nil {
		resp.Status.Code = int(FAILED)
		resp.Status.Message = err.Error()
		return &resp
	}

	status, desc, payID, err := i.adapter.Payment(req.Account, req.ProviderServiceID, fmt.Sprintf("%.2f", amount), fmt.Sprint(req.ID), "")
	if err != nil {
		resp.Status.Code = int(FAILED)
		resp.Status.Message = err.Error()
		return &resp
	}
	switch status {
	case 310:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID = payID
	case 320:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = desc
	default:
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = desc
		resp.ReceiverTrnID = payID
	}
	return &resp
}


func (i integration) PostCheck(operID string, req *PostCheckRequestBody) *PostCheckResponseBody {
	var resp PostCheckResponseBody
	status, desc, err := i.adapter.PostCheck(fmt.Sprint(req.ID))
	if err != nil {
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = err.Error()
		return &resp
	}
	switch status {
	case 310:
		resp.Status.Code = int(PaymentStatusProcessedByGateway)
		resp.Status.Message = desc
	case 320, 404:
		resp.Status.Code = int(PaymentStatusRejectedByGateway)
		resp.Status.Message = desc
	default:
		resp.Status.Code = int(PaymentStatusSendedToGateway)
		resp.Status.Message = desc
	}
	return &resp
}