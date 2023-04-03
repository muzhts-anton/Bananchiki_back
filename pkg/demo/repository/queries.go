package demorep

const (
	queryGetAllPres = `
	SELECT id, code FROM presentation;
	`

	queryGetPresVm = `
	SELECT viewmode FROM presentation WHERE id = $1;
	`

	queryGetCurrentDemoSlideType = `
	SELECT slideorder.type, slideorder.item_id, presentation.demo_idx
	FROM slideorder
	JOIN presentation ON slideorder.presentation_id = presentation.id
	WHERE presentation.id = $1 AND slideorder.idx = presentation.demo_idx;
	`

	queryGetConvertedSlide = `
	SELECT name, width, height FROM convertedslide WHERE id = $1;
	`

	queryGetQuiz = `
	SELECT id, type, question, background, font_color, font_size, graph_color
	FROM quiz
	WHERE id = $1;
	`
)