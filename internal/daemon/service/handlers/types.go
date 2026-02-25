package handlers

import "net/http"

type Head struct {
	Title       string
	Description string
	Keywords    string
}

const internalServerError = `
<!DOCTYPE html>
<body>
    <h1>Ошибка</h1>
    <p>Произошла ошибка при работе сайта, попробуйте обновить страницу магазина позднее или обратиться по телефону.</p>
    <p>Приносим свои извинения за предоставленные неудобства</p>
</body>
`

func HandleInternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(internalServerError))
}
