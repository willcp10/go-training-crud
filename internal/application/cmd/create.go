package cmd

func (u *UserCommandServiceImpl) CreateUser(cmd CreateUserCommand) error {
	return u.userRepository.CreateUser(cmd.User)
}