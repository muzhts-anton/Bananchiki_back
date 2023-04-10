package profrep

const (
	queryGetUserInfo = `
	SELECT username, email, imgsrc FROM users WHERE id = $1;
	`

	queryGetAllPres = `
	SELECT id, name, code, quiz_num, converted_slide_num
	FROM presentation
	WHERE creator_id = $1;
	`
)
