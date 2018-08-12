package main

const graphql = "https://api.github.com/graphql";

type auth struct {
	user string
	public_repo string
	repo string
	repo_deployment string
}

type repo struct {
	status string
	hook string
	org string
	public_key string
	gpg_key string
}