package model

type Chat struct{
	ID int `json:"id" db:"id"`
	FirstMemberID int `json:"firstMemberId" db:"first_member_id"`
	SecondMemberID int `json:"secondMemberId" db:"second_member_id"`
}