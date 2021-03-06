package actions

import (

  "fmt"
  "github.com/gobuffalo/buffalo"
  "github.com/gobuffalo/pop"
  "golang/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Employee)
// DB Table: Plural (employees)
// Resource: Plural (Employees)
// Path: Plural (/employees)
// View Template Folder: Plural (/templates/employees/)

// EmployeesResource is the resource for the Employee model
type EmployeesResource struct{
  buffalo.Resource
}

// List gets all Employees. This function is mapped to the path
// GET /employees
func (v EmployeesResource) List(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  employees := &models.Employees{}

  // Paginate results. Params "page" and "per_page" control pagination.
  // Default values are "page=1" and "per_page=20".
  q := tx.PaginateFromParams(c.Params())

  // Retrieve all Employees from the DB
  if err := q.All(employees); err != nil {
    return err
  }

  // Add the paginator to the context so it can be used in the template.
  c.Set("pagination", q.Paginator)

  return c.Render(200, r.Auto(c, employees))
}

// Show gets the data for one Employee. This function is mapped to
// the path GET /employees/{employee_id}
func (v EmployeesResource) Show(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Employee
  employee := &models.Employee{}

  // To find the Employee the parameter employee_id is used.
  if err := tx.Find(employee, c.Param("employee_id")); err != nil {
    return c.Error(404, err)
  }

  return c.Render(200, r.Auto(c, employee))
}

// New renders the form for creating a new Employee.
// This function is mapped to the path GET /employees/new
func (v EmployeesResource) New(c buffalo.Context) error {
  return c.Render(200, r.Auto(c, &models.Employee{}))
}
// Create adds a Employee to the DB. This function is mapped to the
// path POST /employees
func (v EmployeesResource) Create(c buffalo.Context) error {
  // Allocate an empty Employee
  employee := &models.Employee{}

  // Bind employee to the html form elements
  if err := c.Bind(employee); err != nil {
    return err
  }

  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Validate the data from the html form
  verrs, err := tx.ValidateAndCreate(employee)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the new.html template that the user can
    // correct the input.
    return c.Render(422, r.Auto(c, employee))
  }

  // If there are no errors set a success message
  c.Flash().Add("success", T.Translate(c, "employee.created.success"))
  // and redirect to the employees index page
  return c.Render(201, r.Auto(c, employee))
}

// Edit renders a edit form for a Employee. This function is
// mapped to the path GET /employees/{employee_id}/edit
func (v EmployeesResource) Edit(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Employee
  employee := &models.Employee{}

  if err := tx.Find(employee, c.Param("employee_id")); err != nil {
    return c.Error(404, err)
  }

  return c.Render(200, r.Auto(c, employee))
}
// Update changes a Employee in the DB. This function is mapped to
// the path PUT /employees/{employee_id}
func (v EmployeesResource) Update(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Employee
  employee := &models.Employee{}

  if err := tx.Find(employee, c.Param("employee_id")); err != nil {
    return c.Error(404, err)
  }

  // Bind Employee to the html form elements
  if err := c.Bind(employee); err != nil {
    return err
  }

  verrs, err := tx.ValidateAndUpdate(employee)
  if err != nil {
    return err
  }

  if verrs.HasAny() {
    // Make the errors available inside the html template
    c.Set("errors", verrs)

    // Render again the edit.html template that the user can
    // correct the input.
    return c.Render(422, r.Auto(c, employee))
  }

  // If there are no errors set a success message
  c.Flash().Add("success", T.Translate(c, "employee.updated.success"))
  // and redirect to the employees index page
  return c.Render(200, r.Auto(c, employee))
}

// Destroy deletes a Employee from the DB. This function is mapped
// to the path DELETE /employees/{employee_id}
func (v EmployeesResource) Destroy(c buffalo.Context) error {
  // Get the DB connection from the context
  tx, ok := c.Value("tx").(*pop.Connection)
  if !ok {
    return fmt.Errorf("no transaction found")
  }

  // Allocate an empty Employee
  employee := &models.Employee{}

  // To find the Employee the parameter employee_id is used.
  if err := tx.Find(employee, c.Param("employee_id")); err != nil {
    return c.Error(404, err)
  }

  if err := tx.Destroy(employee); err != nil {
    return err
  }

  // If there are no errors set a flash message
  c.Flash().Add("success", T.Translate(c, "employee.destroyed.success"))
  // Redirect to the employees index page
  return c.Render(200, r.Auto(c, employee))
}
