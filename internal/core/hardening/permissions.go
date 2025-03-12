package hardening

import (
    "errors"
)

// Permission represents a user permission
type Permission struct {
    Name        string
    Description string
}

// User represents a user with permissions
type User struct {
    Username    string
    Permissions []Permission
}

// NewPermission creates a new permission
func NewPermission(name, description string) Permission {
    return Permission{
        Name:        name,
        Description: description,
    }
}

// AddPermission adds a permission to a user
func (u *User) AddPermission(permission Permission) {
    u.Permissions = append(u.Permissions, permission)
}

// HasPermission checks if a user has a specific permission
func (u *User) HasPermission(permissionName string) bool {
    for _, p := range u.Permissions {
        if p.Name == permissionName {
            return true
        }
    }
    return false
}

// RemovePermission removes a permission from a user
func (u *User) RemovePermission(permissionName string) error {
    for i, p := range u.Permissions {
        if p.Name == permissionName {
            u.Permissions = append(u.Permissions[:i], u.Permissions[i+1:]...)
            return nil
        }
    }
    return errors.New("permission not found")
}
Explanation:
