package storage

const (
	queryGetUser = `
	select users.user_id as user_id,
	       users.login as login,
	       users.password as password,
	       users.otp_enabled as otp_enabled,
	       users.otp_key as otp_key,
	       roles.id as role_id,
	       roles.access as role_access,
	       roles.name as role_name
	from users join roles on roles.role_id = users.role_id
	where login = $1 or user_id = $2
`
	queryCreateUser = `
	insert into users( login, password, otp_enabled, otp_key, role_id)
	values(:login, :password, :otp_enabled, :otp_key, :role_id)
`
	queryUpdateUser = `
	update users
	set id = $3,
	    login = $4,
	    password = $5,
	    otp_enabled = $6,
	    otp_key = $7,
	    role_id = $8
	where login = $1 || id = $2
	returning 1
`
	queryDeleteUser = `
	delete from users
	where login = $1 or user_id = $2
	returning 1
`
	queryCreateRole = `
	insert into roles values($1, $2) 
	returning role_id;
`
	queryGetRole = `
	select id as role_id,
	       access as role_access,
	       name as role_name
	from roles
	where id = $1 or role_name = $2
`
	queryUpdateRole = `
	update roles
	set role_id = $3,
	    role_access = $4,
	    role_name = $5
	where role_id = $1 or role_name = $2
	returning 1
`
	queryDeleteRole = `
	delete from roles
	where role_id = $1 or role_name = $2
`
)
