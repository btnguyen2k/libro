# Libro - Data Model

## Introduction

Libro is multitenancy. That means Libro is able to host documentation of multiple products (entity `Product`). Each product documentation consists of multiple topics (entity `Topic`). And each topic can have one or more documentation pages (entity `Page`).

Also, the documentation hosted on Libro can be viewed publicly. In order to modify the documentation, however, user (entity `User`) must be authenticated.

Each entity is modeled as a business object (BO) and managed via a data-access-object (DAO). BO and DAO are built on top [henge](https://github.com/btnguyen2k/henge).


## User: BO & DAO

The business object `User` represents a user account.

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


## Product: BO & DAO

The business object `Product` represents a product to be documented.

**BO attributes:**

|Name   |Type  |Description|
|-------|------|-----------|
|name   |string|Product's name, for displaying purpose|
|ispub  |bool  |Flag to mark if the product is enabled/published|
|desc   |string|Description text|

> `Product` inherites other attibutes from [henge](https://github.com/btnguyen2k/henge)'s BO.

**DAO functions:**

|Function|Description|
|-------------------------------|-----------|
|`Delete(bo *Product) (bool, error)`|removes the specified business object from storage|
|`Create(bo *Product) (bool, error)`|persists a new business object to storage|
|`Get(id string) (*Product, error)` |retrieves a business object from storage|
|`GetN(fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Product, error)`|retrieves N business objects from storage|
|`GetAll(filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Product, error)`|retrieves all available business objects from storage|
|`Update(bo *Product) (bool, error)`|modifies an existing business object|


## Topic: BO & DAO

The business object `Topic` is a collection of documents on a same subject.

**BO attributes:**

|Name   |Type  |Description|
|-------|------|-----------|
|prod   |string|(*top-level attribute*) Id of the product this topic belongs to|
|title  |string|Topic title, for displaying purpose|
|icon   |string|Topic icon id, for displaying purpose|
|summary|string|Topic summary text|
|pos    |int   |Position index, for ordering/sorting purpose|

**DAO functions:**

|Function|Description|
|-------------------------------|-----------|
|`Delete(bo *Topic) (bool, error)`|removes the specified business object from storage|
|`Create(bo *Topic) (bool, error)`|persists a new business object to storage|
|`Get(id string) (*Topic, error)` |retrieves a business object from storage|
|`GetN(prod *Product, fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Topic, error)`|retrieves N business objects from storage|
|`GetAll(prod *Product, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Topic, error)`|retrieves all available business objects from storage|
|`Update(bo *Topic) (bool, error)`|modifies an existing business object|


## Page: BO & DAO

The business object `Page` represents a documentation page.

**BO attributes:**

|Name   |Type  |Description|
|-------|------|-----------|
|prod   |string|(*top-level attribute*) Id of the product this document page belongs to|
|topic  |string|(*top-level attribute*) Id of the topic this document page belongs to|
|title  |string|Page title, for displaying purpose|
|icon   |string|Page icon id, for displaying purpose|
|summary|string|Page summary text|
|pos    |int   |Position, for ordering/sorting purpose|
|content|string|Page content text|

**DAO functions:**

|Function|Description|
|-------------------------------|-----------|
|`Delete(bo *Page) (bool, error)`|removes the specified business object from storage|
|`Create(bo *Page) (bool, error)`|persists a new business object to storage|
|`Get(id string) (*Page, error)` |retrieves a business object from storage|
|`GetN(topic *Topic, fromOffset, maxNumRows int, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Page, error)`|retrieves N business objects from storage|
|`GetAll(topic *Topic, filter godal.FilterOpt, sorting *godal.SortingOpt) ([]*Page, error)`|retrieves all available business objects from storage|
|`Update(bo *Page) (bool, error)`|modifies an existing business object|
