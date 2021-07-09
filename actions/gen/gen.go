package gen


//go:generate ./scaffold_action.sh SubmitAttemptMutation --derive-from "insert_submissions_one"
//go:generate bash -c "cd ../ && gqlgenc"
