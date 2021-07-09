package action

type Mutation struct {
	Unlock *UnlockMutationOutput `json:"unlock"`
}

type SubmitFlagMutationOutput struct {
	Id int `json:"id"`
}

type PurchaseUnlockableOutput struct {
	Id int `json:"id"`
}

type PurchaseUnlockableMutationOutput struct {
	Id int `json:"id"`
}

type UnlockMutationOutput struct {
	Id int `json:"id"`
}

type JsonWebToken struct {
	Token string `json:"token"`
}

type CreateUserOutput struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type LogoutResponse struct {
	Success *bool `json:"success"`
}

type SubmitMutationOutput struct {
	Id int `json:"id"`
}

type SubmitAttemptMutationOutput struct {
	Id int `json:"id"`
}

type RegisterOutput struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

type SubmitArgs struct {
	Challenge_id int    `json:"challenge_id"`
	Provided     string `json:"provided"`
}

type UnlockArgs struct {
	Target *int    `json:"target"`
	Type   *string `json:"type"`
}

type GraphQLError struct {
	Message string `json:"message"`
}
