package cmd

func (u *UserCommandServiceImpl) DeleteUser(cmd DeleteUserCommand) error {
	return u.userRepository.DeleteUser(cmd.ID)
}