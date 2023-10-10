package database

func DestroyRoom(gid uint32) error {
	sqlMutex.Lock()
	defer sqlMutex.Unlock()
	_, err := SQLite.Exec(
		`DELETE FROM lm2_rooms WHERE gid=?`, gid,
	)
	if err != nil {
		return err
	}
	return nil
}
