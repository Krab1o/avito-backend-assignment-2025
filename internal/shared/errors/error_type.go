package errs

type ErrorType int

const (
	//Repo layer errors	
	DBError ErrorType = iota // ошибка уровня базы данных

	// Service layer errors
	SemanticError // ошибка логики действия (например, передача денег самому себе)
	NotFound // какой-либо ресурс не найден
	Unauthorized // не авторизован
	ServiceError // другая ошибка сервисного слоя
)


func (e ErrorType) String() string {
	switch e {
	case DBError:
		return "Database Error"
	case SemanticError:
		return "Semantic Error"
	case NotFound:
		return "Not Found Error"
	case Unauthorized:
		return "Authorization Error"
	case ServiceError:
		return "Service Layer Error"
	default:
		return "Unknown Error"
	}
}