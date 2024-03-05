package http

import (
	"net/http"
	"server/db/redis"
)

func AuthRequest(writer http.ResponseWriter, request *http.Request) {
	parsedBody := parseBody(request, &writer)

	if parsedBody["token"] != nil && parsedBody["token"] == AuthToken {
		session := randSession()

		err := redis.SetSession(session)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(session))
	} else {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte("There is no token!"))
	}

	return
}
