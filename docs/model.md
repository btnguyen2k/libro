# Libro - Data Model

Libro's business objects (BO) and data-access-objects (DAO) are built on top [henge](https://github.com/btnguyen2k/henge).

## App: BO & DAO

The business object `App` represents a registered application.

**BO attributes:**

|Name   |Type  |Description|
|-------|------|-----------|
|name |string|App's name, for displaying purpose|
|isvis|bool  |Flag to mark if the app is enabled/visible|
|desc |string|Description text|

**DAO functions:**

|Function|Description|
|-------------------------------|-----------|
|`Delete(bo *App) (bool, error)`|removes the specified business object from storage|
|`Create(bo *App) (bool, error)`|persists a new business object to storage|
|`Get(id string) (*App, error)` |retrieves a business object from storage|
|`GetN(fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*App, error)`|retrieves N business objects from storage|
|`GetAll(filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*App, error)`|retrieves all available business objects from storage|
|`Update(bo *App) (bool, error)`|modifies an existing business object|

## User: BO & DAO

The business object `User` represents a user account in the application.

**BO attributes:**

|Name |Type  |Description|
|-----|------|-----------|
|mid  |string|User's mask-id, also a unique id used when we do not wish to expose user's id|
|pwd  |string|User's password (hashed) (*)|
|dname|string|Display name, used for displaying purpose|
|isadm|bool  |Flag to mark if user has administrative privilege|

> `User` inherites other attibutes from [henge](https://github.com/btnguyen2k/henge)'s BO.

(*) User's password can be empty if authenticated via external source (e.g. [exter](https://github.com/btnguyen2k/exter)).

**DAO functions:**

|Function|Description|
|--------------------------------|-----------|
|`Delete(bo *User) (bool, error)`|removes the specified business object from storage|
|`Create(bo *User) (bool, error)`|persists a new business object to storage|
|`Get(id string) (*User, error)` |retrieves a business object from storage|
|`GetN(fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*User, error)`|retrieves N business objects from storage|
|`GetAll(filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*User, error)`|retrieves all available business objects from storage|
|`Update(bo *User) (bool, error)`|modifies an existing business object|
