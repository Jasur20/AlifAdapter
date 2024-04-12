package integration

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

