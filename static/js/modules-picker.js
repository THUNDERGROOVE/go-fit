
function ItemCategoryChange() {
	var select = document.getElementById("category-picker");
	newSet = GetArrayFromMapping(select.value);
	var sel = document.getElementById("module-picker")
	sel.options.length = 0;
	for (var i = newSet.length - 1; i >= 0; i--) {
		sel.options.add(new Option(newSet[i].text, newSet[i].value))
	};
}
function GetArrayFromMapping(val) {
	var out = [];
	$.each(TypeMapping, function(k, v){
		if (k == val) {
			out = v
			return
		}
	});
	return out
}
function SetupPicker() {
	var select = document.getElementById("category-picker");
	select.options.length = 0;
	for (var i = moduleTypes.length - 1; i >= 0; i--) {
		select.options.add(new Option(moduleTypes[i].text,moduleTypes[i].text))
	};
}
SetupPicker()
ItemCategoryChange()
