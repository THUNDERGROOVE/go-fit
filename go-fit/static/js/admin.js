function ReloadTemplates() {
	var resp = HTTPGet('/admin/reload')
	if (resp == "success") {
		$("#alert-success-text").html("Successfully reloaded templates.")
		$("#alert-success").show()
	} else {
		$("#alert-error-text").html("Something bad happened while reloading templates.")
		$("#alert-error").show()
	}
}