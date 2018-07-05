package auth

type UserManager interface {
	/*
	Get user's name
	 */
    GetUsername() string

    /*
    Always return False. This is a way of comparing User objects to
        anonymous users.
     */
    IsAnonymous() bool

    /*
    Always return True. This is a way to tell if the user has been
        authenticated in templates.
     */
	IsAuthenticated() bool


	/*
	set new password
	 */
	SetPassword(string) bool

	/*
	   Return a boolean of whether the raw_password was correct. Handles
        hashing formats behind the scenes.
	 */
	CheckPassword(string) bool

	GetFullName()  string

	GerShortName() string

	/*
	Returns True if the user has the specified permission. This method
        queries all available auth backends, but returns immediately if any
        backend returns True. Thus, a user who has permission from a single
        auth backend is assumed to have permission in general. If an object is
        provided, permissions for this specific object are checked.
	 */
	HasPerm(string) bool

	/*
	Returns True if the user has each of the specified permissions. If
        object is passed, it checks if the user has all required perms for this
        object.
	 */
    HasPerms([]string) bool

    /*
    Sends an email to this User.
     */
     EmailUser(string,string,string) error
}
