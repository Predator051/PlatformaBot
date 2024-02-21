package redis

import "time"

func SetSession(session string) error {
	client := GetClient()
	status := client.Set(ctx, session, "user_session", time.Minute*60)

	return status.Err()
}

func SessionExist(session string) bool {
	client := GetClient()
	status := client.Exists(ctx, session)

	num, err := status.Result()
	return err != nil && num > 0
}
