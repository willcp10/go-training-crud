package cmd

func (u *UserCommandServiceImpl) UpdateUser(cmd UpdateUserCommand) error {
	return u.userRepository.UpdateUser(cmd.User)
}
