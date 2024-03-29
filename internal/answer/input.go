package answer

type AnswerQuestionInput struct {
	QuestionId     string `json:"question_id"`
	Answer         string `json:"answer"`
	AnswerDuration int64  `json:"answer_duration"`
}

type AnswerToCorrect struct {
	QuestionId      string `json:"question_id"`
	AnswerId        string `json:"answer_id"`
	Answer          string `json:"answer"`
	ReferenceAnswer string `json:"reference_answer"`
	AnswerDuration  int64  `json:"answer_duration"`
}
