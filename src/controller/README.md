# Controllers folder:

* this folder will contain all controllers regarding all routes:

* main.go file will handle all routes directly with the use of all the controllers router.HandleFunc("/", controllers.nameFunction), and serve all the files in the public folder(images, js, css)

* all the controllers will have this layout:

`type nameController struct{
 template *template.Template  }`

 `func (this *nameController) nameFunction(w http.ResponseWriter, req *http.Request){}`

   * we are going to retrieve all the models needed alongside the viewmodels which are used to create what needs to be displayed on the page

   * use mux package from gorilla to simplify the use of routers
