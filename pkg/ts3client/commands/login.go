package commands

func Login(username, password string) Command {
	return Command{
		Name: "login",
		Parameters: map[string][]string{
			"client_login_name":     {username},
			"client_login_password": {password},
		},
		Options: nil,
	}
}
