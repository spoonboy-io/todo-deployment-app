package templates

import "fmt"

func Index() string {

	// change this to demo different build
	title := "Morpheus Todo List"
	version := "Version 1.0.0"
	backgroundColor := "#666666"

	// the template don't edit this
	template := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css"/>
    <link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Open+Sans:ital,wght@0,300;0,400;0,600;0,700;0,800;1,300;1,400;1,600;1,700;1,800&amp;display=swap"/>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css"/>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-datepicker/1.9.0/css/bootstrap-datepicker.standalone.min.css"/>


    <title>%s</title>

    <style>
        body {
            font-family: "Open Sans", sans-serif;
            line-height: 1.6;
			background-color: %s
        }

        .add-todo-input,
        .edit-todo-input {
            outline: none;
        }

        .add-todo-input:focus,
        .edit-todo-input:focus {
            border: none !important;
            box-shadow: none !important;
        }

        .view-opt-label,
        .date-label {
            font-size: 0.8rem;
        }

        .edit-todo-input {
            font-size: 1.7rem !important;
        }

        .todo-actions {
            visibility: hidden !important;
        }

        .todo-item:hover .todo-actions {
            visibility: visible !important;
        }

        .todo-item.editing .todo-actions .edit-icon {
            display: none !important;
        }
    </style>
</head>
<body>
<div class="container m-5 p-2 rounded mx-auto bg-light shadow">
    <!-- App title section -->
    <div class="row m-1 p-4">
        <div class="col">
            <div class="p-1 h1 text-primary text-center mx-auto display-inline-block">
                <i class="fa fa-check bg-primary text-white rounded p-2"></i>
                <p>%s</p>
				<p style="font-size:22px;">%s</p>
            </div>
        </div>
    </div>
    <!-- Create todo section -->
    <div class="row m-1 p-3">
        <div class="col col-11 mx-auto">
            <form action="/" method="POST" class="row bg-white rounded shadow-sm p-2 add-todo-wrapper align-items-center justify-content-center">
                <div class="col">
                    <input name="Item" class="form-control form-control-lg border-0 add-todo-input bg-transparent rounded" type="text" placeholder="Add new ..">
                </div>
                <div class="col-auto px-0 mx-0 mr-2">
                    <button type="submit" class="btn btn-primary">Add</button>
                </div>
            </form>
        </div>
    </div>
    <div class="p-2 m-2 mx-4 border-black-25 border-bottom"></div>
    <!-- Todo list section -->
    <div class="row mx-1 px-5 pb-3 w-80">
        <div class="col mx-auto">
            <!-- Todo Item-->
            {{range .Todos}}
            <div class="row px-3 align-items-center todo-item editing rounded">
                <div class="col px-1 m-1 d-flex align-items-center">
                    <input type="text" class="form-control form-control-lg border-0 edit-todo-input bg-transparent rounded px-3 d-none" readonly value="{{.}}" title="{{.}}" />
                    <input  id="{{.}}"  type="text" class="form-control form-control-lg border-0 edit-todo-input rounded px-3" value="{{.}}" />
                </div>
                <div class="col-auto m-1 p-0 px-3 d-none">
                </div>
                <div class="col-auto m-1 p-0 todo-actions">
                    <div class="row d-flex align-items-center justify-content-end">
                        <h5 class="m-0 p-0 px-2">
                            <i onclick="updateDb('{{.}}')" class="fa fa-pencil text-warning btn m-0 p-0" data-toggle="tooltip" data-placement="bottom" title="Edit todo"></i>
                        </h5>
                        <h5 class="m-0 p-0 px-2">
                            <i onclick="removeFromDb('{{.}}')" class="fa fa-trash-o text-danger btn m-0 p-0" data-toggle="tooltip" data-placement="bottom" title="Delete todo"></i>
                        </h5>
                    </div>
                </div>
            </div>
            {{end}}
        </div>
    </div>
</div>
</form>
<script>
    function removeFromDb(item){
		let req = "/delete?item=" + item;
        fetch(req, {method: "Delete"}).then(res =>{
            if (res.status == 200){
                window.location.pathname = "/"
            }
        })
    }

    function updateDb(item) {
        let input = document.getElementById(item)
        let newitem = input.value
		let req = "/update?olditem=" + item + "&newitem=" + newitem
        fetch(req, {method: "PUT"}).then(res =>{
            if (res.status == 200){
                alert("Database updated")
                window.location.pathname = "/"
            }
        })
    }
</script>
<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js"></script>
<script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js"></script>
<script src="https://stackpath.bootstrapcdn.com/bootlint/1.1.0/bootlint.min.js"></script>
</body>
</html>
`

	return fmt.Sprintf(template, title, backgroundColor, title, version)
}