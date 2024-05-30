/*
Functions for index.html

NETWORK FUNCTIONS
    completeTodo()  POST /todo/:id      update completion status of todo item
    updateTodo()    POST /todo/:id      update todo item using contenteditable fields
    deleteTodo()    DELETE /todo/:id    delete todo item
UTIL FUNCTIONS
    validateForm()                      client-side validation
*/


/*
Complete a todo item with the given ID
    @param {id} string      id of the todo item to update
*/
function completeTodo(id) {
    // Convert HTML elements into a form data object
    var formData = new FormData()
    formData.append('completed', document.getElementById('checkbox:' + id).checked)

    // Send form data via AJAX request
    var xhr = new XMLHttpRequest()
    xhr.open("POST", "/todo/" + id, true)
    xhr.onload = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
        console.log('POST checkbox data success')
        } else {
            console.error("POST checkbox data failed", xhr.status, xhr.statusText)
        }
    }
    xhr.send(formData)
}


/*
Update todo item using the above contenteditable fields 
    @param {id} string      id of the todo item to update
    @param {key} string     hash key of todo item field we are updating
*/
function updateTodo(id, key) {
    // Client-side form validation
    var value = document.getElementById(key + ":" + id).innerText
    if (!validateForm(key, value)) {
        console.error("Form validation failed", value)
        return
    }
    
    // Convert HTML element into a form data object
    var formData = new FormData()
    formData.append('id', id)
    formData.append(key, value)

    // Send form data via AJAX request
    var xhr = new XMLHttpRequest()
    xhr.open("POST", "/todo/" + id, true)
    xhr.onload = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
        console.log('POST form data success')
        } else {
            console.error("POST form data failed", xhr.status, xhr.statusText)
        }
    }
    xhr.send(formData)
}

/*
Delete a todo item with the given ID
    @param {id} string      id of the todo item to delete
*/
function deleteTodo(id) {
    var xhr = new XMLHttpRequest()
    xhr.open("DELETE", "/todo/" + id, true)
    xhr.onload = function () {
        if (xhr.readyState == 4 && xhr.status == 200) {
        console.log('DELETE success')
        } else {
            console.error("DELETE failed", xhr.status, xhr.statusText)
        }
    }
    xhr.send()
    window.location.href = "/"
}


/*
Client-side validation for form data
Server-side sanitation will also occur
    @param {key}   string      key for the todo item field (allows for conditional validation)
    @param {value} string      text to validate
*/
function validateForm(key, value) {
    // Sanitize HTML tags
    value = value.replace(/<[^>]*>?/gm, '');

    // Title cannot be empty
    return !(key == "title" && value.trim().length == 0)
}