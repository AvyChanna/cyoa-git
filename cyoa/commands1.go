package cyoa

type Context struct {
}

var Commands struct {
	Init   InitCmd   `cmd:"" help:"Init a repo for cyoa story"`
	Add    AddCmd    `cmd:"" help:"Add a chapter"`
	Remove RemoveCmd `cmd:"" help:"Remove a chapter"`
	Update UpdateCmd `cmd:"" help:"Update(Replace) a chapter"`
	List   ListCmd   `cmd:"" help:"Print list of chapters"`
	Move   MoveCmd   `cmd:"" help:"Move a chapter below another chapter"`
}
